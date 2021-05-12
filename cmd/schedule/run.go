package schedule

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/kutifyhq/autifyctl/pkg/client"
	"github.com/spf13/cobra"
)

type runOption struct {
	accessToken string
}

func newRunCommand(so *scheduleOption) *cobra.Command {
	opts := &runOption{}

	cmd := &cobra.Command{
		Use:     "run",
		Aliases: []string{"r"},
		Args:    cobra.ExactArgs(1),
		RunE:    scheduleRunRun(opts),
	}

	cmd.PersistentFlags().StringVar(&opts.accessToken, "access-token", os.Getenv("AUTIFYCTL_ACCESS_TOKEN"), "Access token to access")
	return cmd
}

func scheduleRunRun(opt *runOption) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		schIDStr := args[0]
		schID, err := strconv.Atoi(schIDStr)
		if err != nil {
			return fmt.Errorf("unable to parse schedule ID to int, type:%T", schIDStr)
		}

		c, err := client.NewClient(opt.accessToken)
		if err != nil {
			return err
		}

		resp, err := c.PostSchedulesScheduleIdWithResponse(ctx, schID)
		if err != nil {
			return err
		}

		if resp.StatusCode() != http.StatusOK {
			apiErrors, err := client.ParseError(resp.Body)
			if err != nil {
				return err
			}

			return fmt.Errorf("API response status code was not 200 got %d with messages: %v", resp.StatusCode(), apiErrors.Errors)
		}

		schedule := (*resp.JSON200)[0]
		fmt.Printf("schedule started with id:%p", schedule.Id)
		return nil
	}
}
