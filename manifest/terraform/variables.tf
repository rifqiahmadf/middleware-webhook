variable "project_id" {}
variable "region" {
  default = "asia-southeast2"
}
variable "service_name" {
  default = "middleware-webhook"
}
variable "image" {
  default = "fitrakz/middleware-webhook:latest"
}
variable "env_variables" {
  type = map(string)
}