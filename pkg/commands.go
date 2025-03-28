package pkg

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	CommandRepo.Flags().StringP("url", "u", "", "url the repo.(eg: https://github.com/informeai/gia)")
	CommandRepo.Flags().BoolP("local", "l", false, "get repo from actual directory.(verify .git folder)")
	CommandRepo.Flags().BoolP("authors", "a", false, "List all authors the repo")
}

var CommandVersion = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long:  `Print the version the gia cli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version - %s\n", VERSION)
	},
}

var CommandRepo = &cobra.Command{
	Use:   "repo",
	Short: "repo commands",
	Long:  `execute query and get informations from repo`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var url string
		urlFlag, err := cmd.Flags().GetString("url")
		if err != nil {
			return err
		}
		// TODO: add validator
		if len(urlFlag) > 0 {
			url = urlFlag
		}

		localFlag, err := cmd.Flags().GetBool("local")
		if err != nil {
			return err
		}
		if localFlag {
			actualDir, err := os.Getwd()
			if err != nil {
				return err
			}
			urlLocal := fmt.Sprintf("file://%s", actualDir)
			url = urlLocal
		}

		gitWrapper := NewGitWrapper()
		if err := gitWrapper.Init(url); err != nil {
			return err
		}
		authorsFlag, err := cmd.Flags().GetBool("authors")
		if err != nil {
			return err
		}
		if authorsFlag {
			authors, err := gitWrapper.Authors()
			if err != nil {
				return err
			}
			fmt.Printf("Authors:\n| commits_quant | name | email |\n")
			for _, author := range authors {
				fmt.Printf("%d\t%s\t%s\n", author.CommitCount, author.Name, author.Email)
			}
			fmt.Printf("\nTotal Authors: %d\n", len(authors))
		}
		return nil
	},
}
