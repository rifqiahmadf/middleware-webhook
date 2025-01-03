

provider "google" {
  project = var.project_id
  region  = var.region
}

resource "google_cloud_run_service" "hello_world" {
  name     = var.service_name
  location = var.region

  template {
    spec {
      containers {
        image = var.image

        dynamic "env" {
          for_each = var.env_variables
          content {
            name  = env.key
            value = env.value
          }
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

resource "google_project_iam_binding" "run_invoker" {
  project = var.project_id
  role    = "roles/run.invoker"

  members = [
    "allUsers"
  ]
}
