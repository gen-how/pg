package cmd

import (
	"fmt"
	"os"

	"github.com/gen-how/pg/internal/generator"
	"github.com/spf13/cobra"
)

var (
	language string
	isApp    bool
	isLib    bool
	useMake  bool
	useCMake bool
	license  string
)

var rootCmd = &cobra.Command{
	Use:   "pg <project-name>",
	Short: "Project generator for C and C++ projects",
	Long:  "A CLI tool to generate initial structure for C and C++ projects",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		// Validate language
		if language != "c" && language != "cpp" {
			fmt.Println("Error: Language must be 'c' or 'cpp'")
			os.Exit(1)
		}

		// Default to app if neither app nor lib is specified
		if !isApp && !isLib {
			isApp = true
		}

		// Default to make if no build system is specified
		if !useMake && !useCMake {
			useMake = true
		}

		config := generator.ProjectConfig{
			Name:     projectName,
			Language: language,
			IsApp:    isApp,
			IsLib:    isLib,
			UseMake:  useMake,
			UseCMake: useCMake,
			License:  license,
		}

		err := generator.Generate(config)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("info: successfully created project: '%s'\n", projectName)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&language, "lang", "l", "", "Specify language (required) [options: c, cpp]")
	rootCmd.Flags().BoolVar(&isApp, "app", false, "Set to create application source code")
	rootCmd.Flags().BoolVar(&isLib, "lib", false, "Set to create library sourece code")
	rootCmd.Flags().BoolVar(&useMake, "make", false, "Set to create Makefile")
	rootCmd.Flags().BoolVar(&useCMake, "cmake", false, "Set to create CMake configuration files")
	rootCmd.Flags().StringVar(&license, "license", "", "Specify license type [options: mit, apache2]")

	rootCmd.MarkFlagRequired("lang")
}
