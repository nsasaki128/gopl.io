package main

import (
	"fmt"
	"log"

	"gopl.io/ch04/ex11/github"
)

func main() {
	search()

}

func search() {
	result, err := github.SearchIssues()
	if err != nil {
		log.Fatal(err)
	}
	for _, issue := range result {
		printIssue(issue)
	}
}

func create(title, body string) {
	issue, err := github.CreateIssue(&github.IssueCreate{Title: title, Body: body}, "dummy-token")
	if err != nil {
		log.Fatal(err)
	}
	printIssue(issue)
}

func printIssue(issue *github.Issue) {
	fmt.Printf("#%-5d %9.9s %.55s\n",
		issue.Number, issue.User.Login, issue.Title)
}
