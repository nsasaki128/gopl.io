package main

import (
	"fmt"
	"log"
	"os"

	"time"

	"gopl.io/ch04/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	now := time.Now()
	monthAgo := now.AddDate(0, -1, 0)
	yearAgo := now.AddDate(-1, 0, 0)
	fmt.Println("more than a year old")
	for _, item := range result.Items {
		if !item.CreatedAt.After(yearAgo) {
			fmt.Printf("#%-5d %9.9s %.55s %v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}
	fmt.Println("less than a year old")
	for _, item := range result.Items {
		if item.CreatedAt.After(yearAgo) && !item.CreatedAt.After(monthAgo) {
			fmt.Printf("#%-5d %9.9s %.55s %v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}
	fmt.Println("less than a month old")
	for _, item := range result.Items {
		if item.CreatedAt.After(monthAgo) {
			fmt.Printf("#%-5d %9.9s %.55s %v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}

}
