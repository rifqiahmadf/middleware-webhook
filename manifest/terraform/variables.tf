variable "project_id" {}
variable "region" {
  default = "us-central1"
}
variable "service_name" {
  default = "hello-world-service"
}
variable "image" {
  default = "test/hello-world"
}
variable "env_variables" {
  type = map(string)
  default = {
    GITHUB_OWNER = "your-username"
    GITHUB_REPO     = "your-name"
    GITHUB_TOKEN    = "your-token"
  }
}