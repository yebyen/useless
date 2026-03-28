package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status of the Useless Machine",
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := getBackend()
		if err != nil {
			return err
		}

		status, err := b.GetStatus(context.Background())
		if err != nil {
			return err
		}

		fmt.Printf("Count: %d\n", status.DailyCount)
		fmt.Printf("Last Pushed: %s\n", status.LastPushed)
		fmt.Printf("Nagging: %v\n", status.IsNagging)
		if status.IsNagging {
			fmt.Println("Don't be lazy! Push it!")
		}
		return nil
	},
}
