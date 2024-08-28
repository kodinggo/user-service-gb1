package console

import (
	"os"

	"github.com/kodinggo/user-service-gb1/internal/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Kodinggo User Service GB1",
	Short: "Kodinggo User Service GB1",
	Long:  `User Service Golang Bootcamp 1 Kodinggo.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	config.SetupLogger()
}
