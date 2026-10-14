resource "urllo_rule" "example" {
  source_urls    = ["source.urllo-acceptance-test.com/"]
  target_url     = "https://target.urllo-acceptance-test.com"
  validate_dns   = false
  forward_params = true
  forward_path   = true
  tags           = ["tag 1", "tag 2"]
}
