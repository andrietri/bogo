package main

import (
	"log"

	"github.com/andrietri/bogo/cmd/createmodule"
	"github.com/andrietri/bogo/cmd/createproject"
	"github.com/andrietri/bogo/cmd/structure"
	"github.com/andrietri/bogo/cmd/version"
	"github.com/spf13/cobra"
)

// Root cobra CLI.
func main() {
	rootCmd := &cobra.Command{
		Use:               "bogo",
		Short:             "Golang CLI generate project",
		Long:              "Golang CLI generate project.",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	rootCmd.AddCommand(getCreateModuleCmd())
	rootCmd.AddCommand(getCreateProjectCmd())
	rootCmd.AddCommand(getVersionCmd())
	rootCmd.AddCommand(getStructureCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func getCreateModuleCmd() *cobra.Command {
	createModuleCmd := &cobra.Command{
		Use:   "create-module [module]",
		Short: "Create new module with the provided module name",
		Long:  "Create new module with the provided module name. If no module name is provided, it defaults to 'hello world'.",
		Args:  cobra.MaximumNArgs(1),
		Run:   createmodule.CLI,
	}
	return createModuleCmd
}

func getVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show bogo version",
		Long:  "Show bogo version.",
		Args:  cobra.MaximumNArgs(0),
		Run:   version.CLI,
	}
	return versionCmd
}

func getStructureCmd() *cobra.Command {
	structureCmd := &cobra.Command{
		Use:   "structure",
		Short: "Explain code structure",
		Long:  "Explain code structure.",
		Args:  cobra.MaximumNArgs(0),
		Run:   structure.CLI,
	}
	return structureCmd
}

func getCreateProjectCmd() *cobra.Command {
	createProjectCmd := &cobra.Command{
		Use:   "create-project [project-name] [module-name]",
		Short: "Create new project with the provided module name",
		Long:  "Create new project with the provided module name.",
		Args:  cobra.MinimumNArgs(2),
		Run:   createproject.CLI,
	}
	return createProjectCmd
}
