package createmodule

import "github.com/spf13/cobra"

// CLI is cobra command used for create new module.
func CLI(cmd *cobra.Command, args []string) {
	module := "hello world"

	if len(args) > 0 {
		module = args[0]
	}

	CreateModule(module)
}
