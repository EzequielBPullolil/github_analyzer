package profileanalyzer

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/EzequielBPullolil/github_analyzer/colors"
	"github.com/gocolly/colly"
)

func searchAllRepositories() []Repository {
	var repos []Repository
	state := colors.Green("found")
	c := colly.NewCollector()
	c.OnHTML("a[itemprop]", func(e *colly.HTMLElement) {
		var r Repository

		r.Name = strings.TrimSpace(e.Text)
		r.Url = fmt.Sprintf("https://github.com/%s", e.Attr("href"))
		repos = append(repos, r)
	})
	c.Visit("https://github.com/EzequielBPullolil?tab=repositories")
	if len(repos) == 0 {
		state = colors.Fail("not found")
	}
	fmt.Printf("Public repositories %s, quantity: %d \n", state, len(repos))
	return repos
}

func AnalyzePublicRepos(rs []Repository) (int, int, int, int) {

	total_commits := 0
	total_no_readme_repos := 0
	empty_repos := 0
	for _, repo := range rs {
		repo.GetData()
		total_commits += repo.CantCommits()
		if !repo.HaveReadme() {
			total_no_readme_repos += 1
		}

		if repo.IsEmpty() {
			empty_repos += 1
			fmt.Println(repo.Url)
		}
	}

	return len(rs), total_commits, total_no_readme_repos, empty_repos
}

func ProfileHaveReadme(username string) string {
	r, _ := http.Get(fmt.Sprintf("https://github.com/%s/%s", username, username))
	if r.StatusCode == 404 {
		return colors.Fail("False")
	}

	return colors.Green("True")
}

func AnalyzeProfile(username string) {
	fmt.Println(colors.Info("Finding public repos....."))
	repos := searchAllRepositories()
	fmt.Println("Analyzing public repositories....")
	cant_repos, total_commits, total_no_readme_repos, empty_repos := AnalyzePublicRepos(repos)
	fmt.Println("Public repositories analyzed")
	fmt.Println("-----------------------------")
	c_average := total_commits / cant_repos
	fmt.Printf("Repositories without readme: %d/%d \n", total_no_readme_repos, cant_repos)
	fmt.Printf("Commit average: %d \n", c_average)
	fmt.Printf("Empty repositories: %d \n", empty_repos)
	fmt.Printf("Profile Readme: %s \n", ProfileHaveReadme(username))
	fmt.Println(colors.Info("------------End------------"))
}
