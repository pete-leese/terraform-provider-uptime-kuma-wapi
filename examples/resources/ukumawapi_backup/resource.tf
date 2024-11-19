resource "ukumawapi_backup" "my_backup" {
  import_handle = "overwrite"
  monitor_list = [
    {
      accepted_statuscodes = [
        "..."
      ]
      active                     = 1
      auth_domain                = "{ \"see\": \"documentation\" }"
      auth_method                = "...my_auth_method..."
      auth_workstation           = "{ \"see\": \"documentation\" }"
      basic_auth_pass            = "{ \"see\": \"documentation\" }"
      basic_auth_user            = "{ \"see\": \"documentation\" }"
      body                       = "{ \"see\": \"documentation\" }"
      database_connection_string = "{ \"see\": \"documentation\" }"
      database_query             = "{ \"see\": \"documentation\" }"
      dns_last_result            = "{ \"see\": \"documentation\" }"
      dns_resolve_server         = "...my_dns_resolve_server..."
      dns_resolve_type           = "...my_dns_resolve_type..."
      docker_container           = "{ \"see\": \"documentation\" }"
      docker_host                = "{ \"see\": \"documentation\" }"
      expiry_notification        = true
      headers                    = "{ \"see\": \"documentation\" }"
      hostname                   = "{ \"see\": \"documentation\" }"
      id                         = 10
      ignore_tls                 = true
      interval                   = 4
      keyword                    = "{ \"see\": \"documentation\" }"
      maxredirects               = 4
      maxretries                 = 1
      method                     = "...my_method..."
      mqtt_password              = "{ \"see\": \"documentation\" }"
      mqtt_success_message       = "{ \"see\": \"documentation\" }"
      mqtt_topic                 = "{ \"see\": \"documentation\" }"
      mqtt_username              = "{ \"see\": \"documentation\" }"
      name                       = "...my_name..."
      notification_id_list = {
        # ...
      }
      port                      = 1
      proxy_id                  = "{ \"see\": \"documentation\" }"
      push_token                = "{ \"see\": \"documentation\" }"
      radius_called_station_id  = "{ \"see\": \"documentation\" }"
      radius_calling_station_id = "{ \"see\": \"documentation\" }"
      radius_password           = "{ \"see\": \"documentation\" }"
      radius_secret             = "{ \"see\": \"documentation\" }"
      radius_username           = "{ \"see\": \"documentation\" }"
      resend_interval           = 3
      retry_interval            = 10
      tags = [
        "{ \"see\": \"documentation\" }"
      ]
      type        = "...my_type..."
      upside_down = true
      url         = "...my_url..."
      weight      = 5
    }
  ]
  notification_list = [
    {
      active     = false
      config     = "...my_config..."
      id         = 5
      is_default = true
      name       = "...my_name..."
      user_id    = 9
    }
  ]
  proxy_list = [
    "{ \"see\": \"documentation\" }"
  ]
  version = "...my_version..."
}