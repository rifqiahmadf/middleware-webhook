package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	gitHubAPIURL = "https://api.github.com/repos/%s/statuses/%s"
)

type GitLabPipelineWebhook struct {
	ObjectAttributes struct {
		ID     int    `json:"id"`
		Ref    string `json:"ref"`
		SHA    string `json:"sha"`
		Status string `json:"status"`
	} `json:"object_attributes"`
}

type GitHubStatusRequest struct {
	State       string `json:"state"`
	Description string `json:"description"`
	Context     string `json:"context"`
}

func mapGitLabStatusToGitHubState(status string) string {
	switch status {
	case "pending":
		return "pending"
	case "running":
		return "pending"
	case "success":
		return "success"
	case "failed":
		return "failure"
	case "canceled":
		return "error"
	default:
		return "error"
	}
}

func updateGitHubStatus(owner, repo, sha, token, state, description string) error {
	url := fmt.Sprintf(gitHubAPIURL, owner+"/"+repo, sha)
	
	statusRequest := GitHubStatusRequest{
		State:       state,
		Description: description,
		Context:     "GitLab Pipeline",
	}

	body, err := json.Marshal(statusRequest)
	if err != nil {
		return fmt.Errorf("failed to marshal status request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		responseBody, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("unexpected response: %s", string(responseBody))
	}

	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var webhook GitLabPipelineWebhook

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &webhook)
	if err != nil {
		http.Error(w, "invalid JSON payload", http.StatusBadRequest)
		return
	}

	state := mapGitLabStatusToGitHubState(webhook.ObjectAttributes.Status)
	owner := os.Getenv("GITHUB_OWNER")
	repo := os.Getenv("GITHUB_REPO")
	token := os.Getenv("GITHUB_TOKEN")

	err = updateGitHubStatus(owner, repo, webhook.ObjectAttributes.SHA, token, state, "Pipeline status from GitLab")
	if err != nil {
		log.Printf("failed to update GitHub status: %v", err)
		http.Error(w, "failed to update GitHub status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/webhook", handler)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
