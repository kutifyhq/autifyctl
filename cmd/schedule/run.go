package schedule

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/kutifyhq/autifyctl/pkg/client"
	"github.com/spf13/cobra"
)

type runOption struct {
	accessToken string
}

func newRunCommand(so *scheduleOption) *cobra.Command {
	opts := &runOption{}

	cmd := &cobra.Command{
		Use:     "run TestPlanID(s)",
		Aliases: []string{"r"},
		Long: `autifyctl schedule run TestPlanID(s) runs TestPlan(s) by schedule ID as below:

	autifyctl schedule run 1234 --access-token $ACCESS_TOKEN

You can also specify mulitple schedule IDs with comma seperator as below:

	autifyctl schedule run 1,2,3,4 --access-token $ACCESS_TOKEN
`,
		Args: cobra.ExactArgs(1),
		RunE: scheduleRunRun(opts),
	}

	cmd.PersistentFlags().StringVar(&opts.accessToken, "access-token", os.Getenv("AUTIFYCTL_ACCESS_TOKEN"), "Access token to access")
	return cmd
}

func scheduleRunRun(opt *runOption) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		schIDStrs := args[0]
		schIDs := []int{}

		for _, schIDStr := range strings.Split(schIDStrs, ",") {
			schID, err := strconv.Atoi(schIDStr)
			if err != nil {
				return fmt.Errorf("unable to parse schedule ID to int, type:%T", schIDStr)
			}

			schIDs = append(schIDs, schID)
		}

		c, err := client.NewClient(opt.accessToken)
		if err != nil {
			return err
		}

		for _, schID := range schIDs {
			resp, err := c.PostSchedulesScheduleIdWithResponse(ctx, schID)
			if err != nil {
				return err
			}

			fmt.Println(schID)

			if resp.StatusCode() != http.StatusOK {
				apiErrors, err := client.ParseError(resp.Body)
				if err != nil {
					return err
				}

				return fmt.Errorf("API response status code was not 200 got %d with messages: %v", resp.StatusCode(), apiErrors.Errors)
			}

			schedule := (*resp.JSON200)[0]
			fmt.Printf("schedule started with id:%p", schedule.Id)
		}

		return nil
	}
}
