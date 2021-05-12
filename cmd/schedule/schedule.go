package schedule

import (
	"github.com/kutifyhq/autifyctl/pkg/option"
	"github.com/spf13/cobra"
)

type scheduleOption struct {
	rootOpt *option.RootOption
}

func NewScheduleCommand(rootOpt *option.RootOption) *cobra.Command {
	opt := &scheduleOption{
		rootOpt: rootOpt,
	}

	cmd := &cobra.Command{
		Use:     "schedule",
		Aliases: []string{"sch"},
	}

	cmd.AddCommand(newRunCommand(opt))
	return cmd
}
