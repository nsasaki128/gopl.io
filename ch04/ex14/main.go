package main

import (
	"html/template"
	"log"

	"net/http"

	"gopl.io/ch04/ex14/github"
)

type UserInformation struct {
	Issues         []*github.Issue
	Milestones     []*github.Milestone
	FollowerUsers  []*github.User
	FollowingUsers []*github.User
}

var userInformationList = template.Must(template.New("userInformation").Parse(`
<h1>nsasaki128 information</h1>
<h2>Issues</h2>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Issues}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
<h2>Milestones</h2>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>Title</th>
</tr>
{{range .Milestones}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
<h2>Follower Users</h2>
<table>
<tr style='text-align: left'>
  <th>Name</th>
</tr>
{{range .FollowerUsers}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Login}}</a></td>
</tr>
{{end}}
</table>
<h2>Following Users</h2>
<table>
<tr style='text-align: left'>
  <th>Name</th>
</tr>
{{range .FollowingUsers}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Login}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	userInformation := getUserInformation()
	if err := userInformationList.Execute(w, userInformation); err != nil {
		log.Fatal(err)
	}
}

func getUserInformation() UserInformation {
	issues, _ := github.SearchIssues()
	milestones, _ := github.SearchMilestones()
	followingUsers, _ := github.SearchFollowingUsers()
	followerUsers, _ := github.SearchFollowerUsers()
	return UserInformation{issues, milestones, followerUsers, followingUsers}
}
