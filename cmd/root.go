package cmd

import (
	"github.com/kutifyhq/autifyctl/cmd/schedule"
	"github.com/kutifyhq/autifyctl/pkg/option"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "autifyctl",
	Short: "CLI tool for Autify API",
}

func init() {
	opt := &option.RootOption{}

	rootCmd.AddCommand(schedule.NewScheduleCommand(opt))
	rootCmd.AddCommand(newVersionCommand())
	rootCmd.AddCommand(newCompletionCmd())
}

func Execute() error {
	return rootCmd.Execute()
}
