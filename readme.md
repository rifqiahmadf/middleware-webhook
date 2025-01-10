# Middleware Webhook
Aplikasi converter pipeline event gitlab ke notifikasi commit github.

## Requirements
- Golang

## Installation
For apps
```bash
go mod download
go run main.go
```


## ENV APPS
Tambahkan pada variable gitlab ci
```bash
GITHUB_OWNER    => Username Github
GITHUB_REPO     => Repo name Github
GITHUB_TOKEN    => Personal Access Token Github
```
