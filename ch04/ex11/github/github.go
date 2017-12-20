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
	var issue Issue
	if err := requestIssue("POST", param, token, IssuesURL, &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func UpdateIssue(number string, param *IssueUpdate, token string) (*Issue, error) {
	//PATCH /repos/:owner/:repo/issues/:number
	var issue Issue
	if err := requestIssue("PATCH", param, token, IssuesURL+"/"+number, &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func requestIssue(method string, param interface{}, token string, url string, issue *Issue) error {
	var body io.Reader
	if param != nil {
		json, err := json.Marshal(param)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(json)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "token "+token)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return err
	}
	return nil
}
