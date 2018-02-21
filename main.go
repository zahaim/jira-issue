package main

import (
	"fmt"
	"os"

	"github.com/andygrunwald/go-jira"
)

func main() {
  var (
    jiraurl = os.Getenv("JIRA_URL")
    username = os.Getenv("USERNAME")
    password = os.Getenv("PASSWORD")
  )

	fmt.Println("Starting Jira issue handler...")

	jiraClient, err := jira.NewClient(nil, jiraurl)
	if err != nil {
		panic(err)
	}
	// base auth
	// jiraClient.Authentication.SetBasicAuth(username, password)

	// sessioncookie auth
	res, err := jiraClient.Authentication.AcquireSessionCookie(username, password)
	if err != nil || res == false {
		fmt.Printf("Result: %v\n", res)
		panic(err)
	}

	// // for existing issues
	issue, _, err := jiraClient.Issue.Get("DC-681", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)

	// // creating issues
	// i := jira.Issue{
	// 	Fields: &jira.IssueFields{
	// 		Description: "Please deploy this app on PROD",
	// 		Type: jira.IssueType{
	// 			Name: "Task",
	// 		},
	// 		Project: jira.Project{
	// 			Key: "DC",
	// 		},
	// 		Summary: "New deployment of some app",
	// 	},
	// }
	// newissue, _, err := jiraClient.Issue.Create(&i)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Printf("%s\n", newissue.Key)
	//
	// // commenting issues
	// // lenght := len(issue.Fields.Comments.Comments)
	// // given := issue.Fields.Comments.Comments[3].Body
	// //
	// // fmt.Printf("%s: %+v, %v, %v, %d, %v\n\n", issue.Key, issue.Fields.Summary, issue.Fields.Status.Name, issue.Fields.Comments.Comments, lenght, given)
	//
	// commentissue := newissue.Key
	// c := &jira.Comment{Body: "deployed on PREPROD"}
	// comment, _, err := jiraClient.Issue.AddComment(commentissue, c)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%v\n", comment.Body)

	// end of story
	fmt.Println("Terminating Jira issue handler...")
}
