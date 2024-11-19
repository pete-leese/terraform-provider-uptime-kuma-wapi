resource "ukumawapi_maintenance" "my_maintenance" {
  active = false
  date_range = [
    "{ \"see\": \"documentation\" }"
  ]
  days_of_month = [
    "{ \"see\": \"documentation\" }"
  ]
  description  = "...my_description..."
  interval_day = 6
  strategy     = "recurring-interval"
  time_range = [
    "{ \"see\": \"documentation\" }"
  ]
  title = "...my_title..."
  weekdays = [
    "{ \"see\": \"documentation\" }"
  ]
}