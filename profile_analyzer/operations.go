package profileanalyzer

import "fmt"

func searchAllRepositories() []Repository {
	var repos []Repository

	return repos
}

func AnalyzeProfile(username string) {
	repos := searchAllRepositories()

	fmt.Println(repos)
}
