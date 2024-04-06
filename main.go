package main

import (
	"errors"

	"github.com/EzequielBPullolil/github_analyzer/colors"
	profileanalyzer "github.com/EzequielBPullolil/github_analyzer/profile_analyzer"
	"github.com/spf13/cobra"
)

var empty_repos_flag bool
var readme_repos_flag bool

func init() {
	mainCMD.PersistentFlags().BoolVarP(&empty_repos_flag, "empty-repos", "", false, "verbose output")
	mainCMD.PersistentFlags().BoolVarP(&readme_repos_flag, "readme-repos", "", false, "verbose output")

}

var mainCMD = &cobra.Command{
	Use:   "Github Analyzer",
	Short: "Create a score of your github",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("arg username required")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		if empty_repos_flag {
			cmd.Println("Finding empty repos")
			profileanalyzer.FindEmptyRepos(username)
		} else if readme_repos_flag {
			cmd.Println("Searching repositories without README")
			profileanalyzer.FindNoReadmeRepos(username)
		} else {
			cmd.Printf("Scoring username @%s..... \n", colors.Magenta(username))
			profileanalyzer.AnalyzeProfile(username)
		}
	},
}

func main() {
	mainCMD.Execute()
}
