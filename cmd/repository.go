/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/fatih/color"
	"github.com/felipem1210/git-helper/githelper"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Create/update a repository",
	Long:  `Create/update a repository in a Github Organization.`,
	Run: func(cmd *cobra.Command, args []string) {
		provider, _ := cmd.Flags().GetString("provider")
		new_repo_json_file, _ := cmd.Flags().GetString("new-repo-json-file")
		repo_info_json_file, _ := cmd.Flags().GetString("repo-info-json-file")
		create, _ := cmd.Flags().GetBool("create")

		if provider == "github" {
			if create {
				myRepos := githelper.MyRepos{}
				myRepos = myRepos.GithubCreateRepos(new_repo_json_file)
				err := githelper.WriteReposToJson(myRepos, repo_info_json_file)
				if err != nil {
					githelper.CheckIfError(err)
				} else {
					color.Green("The json file %s with repo info was written sucessfully", repo_info_json_file)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
	repoCmd.PersistentFlags().String("new-repo-json-file", "new_repos.json", "The json file needed to create new repositories.")
	repoCmd.PersistentFlags().StringP("provider", "p", "", "A provider to choose, options: gitub, gitlab")
	repoCmd.MarkPersistentFlagRequired("provider")
	repoCmd.PersistentFlags().BoolP("create", "c", false, "Create the repositories from repo-info-json-file")
}
