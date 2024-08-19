package apps

import (
	"fmt"

	"github.com/spf13/cobra"
)

var FiberTemplCmd = &cobra.Command{
	Use:   "fibertempl",
	Short: "Create a Fiber Templ project structure",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a project name")
			return
		}
		projectName := args[0]
		createFiberProject(projectName)
	},
}
func createFiberTemplProject(projectName string) {
	dirs := []string{
		projectName,
		filepath.Join(projectName, "static"),
		filepath.Join(projectName, "templates"),
		filepath.Join(projectName, "static", "src"),
		filepath.Join(projectName, "static", "css"),
		filepath.Join(projectName, "static", "js"),
		filepath.Join(projectName, "static", "img"),
	}
}