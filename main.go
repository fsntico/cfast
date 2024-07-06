package main

import (
	"CreateApp/apps"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "create"}
	rootCmd.AddCommand(apps.FlaskCmd)
	rootCmd.AddCommand(apps.FiberCmd)
	rootCmd.Execute()
}

// go build -o cpjt.exe .
