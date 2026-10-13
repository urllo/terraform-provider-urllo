output "host_names" {
  description = "Every source host on the account."
  value       = [for h in data.urllo_hosts.all.hosts : h.name]
}

output "rule_count" {
  description = "Number of redirect rules on the account."
  value       = length(data.urllo_rules.all.rules)
}

# output "custom_404_body_present" {
#   description = "Optional custom 404 body contentpresent"
#   value       = urllo_host.custom_404.not_found_action.custom_404_body_present
# }

output "urllo_rule_example" {
  description = "Full output from example"
  value       = urllo_rule.example
}

output "example_requests_past_week" {
  description = "Requests processed by the example rule over the past week."
  value       = one(data.urllo_rules.example_analytics.rules).analytics.requests_processed
}
