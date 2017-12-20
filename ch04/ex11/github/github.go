package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const IssuesURL = "https://api.github.com/repos/nsasaki128/gopl.io/issues"

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at""`
	Body      string
}

type IssueCreate struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Assignee  string   `json:"assignee,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}

type IssueUpdate struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Assignee  string   `json:"assignee,omitempty"`
	State     string   `json:"state,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"'`
}

func SearchIssues() ([]*Issue, error) {
	resp, err := http.Get(IssuesURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var issues []*Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return issues, nil
}

func CreateIssue(param *IssueCreate, token string) (*Issue, error) {
	//POST /repos/:owner/:repo/issues
	var body io.Reader
	if param != nil {
		json, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(json)
	}

	req, err := http.NewRequest("POST", IssuesURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+token)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var issue *Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}

	return issue, nil
}

func UpdateIssue(number string) {
	//PATCH /repos/:owner/:repo/issues/:number

}
