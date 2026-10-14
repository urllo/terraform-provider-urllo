data "urllo_hosts" "all" {}

data "urllo_rules" "all" {}

# Request-volume analytics for the example rule over the past week.
# analytics_end_date is left unset -- the API defaults it to yesterday.
data "urllo_rules" "example_analytics" {
  analytics_start_date = formatdate("YYYY-MM-DD", timeadd(timestamp(), "-168h"))
  include_analytics    = true
  source_query         = tolist(urllo_rule.example.source_urls)[0]
}
