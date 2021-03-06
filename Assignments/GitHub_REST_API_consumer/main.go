package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

//sub-struct in ReposInfoJson
type Owner struct {
	Login               string
	Id                  int
	Node_id             string
	Avatar_url          string
	Gravatar_id         string
	Url                 string
	Html_url            string
	Followers_url       string
	Following_url       string
	Gists_url           string
	Starred_url         string
	Subscriptions_url   string
	Organizations_url   string
	Repos_url           string
	Events_url          string
	Received_events_url string
	Type                string
	Site_admin          bool
}

//sub-struct in ReposInfoJson
type Licence struct {
	Key     string
	Name    string
	Spdx_id string
	Url     string
	Node_id string
}

type Userinfo struct {
	Login               string
	Id                  int
	Node_id             string
	Avatar_url          string
	Gravatar_id         string
	Url                 string
	Html_url            string
	Followers_url       string
	Following_url       string
	Gists_url           string
	Starred_url         string
	Subscriptions_url   string
	Organizations_url   string
	Repos_url           string
	Events_url          string
	Received_events_url string
	Type                string
	Site_admin          bool
	Name                string
	Company             string
	Blog                string
	Location            string
	Email               string
	Hireable            bool
	Bio                 string
	Twitter_username    string
	Public_repos        int
	Public_gists        int
	Followers           int
	Following           int
	Created_at          string
	Updated_at          string
}

type ReposInfoJson struct {
	Id                int
	Node_id           string
	Name              string
	Full_name         string
	Private           bool
	Owner             Owner
	Html_url          string
	Description       string
	Fork              bool
	Url               string
	Forks_url         string
	Keys_url          string
	Collaborators_url string
	Teams_url         string
	Hooks_url         string
	Issue_events_url  string
	Events_url        string
	Assignees_url     string
	Branches_url      string
	Tags_url          string
	Blobs_url         string
	Git_tags_url      string
	Git_refs_url      string
	Trees_url         string
	Statuses_url      string
	Languages_url     string
	Stargazers_url    string
	Contributors_url  string
	Subscribers_url   string
	Subscription_url  string
	Commits_url       string
	Git_commits_url   string
	Comments_url      string
	Issue_comment_url string
	Contents_url      string
	Compare_url       string
	Merges_url        string
	Archive_url       string
	Downloads_url     string
	Issues_url        string
	Pulls_url         string
	Milestones_url    string
	Notifications_url string
	Labels_url        string
	Releases_url      string
	Deployments_url   string
	Created_at        string
	Updated_at        string
	Pushed_at         string
	Git_url           string
	Ssh_url           string
	Clone_url         string
	Svn_url           string
	Homepage          string
	Size              int
	Stargazers_count  int
	Watchers_count    int
	Language          string
	Has_issues        bool
	Has_projects      bool
	Has_downloads     bool
	Has_wiki          bool
	Has_pages         bool
	Forks_count       int
	Mirror_url        string
	Archived          bool
	Disabled          bool
	Open_issues_count int
	License           Licence
	Forks             int
	Open_issues       int
	Watchers          int
	Default_branch    string
}

type ReposInfoArray []ReposInfoJson

func main() {
	apiUsername := getUsername()
	reposLoop := getInput(apiUsername)
	getRepos(reposLoop, apiUsername)
}

func getUsername() string {
	var username string
	fmt.Printf("Enter the username of the github user : ")
	fmt.Scanln(&username)
	return username
}

func getInput(apiUsername string) int {

	url := "https://api.github.com/users/" + apiUsername

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	//Send request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	//Read the response body.
	bodyJson, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var users Userinfo
	json.Unmarshal([]byte(bodyJson), &users)

	if users.Name != "" {

		userDataStr := []string{"Name               : " + users.Name + "\n\nUsername           : " + users.Login + "\n\nBio                : " + users.Bio + "\nPublic Repositories: " + strconv.Itoa(users.Public_repos) + "\n\nFollowers          : " + strconv.Itoa(users.Followers) + "\n\nFollowing          : " + strconv.Itoa(users.Following)}

		file, err := os.OpenFile(apiUsername+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}

		defer file.Close()

		datawriter := bufio.NewWriter(file)

		for _, data := range userDataStr {
			_, _ = datawriter.WriteString(data + "\n")
		}

		datawriter.Flush()

	} else {
		fmt.Println("GitHub user [" + apiUsername + "] not found!!")
	}
	return users.Public_repos
}

func getRepos(reposLoop int, apiUsername string) {

	url := "http://api.github.com/users/" + apiUsername + "/repos?per_page=100"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	bodyJson, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var reposArray ReposInfoArray
	json.Unmarshal([]byte(bodyJson), &reposArray)

	var loopCondition int
	if reposLoop >= 100 {
		loopCondition = 100
	} else {
		loopCondition = reposLoop
	}

	for i := 0; i < loopCondition; i++ {

		repoDataStr := []string{"\nRepository No[" + strconv.Itoa(i+1) + "]: " + reposArray[i].Name + ". \nAvailable at    : " + reposArray[i].Html_url}

		file, err := os.OpenFile(apiUsername+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		defer file.Close()

		datawriter := bufio.NewWriter(file)

		for _, data := range repoDataStr {
			_, _ = datawriter.WriteString(data + "\n")
		}

		datawriter.Flush()
	}
}
