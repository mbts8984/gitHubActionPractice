package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v37/github"
	"golang.org/x/oauth2"
)

type Package struct {
	FullName      string
	LastUpdatedBy string
}

var namedThing = github.RawOptions{Type: 1}

// var commentStruct = &github.PullRequestComment{Body: github.String(s"this comment is from local")}
var thingNotToInclude = thisIsntASecretButJustLook

//adding a ccomment for testing purposes

// var comment = github.PullRequestListCommentsOptions

func main() {
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	client := github.NewClient(tokenClient)
	//first test
	repo, _, err := client.Repositories.Get(context, "mbts8984", "gitHubActionPractice")

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	pack := &Package{
		FullName: *repo.FullName,
	}

	fmt.Printf("%+v\n", pack)
	// second test
	// commitInfo, _, err := client.Repositories.ListCommits(context, "mbts8984", "gitHubActionPractice", nil)

	// if err != nil {
	// 	fmt.Printf("Problem in commit information %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("%+v\n", commitInfo[0]) // Last commit information

	// third test
	// prInfo, _, err := client.PullRequests.Get(context, "mbts8984", "gitHubActionPractice", 4)
	// fmt.Printf("LOOK HERE HUMAN", prInfo)

	"fourth test"
	// getPrDifUrl, _, err := client.PullRequests.GetReview(context, "mbts8984", "gitHubActionPractice", 4, 723664571)
	// fmt.Printf("%v+v\n", getPrDifUrl)

	//fifth test
	// getDiffUrl, _, err := client.PullRequest.getDiffUrl()
	// thing := prInfo.GetDiffURL()
	// fmt.Printf(thing)
	// TEST

	getRawStuff, _, err := client.PullRequests.GetRaw(context, "mbts8984", "gitHubActionPractice", 2, namedThing)
	fmt.Printf("RAW GROSS", getRawStuff)
	if err != nil {
		fmt.Printf("Problem in commit information %v\n", err)
		os.Exit(1)
	}
	// data, err := ioutil.ReadFile("one_file.patch")
	// lineNumber := 1
	// file, err := os.Open("one_file.patch")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// fmt.Printf("File Contents: %s", data)

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	// fmt.Println("LINE ", scanner.Text())
	// 	if strings.Contains(scanner.Text(), "+") {
	// 		fmt.Println("Line", lineNumber)
	// 	}
	// 	lineNumber++
	// }
	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// makeAComment, _, err := client.PullRequests.CreateComment(context, "mbts8984", "gitHubActionPractice", 4, commentStruct)
	// if err != nil {
	// 	fmt.Printf("Problem in commit information %v\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Print("comment: ", makeAComment)
}
