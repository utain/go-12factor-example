package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"text/template"

	_ "embed"

	"github.com/utain/go/example/command"
	"github.com/utain/go/example/internal/adapters/viperadapter"
)

//go:embed banner.txt
var bannerText string

func main() {
	if err := viperadapter.Parse(); err != nil {
		log.Fatal("Failed to read configuration", err)
	}
	version := command.Versions{
		Runtime: fmt.Sprintf("%s %s %s", runtime.Version(), runtime.GOARCH, runtime.Compiler),
	}
	tmpl, err := template.New("banner").Parse(bannerText)
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(os.Stdout, version)
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
