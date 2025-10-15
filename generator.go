package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"
)

//go:embed tmpl
var templatesFS embed.FS

func createPluginSkeleton(info PluginInfo) error {
	// check if output directory exists, if not create it, if exists, return error
	stat, err := os.Stat(info.Output)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(info.Output, 0755); err != nil {
			return err
		}
		if err := os.MkdirAll(info.Output + "/src", 0755); err != nil {
			return err
		}
	} else {
		if err != nil {
			return err
		}
		if !stat.IsDir() {
			return err
		}
		if stat.Size() != 0 {
			return fmt.Errorf("output directory %s is not empty", info.Output)
		}
	}
	info.NowTime = time.Now().Format("20060102150405")
	tmpls := template.Must(template.ParseFS(templatesFS, "tmpl/**/*.tmpl", "tmpl/*.tmpl"))
	for _, fileName := range fileNames {
		var buf bytes.Buffer
		var baseFileName = strings.ReplaceAll(fileName, "src/", "")
		err = tmpls.ExecuteTemplate(&buf, baseFileName+".tmpl", info)
		if err != nil {
			return err
		}
		err = os.WriteFile(info.Output+"/"+fileName, buf.Bytes(), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

var fileNames = []string{
	"src/plugin.go",
	"src/define.go",
	"Makefile",
	"src/go.mod",
}
