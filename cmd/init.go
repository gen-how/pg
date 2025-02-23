package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <project-name>",
	Short: "Initialize a new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		lang, _ := cmd.Flags().GetString("lang")
		buildSys, _ := cmd.Flags().GetString("build-sys")
		langStd, _ := cmd.Flags().GetString("std")

		fmt.Printf("Creating %s project: %s\n", lang, projectName)
		fmt.Printf("Using build system: %s\n", buildSys)
		if langStd != "" {
			fmt.Printf("Using language standard: %s\n", langStd)
		}
		// TODO: Implement project generation logic
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Init-specific flags
	initCmd.Flags().StringP("lang", "l", "c", "Set programming language (c/cpp/cu)")
	initCmd.Flags().StringP("build-sys", "b", "cmake", "Set build system (cmake/make)")
	initCmd.Flags().StringP("std", "s", "", "Set language standard (e.g., c11, c++17)")
	initCmd.Flags().Bool("tests", true, "Include test directory")
	initCmd.Flags().String("license", "MIT", "License type")
}
