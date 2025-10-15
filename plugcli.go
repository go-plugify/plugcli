package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func main() {
	Execute()
}

var rootCmd = &cobra.Command{
	Use:   "plugcli",
	Short: "PlugCLI - A plugin project scaffolding generator",
	Long:  `PlugCLI helps you quickly create a new Golang plugin project structure with predefined templates.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}

type PluginInfo struct {
	Name        string
	Description string
	Version     string
	NowTime     string
	Output      string
}

var createCmd = &cobra.Command{
	Use:   "create [plugin name]",
	Short: "Create a new plugin project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		info := PluginInfo{Name: name}

		qs := []*survey.Question{
			{
				Name:   "description",
				Prompt: &survey.Input{Message: "Plugin description:"},
			},
			{
				Name:   "version",
				Prompt: &survey.Input{Message: "Version:", Default: "0.0.1"},
			},
			{
				Name:   "output",
				Prompt: &survey.Input{Message: "Output directory:", Default: "./" + name},
			},
		}

		if err := survey.Ask(qs, &info); err != nil {
			return err
		}

		fmt.Println("✔ Creating plugin skeleton...")
		if err := createPluginSkeleton(info); err != nil {
			return err
		}
		fmt.Println("✔ Successfully created at:", info.Output)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
