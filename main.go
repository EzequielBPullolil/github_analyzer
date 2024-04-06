package main

import (
	"errors"
	"fmt"

	"github.com/EzequielBPullolil/github_analyzer/colors"
	profileanalyzer "github.com/EzequielBPullolil/github_analyzer/profile_analyzer"
	"github.com/spf13/cobra"
)

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

		if len(args) == 1 {
			username := args[0]
			fmt.Printf("Scoring username @%s..... \n", colors.Magenta(username))
			profileanalyzer.AnalyzeProfile(username)
		}
	},
}

func main() {

	mainCMD.Execute()
}
