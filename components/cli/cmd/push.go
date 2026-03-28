package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push the button on the Useless Machine",
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := getBackend()
		if err != nil {
			return err
		}

		result, err := b.PushButton(context.Background())
		if err != nil {
			return err
		}

		fmt.Println(result)
		return nil
	},
}
