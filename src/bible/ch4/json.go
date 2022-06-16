package ch4

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

type TestJsonStruct struct {
	Page int `json:"page"`
	Fruits []string
}

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items          []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}


func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func Practice410(TimeInterval int64) {
	issuesStrings := []string{
		"repo:golang/go is:open json decoder",
	}
	result, err := SearchIssues(issuesStrings)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(result)
	fmt.Printf("%d issues:\n", result.TotalCount)

	t := time.Now()
	Tutc := t.Unix()
	for _, item := range result.Items {
		// convert item createdt to unix
		itemUnix := item.CreatedAt.Unix()
		if Tutc - itemUnix < TimeInterval{
			fmt.Printf("#%-5d %9.9s %.55s %v \n",
				item.Number, item.User.Login, item.Title, item.CreatedAt.UTC())
		}
	}
}
