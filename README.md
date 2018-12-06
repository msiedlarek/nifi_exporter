# Nifi Exporter

Interrogates a running Apache NiFi instance and makes the following
information available via a metrics endpoint to a prometheus scraper:

## Configuration

The configuration comes from the configuration yaml file passed to the exporter at startup
e.g.

``` bash
./nifi_exporter /etc/nifi_exporter/config.yml
```

### Minimal configuration

The nifi exporter expects at least the following keys to be present in the yaml config file:

``` yaml
---
exporter:
  listenAddress: 127.0.0.1:9103
nodes:
  - url: http://localhost:8080
    username: xxxxxx
    password: xxxxxx
```

## Example output

``` text
# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 12
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 1.737064e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 1.737064e+06
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 1.443204e+06
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 1395
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 2.234368e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 1.737064e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 6.3307776e+07
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 3.178496e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 11434
# HELP go_memstats_heap_released_bytes_total Total number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes_total counter
go_memstats_heap_released_bytes_total 0
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 6.6486272e+07
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 0
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 12829
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 6912
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 45752
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 49152
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.473924e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 1.301364e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 622592
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 622592
# HELP go_memstats_sys_bytes Number of bytes obtained by system. Sum of all system allocations.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 7.2153336e+07
# HELP nifi_info NiFi version info.
# TYPE nifi_info gauge
nifi_info{env="testing",node_id="aggregate",version="1.7.1"} 1
# HELP nifi_jvm_available_processor_count Number of available processors if supported by the underlying system.
# TYPE nifi_jvm_available_processor_count gauge
nifi_jvm_available_processor_count{env="testing",node_id="aggregate"} 36
# HELP nifi_jvm_daemon_thread_count Number of daemon threads.
# TYPE nifi_jvm_daemon_thread_count gauge
nifi_jvm_daemon_thread_count{env="testing",node_id="aggregate"} 28
# HELP nifi_jvm_free_heap_bytes The number of bytes that are allocated to the JVM heap but not currently being used.
# TYPE nifi_jvm_free_heap_bytes gauge
nifi_jvm_free_heap_bytes{env="testing",node_id="aggregate"} 9.637522384e+09
# HELP nifi_jvm_free_non_heap_bytes Total number of free non-heap bytes available to the JVM.
# TYPE nifi_jvm_free_non_heap_bytes gauge
nifi_jvm_free_non_heap_bytes{env="testing",node_id="aggregate"} 1.086156e+07
# HELP nifi_jvm_gc_collection_time_ms_total The total number of milliseconds spent garbage collecting.
# TYPE nifi_jvm_gc_collection_time_ms_total counter
nifi_jvm_gc_collection_time_ms_total{env="testing",gc="G1 Old Generation",node_id="aggregate"} 0
nifi_jvm_gc_collection_time_ms_total{env="testing",gc="G1 Young Generation",node_id="aggregate"} 1565
# HELP nifi_jvm_gc_collection_total The number of times garbage collection has run.
# TYPE nifi_jvm_gc_collection_total counter
nifi_jvm_gc_collection_total{env="testing",gc="G1 Old Generation",node_id="aggregate"} 0
nifi_jvm_gc_collection_total{env="testing",gc="G1 Young Generation",node_id="aggregate"} 54
# HELP nifi_jvm_heap_bytes The total number of bytes that are available for the JVM heap to use.
# TYPE nifi_jvm_heap_bytes gauge
nifi_jvm_heap_bytes{env="testing",node_id="aggregate"} 1.2884901888e+10
# HELP nifi_jvm_info JVM version info.
# TYPE nifi_jvm_info gauge
nifi_jvm_info{env="testing",node_id="aggregate",vendor="Oracle Corporation",version="1.8.0_181"} 1
# HELP nifi_jvm_max_heap_bytes The maximum number of bytes that can be used by the JVM.
# TYPE nifi_jvm_max_heap_bytes gauge
nifi_jvm_max_heap_bytes{env="testing",node_id="aggregate"} 1.2884901888e+10
# HELP nifi_jvm_max_non_heap_bytes The maximum number of bytes that the JVM can use for non-heap purposes.
# TYPE nifi_jvm_max_non_heap_bytes gauge
nifi_jvm_max_non_heap_bytes{env="testing",node_id="aggregate"} -1
# HELP nifi_jvm_non_heap_bytes Total number of bytes allocated to the JVM not used for heap.
# TYPE nifi_jvm_non_heap_bytes gauge
nifi_jvm_non_heap_bytes{env="testing",node_id="aggregate"} 2.03776e+08
# HELP nifi_jvm_processor_load_avg The processor load average if supported by the underlying system.
# TYPE nifi_jvm_processor_load_avg gauge
nifi_jvm_processor_load_avg{env="testing",node_id="aggregate"} 0
# HELP nifi_jvm_thread_count Total number of threads.
# TYPE nifi_jvm_thread_count gauge
nifi_jvm_thread_count{env="testing",node_id="aggregate"} 139
# HELP nifi_jvm_used_heap_bytes The number of bytes of JVM heap that are currently being used.
# TYPE nifi_jvm_used_heap_bytes gauge
nifi_jvm_used_heap_bytes{env="testing",node_id="aggregate"} 3.247379504e+09
# HELP nifi_jvm_used_non_heap_bytes Total number of bytes used by the JVM not in the heap space.
# TYPE nifi_jvm_used_non_heap_bytes gauge
nifi_jvm_used_non_heap_bytes{env="testing",node_id="aggregate"} 1.9291444e+08
# HELP nifi_os_info Operating system version info.
# TYPE nifi_os_info gauge
nifi_os_info{arch="amd64",env="testing",name="Linux",node_id="aggregate",version="4.14.72-73.55.amzn2.x86_64"} 1
# HELP nifi_conn_flow_files_queued The number of FlowFiles that are currently queued in the connection
# TYPE nifi_conn_flow_files_queued gauge
nifi_conn_flow_files_queued{connection_id="026073af-c1c7-3d0e-2b43-e8baf6bb875b",connection_name="success",destination_name="in",env="testing",group_id="57b9c87e-7ee2-3cf1-2a76-5d75762bb2dd",node_id="aggregate",source_name="GenerateFlowFile"} 0
nifi_conn_flow_files_queued{connection_id="06e8674b-a9a2-300a-5fea-3bcb312a7a0e",connection_name="failure",destination_name="fail5",env="testing",group_id="053742a0-5d93-379d-8232-ff6e6c21794d",node_id="aggregate",source_name="ExecuteSQL"} 4
nifi_conn_flow_files_queued{connection_id="0818d0fd-9d0f-36b4-db68-596dfe02569e",connection_name="success",destination_name="out",env="testing",group_id="48f00c1b-c9c7-30cf-e3f8-52f0745f3dfd",node_id="aggregate",source_name="PutDistributedMapCache"} 0
nifi_conn_flow_files_queued{connection_id="0bd6ab8a-29f8-37df-b345-16bdcf1abb9a",connection_name="success",destination_name="ExecuteScript",env="testing",group_id="263a63da-e085-3f02-1818-f5dfaa5504f0",node_id="aggregate",source_name="Set Credentials & Properties"} 0
nifi_conn_flow_files_queued{connection_id="134df201-97bf-3eed-f937-445c48af300d",connection_name="splits",destination_name="ExtractText",env="testing",group_id="d99c18ca-e627-3085-7280-3aabf4c44eba",node_id="aggregate",source_name="SplitText"} 0
nifi_conn_flow_files_queued{connection_id="16122638-c888-3192-a3f2-c6578fa7bced",connection_name="Response",destination_name="RouteOnContent",env="testing",group_id="053742a0-5d93-379d-8232-ff6e6c21794d",node_id="aggregate",source_name="InvokeHTTP"} 0
nifi_conn_flow_files_queued{connection_id="1d681a65-c462-30eb-0f17-c56423642902",connection_name="No Retry, Failure",destination_name="UpdateAttribute",env="testing",group_id="053742a0-5d93-379d-8232-ff6e6c21794d",node_id="aggregate",source_name="InvokeHTTP"} 0
# HELP nifi_pg_active_thread_count The active thread count for this process group.
# TYPE nifi_pg_active_thread_count gauge
nifi_pg_active_thread_count{env="testing",group="my_processor_group",node_id="aggregate"} 0
# HELP nifi_pg_bulletin_5m_count Number of bulletins posted during last 5 minutes.
# TYPE nifi_pg_bulletin_5m_count gauge
nifi_pg_bulletin_5m_count{env="testing",group="my_processor_group",level="ERROR"} 0
nifi_pg_bulletin_5m_count{env="testing",group="my_processor_group",level="INFO"} 0
nifi_pg_bulletin_5m_count{env="testing",group="my_processor_group",level="WARNING"} 0
# HELP nifi_pg_component_count The number of components in this process group.
# TYPE nifi_pg_component_count gauge
nifi_pg_component_count{env="testing",group="my_processor_group",status="disabled"} 0
nifi_pg_component_count{env="testing",group="my_processor_group",status="invalid"} 2
nifi_pg_component_count{env="testing",group="my_processor_group",status="running"} 0
nifi_pg_component_count{env="testing",group="my_processor_group",status="stopped"} 5
# HELP nifi_pg_in_bytes_5m_count The number of bytes that have come into this ProcessGroup in the last 5 minutes
# TYPE nifi_pg_in_bytes_5m_count gauge
nifi_pg_in_bytes_5m_count{env="testing",group="my_processor_group",node_id="aggregate"} 0
# HELP nifi_pg_in_flow_files_5m_count The number of FlowFiles that have come into this ProcessGroup in the last 5 minutes
# TYPE nifi_pg_in_flow_files_5m_count gauge
nifi_pg_in_flow_files_5m_count{env="testing",group="my_processor_group",node_id="aggregate"} 0
# HELP nifi_pg_out_bytes_5m_count The number of bytes transferred out of this ProcessGroup in the last 5 minutes
# TYPE nifi_pg_out_bytes_5m_count gauge
nifi_pg_out_bytes_5m_count{env="testing",group="my_processor_group",node_id="aggregate"} 0
# HELP nifi_pg_out_flow_files_5m_count The number of FlowFiles transferred out of this ProcessGroup in the last 5 minutes
# TYPE nifi_pg_out_flow_files_5m_count gauge
nifi_pg_out_flow_files_5m_count{env="testing",group="my_processor_group",node_id="aggregate"} 0
# HELP nifi_pg_queued_bytes The number of bytes that are queued up in this ProcessGroup right now
# TYPE nifi_pg_queued_bytes gauge
nifi_pg_queued_bytes{env="testing",group="my_processor_group",node_id="aggregate"} 0
# HELP nifi_pg_queued_flow_files_count The number of FlowFiles that are queued up in this ProcessGroup right now
# TYPE nifi_pg_queued_flow_files_count gauge
nifi_pg_queued_flow_files_count{env="testing",group="my_processor_group",node_id="aggregate"} 0
# HELP nifi_pg_read_bytes_5m_count The number of bytes read by components in this ProcessGroup in the last 5 minutes
# TYPE nifi_pg_read_bytes_5m_count gauge
nifi_pg_read_bytes_5m_count{env="testing",group="my_processor_group",node_id="aggregate"} 0
# HELP nifi_pg_received_bytes_5m_count The number of bytes received from external sources by components within this ProcessGroup in the last 5 minutes
# TYPE nifi_pg_received_bytes_5m_count gauge
nifi_pg_received_bytes_5m_count{env="testing",group="my_processor_group",node_id="aggregate"} 0
# HELP nifi_pg_received_flow_files_5m_count The number of FlowFiles received from external sources by components within this ProcessGroup in the last 5 minutes
# TYPE nifi_pg_received_flow_files_5m_count gauge
nifi_pg_received_flow_files_5m_count{env="testing",group="my_processor_group",node_id="aggregate"} 0
# HELP nifi_pg_sent_bytes_5m_count The number of bytes sent to an external sink by components within this ProcessGroup in the last 5 minutes
# TYPE nifi_pg_sent_bytes_5m_count gauge
nifi_pg_sent_bytes_5m_count{env="testing",group="my_processor_group",node_id="aggregate"} 0
# HELP nifi_pg_sent_flow_files_5m_count The number of FlowFiles sent to an external sink by components within this ProcessGroup in the last 5 minutes
# TYPE nifi_pg_sent_flow_files_5m_count gauge
nifi_pg_sent_flow_files_5m_count{env="testing",group="my_processor_group",node_id="aggregate"} 0
# HELP nifi_pg_transferred_bytes_5m_count The number of bytes transferred in this ProcessGroup in the last 5 minutes
# TYPE nifi_pg_transferred_bytes_5m_count gauge
nifi_pg_transferred_bytes_5m_count{env="testing",group="my_processor_group",node_id="aggregate"} 0
# HELP nifi_pg_transferred_flow_files_5m_count The number of FlowFiles transferred in this ProcessGroup in the last 5 minutes
# TYPE nifi_pg_transferred_flow_files_5m_count gauge
nifi_pg_transferred_flow_files_5m_count{env="testing",group="my_processor_group",node_id="aggregate"} 0
# HELP nifi_pg_written_bytes_5m_count The number of bytes written by components in this ProcessGroup in the last 5 minutes
# TYPE nifi_pg_written_bytes_5m_count gauge
nifi_pg_written_bytes_5m_count{env="testing",group="my_processor_group",node_id="aggregate"} 0
# HELP nifi_stor_free_bytes The number of bytes of free space.
# TYPE nifi_stor_free_bytes gauge
nifi_stor_free_bytes{env="testing",location="default",node_id="aggregate",type="content"} 1.92030683136e+11
nifi_stor_free_bytes{env="testing",location="default",node_id="aggregate",type="flow_file"} 1.92030683136e+11
nifi_stor_free_bytes{env="testing",location="default",node_id="aggregate",type="provenance"} 1.92030683136e+11
# HELP nifi_stor_size_bytes The number of bytes of total space.
# TYPE nifi_stor_size_bytes gauge
nifi_stor_size_bytes{env="testing",location="default",node_id="aggregate",type="content"} 2.14735761408e+11
nifi_stor_size_bytes{env="testing",location="default",node_id="aggregate",type="flow_file"} 2.14735761408e+11
nifi_stor_size_bytes{env="testing",location="default",node_id="aggregate",type="provenance"} 2.14735761408e+11
# HELP nifi_stor_used_bytes The number of bytes of used space.
# TYPE nifi_stor_used_bytes gauge
nifi_stor_used_bytes{env="testing",location="default",node_id="aggregate",type="content"} 2.2705078272e+10
nifi_stor_used_bytes{env="testing",location="default",node_id="aggregate",type="flow_file"} 2.2705078272e+10
nifi_stor_used_bytes{env="testing",location="default",node_id="aggregate",type="provenance"} 2.2705078272e+10
```

## Building

Set up a [golang environment](https://golang.org/doc/install)
Checkout this repository to `$GOPATH/src/github.com/msiedlarek/nifi_exporter`
From the root of this repository run `go build`