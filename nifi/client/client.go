package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"

	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
)

const tokenExpirationMargin = time.Minute

type Credentials struct {
	Username string
	Password string
}

type Client struct {
	baseURL     string
	client      http.Client
	credentials url.Values

	token                    string
	tokenMx                  sync.Mutex
	tokenExpirationTimestamp int64
}

type jwtPayload struct {
	Audience          string `json:"aud"`
	ExpirationTime    int64  `json:"exp"`
	IssuedAt          int64  `json:"iat"`
	Issuer            string `json:"iss"`
	Kid               int    `json:"kid"`
	PreferredUsername string `json:"preferred_username"`
	Subject           string `json:"sub"`
}

func NewClient(baseURL, username, password, caCertificates string) (*Client, error) {
	c := Client{
		baseURL: strings.TrimRight(baseURL, "/") + "/nifi-api",
		credentials: url.Values{
			"username": []string{username},
			"password": []string{password},
		},
	}
	if caCertificates != "" {
		certPool := x509.NewCertPool()
		if ok := certPool.AppendCertsFromPEM([]byte(caCertificates)); !ok {
			return nil, errors.New("Invalid CA certificates.")
		}
		for _, der := range certPool.Subjects() {
			var rdn pkix.RDNSequence
			if _, err := asn1.Unmarshal(der, &rdn); err != nil {
				return nil, errors.Trace(err)
			}
			var name pkix.Name
			name.FillFromRDNSequence(&rdn)
			log.WithFields(log.Fields{
				"commonName":   name.CommonName,
				"organization": name.Organization,
			}).Infof("Loaded CA certificate for %s: %s", baseURL, name.CommonName)
		}
		c.client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: certPool,
			},
		}
	}
	return &c, nil
}

func (c *Client) GetCounters(nodewise bool, clusterNodeId string) (*CountersDTO, error) {
	query := url.Values{}
	if nodewise {
		query.Add("nodewise", "1")
	} else {
		query.Add("nodewise", "0")
	}
	if len(clusterNodeId) > 0 {
		query.Add("clusterNodeId", clusterNodeId)
	}

	var entity CountersEntity
	if err := c.request("/counters", query, &entity); err != nil {
		return nil, errors.Trace(err)
	}
	return &entity.Counters, nil
}

func (c *Client) GetProcessGroup(id string) (*ProcessGroupEntity, error) {
	var entity ProcessGroupEntity
	if err := c.request("/process-groups/"+id, nil, &entity); err != nil {
		return nil, errors.Trace(err)
	}
	return &entity, nil
}

func (c *Client) GetProcessGroups(parentID string) ([]ProcessGroupEntity, error) {
	var entity ProcessGroupsEntity
	if err := c.request("/process-groups/"+parentID+"/process-groups", nil, &entity); err != nil {
		return nil, errors.Trace(err)
	}
	return entity.ProcessGroups, nil
}

// GetConnections traverses the process group hierarchy returning information about
// all connections
func (c *Client) GetConnections(parentID string) ([]ConnectionEntity, error) {
	var entity ConnectionsEntity
	if err := c.getDeepConnections(parentID, &entity); err != nil {
		return nil, err
	}
	return entity.Connections, nil

}

func (c *Client) getDeepConnections(parentID string, connectionsEntity *ConnectionsEntity) error {
	var entity ConnectionsEntity

	// Get the connections for the current process group
	if err := c.request("/process-groups/"+parentID+"/connections", nil, &entity); err != nil {
		return errors.Trace(err)
	}

	// And the child process groups

	var pgentity ProcessGroupsEntity
	if err := c.request("/process-groups/"+parentID+"/process-groups", nil, &pgentity); err != nil {
		return errors.Trace(err)
	}

	for _, pg := range pgentity.ProcessGroups {
		if err := c.getDeepConnections(pg.ID, connectionsEntity); err != nil {
			return err
		}
	}
	connectionsEntity.Connections = append(connectionsEntity.Connections, entity.Connections...)
	return nil
}

// GetDeepProcessGroups traverses the process group hierarchy returning information about
// this and all child process groups
func (c *Client) GetDeepProcessGroups(parentID string) ([]ProcessGroupEntity, error) {
	var entity ProcessGroupsEntity
	if err := c.getDeepProcessGroups(parentID, &entity); err != nil {
		return nil, err
	}
	return entity.ProcessGroups, nil

}

func (c *Client) getDeepProcessGroups(parentID string, groupsEntity *ProcessGroupsEntity) error {
	var entity ProcessGroupsEntity
	if err := c.request("/process-groups/"+parentID+"/process-groups", nil, &entity); err != nil {
		return errors.Trace(err)
	}

	for _, pg := range entity.ProcessGroups {
		if err := c.getDeepProcessGroups(pg.ID, groupsEntity); err != nil {
			return err
		}
	}
	groupsEntity.ProcessGroups = append(groupsEntity.ProcessGroups, entity.ProcessGroups...)
	return nil
}

func (c *Client) GetSystemDiagnostics(nodewise bool, clusterNodeId string) (*SystemDiagnosticsDTO, error) {
	query := url.Values{}
	if nodewise {
		query.Add("nodewise", "1")
	} else {
		query.Add("nodewise", "0")
	}
	if len(clusterNodeId) > 0 {
		query.Add("clusterNodeId", clusterNodeId)
	}

	var entity SystemDiagnosticsEntity
	if err := c.request("/system-diagnostics", query, &entity); err != nil {
		return nil, errors.Trace(err)
	}
	return &entity.SystemDiagnostics, nil
}

func (c *Client) request(path string, query url.Values, responseEntity interface{}) error {
	token, err := c.getToken()
	if err != nil {
		return errors.Trace(err)
	}

	reqURL := c.baseURL + path
	log.WithField("url", reqURL).Info("Requesting api resource......")
	if query != nil && len(query) > 0 {
		reqURL += "?" + query.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return errors.Annotate(err, "Error while preparing API request")
	}
	req.Header.Add("Authorization", token)

	resp, err := c.client.Do(req)
	if err != nil {
		return errors.Annotate(err, "NiFi API request failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(responseEntity); err != nil {
			return errors.Annotate(err, "Invalid JSON response from NiFi")
		}
		return nil
	}

	messageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Annotate(err, "Couldn't read error message from API response")
	}
	message := fmt.Sprintf(
		"API call returned an error: %s: %s",
		resp.Status,
		string(messageBytes),
	)

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		return errors.Unauthorizedf(message)
	} else {
		return errors.New(message)
	}

}

func (c *Client) getToken() (string, error) {
	if atomic.LoadInt64(&c.tokenExpirationTimestamp) < time.Now().Add(tokenExpirationMargin).Unix() {
		c.authenticate()
	}
	return c.token, nil
}

func (c *Client) authenticate() error {
	c.tokenMx.Lock()
	defer c.tokenMx.Unlock()
	if c.tokenExpirationTimestamp > time.Now().Add(tokenExpirationMargin).Unix() {
		return nil
	}
	log.WithFields(log.Fields{
		"url":      c.baseURL,
		"username": c.credentials.Get("username"),
	}).Info("Authentication token has expired, reauthenticating...")

	resp, err := c.client.PostForm(c.baseURL+"/access/token", c.credentials)
	if err != nil {
		return errors.Annotate(err, "Couldn't request access token from NiFi")
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Annotate(err, "Couldn't read access token response from NiFi")
	}
	body := strings.TrimSpace(string(bodyBytes))
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		jwtParts := strings.SplitN(body, ".", 3)
		if len(jwtParts) < 2 {
			return errors.Annotate(err, "Invalid access token response from NiFi: Missing JWT payload")
		}
		jwtPayloadJson, err := base64.RawURLEncoding.DecodeString(jwtParts[1])
		if err != nil {
			return errors.Annotate(err, "Invalid access token response from NiFi: Payload is not valid Base64")
		}
		var payload jwtPayload
		if err := json.Unmarshal(jwtPayloadJson, &payload); err != nil {
			return errors.Annotate(err, "Invalid access token response from NiFi: Payload is not valid JSON")
		}

		c.token = "Bearer " + body
		atomic.StoreInt64(&c.tokenExpirationTimestamp, payload.ExpirationTime)

		log.WithFields(log.Fields{
			"url":             c.baseURL,
			"username":        c.credentials.Get("username"),
			"tokenExpiration": time.Unix(c.tokenExpirationTimestamp, 0).String(),
		}).Info("Authentication successful.")
		return nil
	} else if resp.StatusCode == http.StatusUnauthorized {
		return errors.Unauthorizedf(body)
	} else {
		return errors.New(body)
	}
}
