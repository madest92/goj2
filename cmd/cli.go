// Package cli ...
package cli

import (
	"log"

	"github.com/madest92/goj2/pkg/render"
	"github.com/spf13/cobra"
)

var (
	cmdRoot = &cobra.Command{
		Use:   "goj2",
		Short: "CLI tool for Jinja2 template rendering",
		Long:  "CLI tool for Jinja2 template rendering",
		Run:   cmdRun,
	}
)

// Execute cobra
func Execute() {
	// Add flags to the root command
	cmdRoot.Flags().StringP("from", "f", "", "Input template file")
	if err := cmdRoot.MarkFlagRequired("from"); err != nil {
		log.Fatal(err)
	}
	cmdRoot.Flags().StringP("to", "t", "", "Output file")
	cmdRoot.Flags().StringSliceP("vars", "v", []string{}, "Variables file(s) in YAML format")
	if err := cmdRoot.MarkFlagRequired("vars"); err != nil {
		log.Fatal(err)
	}

	// Execute the root command
	if err := cmdRoot.Execute(); err != nil {
		log.Fatal(err)
	}
}

func cmdRun(cmd *cobra.Command, _ []string) {
	from, _ := cmd.Flags().GetString("from")
	to, _ := cmd.Flags().GetString("to")
	varsFiles, _ := cmd.Flags().GetStringSlice("vars")

	// Check for required parameters
	if from == "" {
		log.Fatal("Input template file is required!")
	}
	if len(varsFiles) == 0 {
		log.Fatal("Variables file is required!")
	}

	// Call function to render the template
	render.Template(from, to, varsFiles)
}
