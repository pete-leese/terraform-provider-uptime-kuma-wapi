---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ukumawapi_monitor Resource - terraform-provider-ukumawapi"
subcategory: ""
description: |-
  Monitor Resource
---

# ukumawapi_monitor (Resource)

Monitor Resource

## Example Usage

```terraform
resource "ukumawapi_monitor" "my_monitor" {
  accepted_statuscodes = [
    "{ \"see\": \"documentation\" }"
  ]
  auth_domain                = "...my_auth_domain..."
  auth_method                = "...my_auth_method..."
  auth_workstation           = "...my_auth_workstation..."
  basic_auth_pass            = "...my_basic_auth_pass..."
  basic_auth_user            = "...my_basic_auth_user..."
  body                       = "...my_body..."
  database_connection_string = "...my_database_connection_string..."
  database_query             = "...my_database_query..."
  dns_resolve_server         = "...my_dns_resolve_server..."
  dns_resolve_type           = "...my_dns_resolve_type..."
  docker_container           = "...my_docker_container..."
  docker_host                = 3
  expiry_notification        = true
  headers                    = "...my_headers..."
  hostname                   = "...my_hostname..."
  ignore_tls                 = true
  interval                   = 0
  keyword                    = "...my_keyword..."
  maxredirects               = 6
  maxretries                 = 9
  method                     = "...my_method..."
  monitor_id                 = 2
  mqtt_password              = "...my_mqtt_password..."
  mqtt_success_message       = "...my_mqtt_success_message..."
  mqtt_topic                 = "...my_mqtt_topic..."
  mqtt_username              = "...my_mqtt_username..."
  name                       = "...my_name..."
  notification_id_list = [
    "{ \"see\": \"documentation\" }"
  ]
  port                      = 6
  proxy_id                  = 8
  radius_called_station_id  = "...my_radius_called_station_id..."
  radius_calling_station_id = "...my_radius_calling_station_id..."
  radius_password           = "...my_radius_password..."
  radius_secret             = "...my_radius_secret..."
  radius_username           = "...my_radius_username..."
  resend_interval           = 7
  retry_interval            = 8
  type                      = "sqlserver"
  upside_down               = true
  url                       = "...my_url..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `monitor_id` (Number)
- `name` (String)
- `type` (String) Enumerate monitor types. must be one of ["http", "port", "ping", "keyword", "grpc-keyword", "dns", "docker", "push", "steam", "gamedig", "mqtt", "sqlserver", "postgres", "mysql", "mongodb", "radius", "redis"]

### Optional

- `accepted_statuscodes` (List of String)
- `auth_domain` (String)
- `auth_method` (String) Enumerate authentication methods for monitors. Default: ""; must be one of ["", "basic", "ntlm", "mtls"]
- `auth_workstation` (String)
- `basic_auth_pass` (String)
- `basic_auth_user` (String)
- `body` (String)
- `database_connection_string` (String)
- `database_query` (String)
- `dns_resolve_server` (String) Default: "1.1.1.1"
- `dns_resolve_type` (String) Default: "A"
- `docker_container` (String) Default: ""
- `docker_host` (Number)
- `expiry_notification` (Boolean) Default: false
- `headers` (String)
- `hostname` (String)
- `ignore_tls` (Boolean) Default: false
- `interval` (Number) Default: 60
- `keyword` (String)
- `maxredirects` (Number) Default: 10
- `maxretries` (Number) Default: 0
- `method` (String) Default: "GET"
- `mqtt_password` (String)
- `mqtt_success_message` (String)
- `mqtt_topic` (String)
- `mqtt_username` (String)
- `notification_id_list` (List of String)
- `port` (Number) Default: 53
- `proxy_id` (Number)
- `radius_called_station_id` (String)
- `radius_calling_station_id` (String)
- `radius_password` (String)
- `radius_secret` (String)
- `radius_username` (String)
- `resend_interval` (Number) Default: 0
- `retry_interval` (Number) Default: 60
- `upside_down` (Boolean) Default: false
- `url` (String)

### Read-Only

- `data` (String) Parsed as JSON.
