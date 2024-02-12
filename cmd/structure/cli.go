package structure

import "github.com/spf13/cobra"

// CLI is cobra command used for create new project.
func CLI(cmd *cobra.Command, args []string) {
	Structure()
}
