package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// Language defines the supported languages
type Language string

const (
    English  Language = "en"
    Chinese  Language = "zh"
    Japanese Language = "ja"
)

// Messages holds all the text messages for different languages
type Messages struct {
    RootUse              string
    RootShort            string
    RootLong             string
    CreateUse            string
    CreateShort          string
    ClientLanguagePrompt string
    PluginIDPrompt       string
    DescriptionPrompt    string
    VersionPrompt        string
    AuthorPrompt         string
    ServerAddrPrompt     string
    OutputDirPrompt      string
    CreatingMessage      string
    SuccessMessage       string
    ErrorMessage         string
    InvalidLangMessage   string
}

var messages = map[Language]Messages{
    English: {
        RootUse:              "plugcli",
        RootShort:            "PlugCLI - A plugin project scaffolding generator",
        RootLong:             `PlugCLI helps you quickly create a new Golang plugin project structure with predefined templates.`,
        CreateUse:            "create [plugin name]",
        CreateShort:          "Create a new plugin project",
        ClientLanguagePrompt: "Choose client script language:",
        PluginIDPrompt:       "Plugin unique ID:",
        DescriptionPrompt:    "Plugin description:",
        VersionPrompt:        "Version:",
        AuthorPrompt:         "Author:",
        ServerAddrPrompt:     "Server API address:",
        OutputDirPrompt:      "Output directory:",
        CreatingMessage:      "✔ Creating plugin skeleton...",
        SuccessMessage:       "✔ Successfully created at:",
        ErrorMessage:         "Error:",
        InvalidLangMessage:   "Invalid language. Supported: en, zh, ja",
    },
    Chinese: {
        RootUse:              "plugcli",
        RootShort:            "PlugCLI - 插件项目脚手架生成器",
        RootLong:             `PlugCLI 帮助您快速创建新的 Golang 插件项目结构和预定义模板。`,
        CreateUse:            "create [插件名称]",
        CreateShort:          "创建新的插件项目",
        ClientLanguagePrompt: "选择客户端脚本语言:",
        PluginIDPrompt:       "插件唯一ID:",
        DescriptionPrompt:    "插件描述:",
        VersionPrompt:        "版本号:",
        AuthorPrompt:         "作者:",
        ServerAddrPrompt:     "服务器API地址:",
        OutputDirPrompt:      "输出目录:",
        CreatingMessage:      "✔ 正在创建插件骨架...",
        SuccessMessage:       "✔ 成功创建于:",
        ErrorMessage:         "错误:",
        InvalidLangMessage:   "无效的语言。支持的语言: en, zh, ja",
    },
    Japanese: {
        RootUse:              "plugcli",
        RootShort:            "PlugCLI - プラグインプロジェクト足場ジェネレーター",
        RootLong:             `PlugCLI は事前定義されたテンプレートで新しい Golang プラグインプロジェクト構造を迅速に作成するのに役立ちます。`,
        CreateUse:            "create [プラグイン名]",
        CreateShort:          "新しいプラグインプロジェクトを作成",
        ClientLanguagePrompt: "クライアントスクリプト言語を選択:",
        PluginIDPrompt:       "プラグインユニークID:",
        DescriptionPrompt:    "プラグインの説明:",
        VersionPrompt:        "バージョン:",
        AuthorPrompt:         "作者:",
        ServerAddrPrompt:     "サーバーAPIアドレス:",
        OutputDirPrompt:      "出力ディレクトリ:",
        CreatingMessage:      "✔ プラグインスケルトンを作成中...",
        SuccessMessage:       "✔ 正常に作成されました:",
        ErrorMessage:         "エラー:",
        InvalidLangMessage:   "無効な言語。サポートされている言語: en, zh, ja",
    },
}

var currentLang Language = English
var currentMessages Messages = messages[English]

func main() {
    Execute()
}

func setLanguage(lang Language) error {
    if msg, exists := messages[lang]; exists {
        currentLang = lang
        currentMessages = msg
        return nil
    }
    fmt.Printf("%s %s\n", messages[English].InvalidLangMessage, lang)
    os.Exit(1)
    return nil
}

func Execute() {
    // Initialize commands with proper language support
    rootCmd := &cobra.Command{
        Use:   currentMessages.RootUse,
        Short: currentMessages.RootShort,
        Long:  currentMessages.RootLong,
        PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
            lang, _ := cmd.Flags().GetString("lang")
            if err := setLanguage(Language(lang)); err != nil {
                return err
            }
            // Update command descriptions after language is set
            cmd.Short = currentMessages.RootShort
            cmd.Long = currentMessages.RootLong
            return nil
        },
    }

    createCmd := &cobra.Command{
        Use:   currentMessages.CreateUse,
        Short: currentMessages.CreateShort,
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            // Update command description based on current language
            cmd.Use = currentMessages.CreateUse
            cmd.Short = currentMessages.CreateShort
            
            name := args[0]
            info := PluginInfo{Name: name}

            clientType := ""
            prompt := &survey.Select{
                Message: currentMessages.ClientLanguagePrompt,
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
                    Prompt:   &survey.Input{Message: currentMessages.PluginIDPrompt},
                    Validate: survey.Required,
                },
                {
                    Name:   "description",
                    Prompt: &survey.Input{Message: currentMessages.DescriptionPrompt},
                },
                {
                    Name:     "version",
                    Prompt:   &survey.Input{Message: currentMessages.VersionPrompt, Default: "0.0.1"},
                    Validate: survey.Required,
                },
                {
                    Name:   "author",
                    Prompt: &survey.Input{Message: currentMessages.AuthorPrompt, Default: "Your Name"},
                },
                {
                    Name:     "serverAddr",
                    Prompt:   &survey.Input{Message: currentMessages.ServerAddrPrompt, Default: "http://localhost:8080/api/v1"},
                    Validate: survey.Required,
                },
                {
                    Name:   "output",
                    Prompt: &survey.Input{Message: currentMessages.OutputDirPrompt, Default: "./" + name},
                },
            }

            if err := survey.Ask(qs, &info); err != nil {
                return err
            }

            fmt.Println(currentMessages.CreatingMessage)
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
            fmt.Println(currentMessages.SuccessMessage, info.Output)
            return nil
        },
    }

    rootCmd.AddCommand(createCmd)
    rootCmd.PersistentFlags().StringP("lang", "l", "en", "Language (en/zh/ja)")

    if err := rootCmd.Execute(); err != nil {
        fmt.Printf("%s %v\n", currentMessages.ErrorMessage, err)
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