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
	ID          string
	Name        string
	Description string
	Version     string
	ServerAddr  string
	Author      string
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

		clientType := ""
		prompt := &survey.Select{
			Message: "Choose client script language:",
			Options: []string{"yaegi", "native_go_plugin"},
			Default: "yaegi",
		}
		err := survey.AskOne(prompt, &clientType)
		if err != nil {
			return err
		}

		qs := []*survey.Question{
			{
				Name:     "id",
				Prompt:   &survey.Input{Message: "Plugin unique ID:"},
				Validate: survey.Required,
			},
			{
				Name:   "description",
				Prompt: &survey.Input{Message: "Plugin description:"},
			},
			{
				Name:     "version",
				Prompt:   &survey.Input{Message: "Version:", Default: "0.0.1"},
				Validate: survey.Required,
			},
			{
				Name:   "author",
				Prompt: &survey.Input{Message: "Author:", Default: "Your Name"},
			},
			{
				Name:     "serverAddr",
				Prompt:   &survey.Input{Message: "Server API address:", Default: "http://localhost:8080/api/v1"},
				Validate: survey.Required,
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
		switch clientType {
		case "native_go_plugin":
			if err := createPluginSkeletonOfNativePlugin(info); err != nil {
				return err
			}
		case "yaegi":
			if err := createPluginSkeletonOfYaegi(info); err != nil {
				return err
			}
		}
		fmt.Println("✔ Successfully created at:", info.Output)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
