package createproject

import "github.com/spf13/cobra"

// CLI is cobra command used for create new project.
func CLI(cmd *cobra.Command, args []string) {
	projectName := args[0]
	moduleName := args[1]

	CreateProject(projectName, moduleName)
}
