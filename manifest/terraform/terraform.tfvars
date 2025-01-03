project_id     = "glowing-service-444205-v1"
region         = "us-central1"
service_name   = "middleware-webhook"
image          = "fitraelbi/middleware-webhook:latest"
env_variables  = {
  GITHUB_OWNER = "fitraelbi"
  GITHUB_REPO     = "middleware-webhook"
  GITHUB_TOKEN    = "custom-token"
}
