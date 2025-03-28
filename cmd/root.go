package cmd

import (
	"os"

	"github.com/informeai/gia/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pkg.CommandVersion)
	rootCmd.AddCommand(pkg.CommandRepo)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gia",
	Short: "A git information analizer",
	Long:  `A git cli for information e analize repos git.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
