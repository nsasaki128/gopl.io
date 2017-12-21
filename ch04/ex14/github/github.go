package github

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const BaseRepoURL = "https://api.github.com/repos/nsasaki128/gopl.io/"
const IssueURL = BaseRepoURL + "issues"
const MilestoneURL = BaseRepoURL + "milestones"

const BaseUserURL = "https://api.github.com/users/nsasaki128/"
const FollowerUserURL = BaseUserURL + "followers"
const FollowingUserURL = BaseUserURL + "following"

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at""`
	Body      string
}

type Milestone struct {
	URL          string
	HTMLURL      string `json:"html_url"`
	LabelsURL    string `json:"labels_url"`
	Number       int
	State        string
	Title        string
	Description  string
	Creator      *User
	OpenIssues   int       `json:"open_issues"`
	ClosedIssues int       `json:"closed_issues"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DueOn        time.Time `json:"due_on"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"'`
}

func SearchIssues() ([]*Issue, error) {
	var results []*Issue
	if err := request("GET", nil, IssueURL, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func SearchMilestones() ([]*Milestone, error) {
	var results []*Milestone
	if err := request("GET", nil, MilestoneURL, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func SearchFollowingUsers() ([]*User, error) {
	var results []*User
	if err := request("GET", nil, FollowingUserURL, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func SearchFollowerUsers() ([]*User, error) {
	var results []*User
	if err := request("GET", nil, FollowerUserURL, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func request(method string, param interface{}, url string, result interface{}) error {
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

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}
	return nil
}
