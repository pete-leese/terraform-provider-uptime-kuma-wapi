---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ukumawapi_maintenance Resource - terraform-provider-ukumawapi"
subcategory: ""
description: |-
  Maintenance Resource
---

# ukumawapi_maintenance (Resource)

Maintenance Resource

## Example Usage

```terraform
resource "ukumawapi_maintenance" "my_maintenance" {
  active = false
  date_range = [
    "{ \"see\": \"documentation\" }"
  ]
  days_of_month = [
    "{ \"see\": \"documentation\" }"
  ]
  description    = "...my_description..."
  interval_day   = 6
  maintenance_id = 8
  strategy       = "recurring-interval"
  time_range = [
    "{ \"see\": \"documentation\" }"
  ]
  title = "...my_title..."
  weekdays = [
    "{ \"see\": \"documentation\" }"
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `maintenance_id` (Number)
- `strategy` (String) Enumerate maintenance strategies. must be one of ["manual", "single", "recurring-interval", "recurring-weekday", "recurring-day-of-month", "cron"]; Requires replacement if changed.
- `title` (String) Requires replacement if changed.

### Optional

- `active` (Boolean) Default: true; Requires replacement if changed.
- `date_range` (List of String) Requires replacement if changed.
- `days_of_month` (List of String) Requires replacement if changed.
- `description` (String) Default: ""; Requires replacement if changed.
- `interval_day` (Number) Default: 1; Requires replacement if changed.
- `time_range` (List of String) Requires replacement if changed.
- `weekdays` (List of String) Requires replacement if changed.

### Read-Only

- `data` (String) Parsed as JSON.