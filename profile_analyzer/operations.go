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

func AnalyzePublicRepos(rs []Repository) (int, int, int) {
	fmt.Println("Calculate public repos score....")
	total_commits := 0
	total_no_readme_repos := 0
	for _, repo := range rs {
		if !repo.HaveReadme() {
			total_no_readme_repos += 1
		}

		total_commits += repo.CalculateCommits()

	}

	fmt.Println("Public repos score calculated")
	fmt.Println("-----------------------------")
	return len(rs), total_commits, total_no_readme_repos
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

	cant_repos, total_commits, total_no_readme_repos := AnalyzePublicRepos(repos)
	c_average := total_commits / cant_repos
	fmt.Printf("Repositories without readme: %d/%d \n", total_no_readme_repos, cant_repos)
	fmt.Printf("Commit average: %d \n", c_average)
	fmt.Println("Empty repositories: 10")
	fmt.Printf("Profile Readme: %s \n", ProfileHaveReadme(username))
	fmt.Println(colors.Info("------------End------------"))
}
