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