package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd *cobra.Command

var _ = RegisterCommandVar(func() {
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "resume",
		Short: "Resume command line",
		Run: func(cmd *cobra.Command, args []string) {
			if path, _ := cmd.Flags().GetString("validate"); path != "" {
				os.Exit(ValidateResume(path))
			}
			// if there are no flags, the default will start the server
			StartResumeServer()
		},
	}
})

var _ = RegisterCommandInit(func() {
	rootCmd.PersistentFlags().String("validate", "", "Validate resume input file")
	// Global cli configurations
	rootCmd.Flags().SortFlags = false
})
