package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/brandon-gottshall/lfx/internal/config"
	"github.com/brandon-gottshall/lfx/internal/doctor"
	"github.com/brandon-gottshall/lfx/internal/paths"
	"github.com/brandon-gottshall/lfx/internal/registry"
	"github.com/brandon-gottshall/lfx/internal/ui"
)

func main() {
	if err := loadConfig(); err != nil {
		ui.PrintError("failed to load config", err)
		os.Exit(1)
	}

	args := os.Args[1:]
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		printHelp()
		return
	}

	switch {
	case len(args) >= 2 && args[0] == "theme" && args[1] == "list":
		listThemes()
	case len(args) >= 1 && args[0] == "doctor":
		runDoctor()
	default:
		ui.PrintError("unknown command", fmt.Errorf("%v", args))
		os.Exit(1)
	}
}

func loadConfig() error {
	cfgPath := config.DefaultPath(paths.LfxConfigDir())
	_, err := config.Load(cfgPath)
	return err
}

func listThemes() {
	repoRoot, err := os.Getwd()
	if err != nil {
		ui.PrintError("failed to determine working directory", err)
		os.Exit(1)
	}

	regRoot := filepath.Join(repoRoot, "registry")
	themes, err := registry.ListThemes(regRoot)
	if err != nil {
		ui.PrintError("failed to read themes registry", err)
		os.Exit(1)
	}

	ui.PrintTitle("Themes")
	if len(themes) == 0 {
		fmt.Println("(none)")
		return
	}
	for _, theme := range themes {
		fmt.Println("- " + theme)
	}
}

func runDoctor() {
	repoRoot, err := os.Getwd()
	if err != nil {
		ui.PrintError("failed to determine working directory", err)
		os.Exit(1)
	}

	regRoot := filepath.Join(repoRoot, "registry")
	result := doctor.Check(regRoot, paths.LfConfigDir())

	ui.PrintTitle("Doctor")
	if len(result.Issues) == 0 {
		fmt.Println("OK")
		return
	}

	for _, issue := range result.Issues {
		fmt.Println("- " + issue)
	}
	os.Exit(1)
}

func printHelp() {
	ui.PrintTitle("lfx")
	fmt.Println("Usage:")
	fmt.Println("  lfx theme list")
	fmt.Println("  lfx doctor")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  theme list   List available themes from registry")
	fmt.Println("  doctor       Check registry and lf config targets")
}
