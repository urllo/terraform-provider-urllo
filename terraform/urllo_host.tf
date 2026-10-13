# Does not require an import as the host already exists in urllo
resource "urllo_host" "example_302" {
  name = trim(one(urllo_rule.example.source_urls), "/") # urllo normalizes the hostname with a trailing /

  acme_enabled = true

  match_options = {
    case_insensitive  = true
    slash_insensitive = true
  }

  not_found_action = {
    forward_params = true
    forward_path   = true
    response_code  = 302
    response_url   = "https://www.example.com"
  }

  # not_found_action = {
  #   response_code = 404
  # }

  # custom_404_body = <<-HTML
  #   <!doctype html>
  #   <html>
  #     <head><title>Page not found</title></head>
  #     <body><h1>404 - Page not found</h1></body>
  #   </html>
  # HTML

  # security = {
  #   https_upgrade             = true
  #   prevent_foreign_embedding = true
  #   hsts_include_sub_domains  = true
  #   hsts_max_age              = 31536000
  #   hsts_preload              = true
  # }
}
