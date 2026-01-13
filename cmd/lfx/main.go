package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/brandon-gottshall/lfx/internal/config"
	"github.com/brandon-gottshall/lfx/internal/doctor"
	"github.com/brandon-gottshall/lfx/internal/install"
	"github.com/brandon-gottshall/lfx/internal/paths"
	"github.com/brandon-gottshall/lfx/internal/registry"
	"github.com/brandon-gottshall/lfx/internal/ui"
)

func main() {
	cfgPath := config.DefaultPath(paths.LfxConfigDir())
	if err := loadConfig(cfgPath); err != nil {
		ui.PrintError("failed to load config", err)
		os.Exit(1)
	}

	args := os.Args[1:]
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		printHelp()
		return
	}

	switch {
	case len(args) >= 3 && args[0] == "theme" && args[1] == "set":
		if len(args) != 3 {
			ui.PrintError("invalid arguments", fmt.Errorf("expected: lfx theme set <name>"))
			os.Exit(1)
		}
		setTheme(args[2])
	case len(args) >= 2 && args[0] == "theme" && args[1] == "list":
		listThemes()
	case len(args) >= 3 && args[0] == "plugin" && args[1] == "install":
		if len(args) != 3 {
			ui.PrintError("invalid arguments", fmt.Errorf("expected: lfx plugin install <name>"))
			os.Exit(1)
		}
		installPlugin(args[2])
	case len(args) >= 3 && args[0] == "plugin" && args[1] == "uninstall":
		if len(args) != 3 {
			ui.PrintError("invalid arguments", fmt.Errorf("expected: lfx plugin uninstall <name>"))
			os.Exit(1)
		}
		uninstallPlugin(args[2])
	case len(args) >= 2 && args[0] == "plugin" && args[1] == "list":
		listPlugins()
	case len(args) >= 1 && args[0] == "doctor":
		runDoctor()
	default:
		ui.PrintError("unknown command", fmt.Errorf("%v", args))
		os.Exit(1)
	}
}

func loadConfig(configPath string) error {
	_, err := config.Load(configPath)
	if err == nil {
		return nil
	}
	switch {
	case errors.Is(err, config.ErrConfigMissing):
		ui.PrintWarning(fmt.Sprintf("missing config at %s; using defaults", configPath), err)
		return nil
	case errors.Is(err, config.ErrConfigInvalid):
		ui.PrintWarning(fmt.Sprintf("invalid config at %s; using defaults", configPath), err)
		return nil
	default:
		return err
	}
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

func listPlugins() {
	repoRoot, err := os.Getwd()
	if err != nil {
		ui.PrintError("failed to determine working directory", err)
		os.Exit(1)
	}

	regRoot := filepath.Join(repoRoot, "registry")
	plugins, err := registry.ListPlugins(regRoot)
	if err != nil {
		ui.PrintError("failed to read plugins registry", err)
		os.Exit(1)
	}

	ui.PrintTitle("Plugins")
	if len(plugins) == 0 {
		fmt.Println("(none)")
		return
	}
	for _, plugin := range plugins {
		fmt.Println("- " + plugin)
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
	fmt.Println("  lfx theme set <name>")
	fmt.Println("  lfx plugin list")
	fmt.Println("  lfx plugin install <name>")
	fmt.Println("  lfx plugin uninstall <name>")
	fmt.Println("  lfx doctor")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  theme list   List available themes from registry")
	fmt.Println("  theme set    Apply a vendored theme")
	fmt.Println("  plugin list  List available plugins from registry")
	fmt.Println("  plugin install  Install a vendored plugin snippet")
	fmt.Println("  plugin uninstall  Remove a previously installed plugin snippet")
	fmt.Println("  doctor       Check registry and lf config targets")
}

func setTheme(themeName string) {
	repoRoot, err := os.Getwd()
	if err != nil {
		ui.PrintError("failed to determine working directory", err)
		os.Exit(1)
	}

	regRoot := filepath.Join(repoRoot, "registry")
	if err := install.ApplyTheme(regRoot, themeName, paths.LfConfigDir()); err != nil {
		ui.PrintError("failed to apply theme", err)
		os.Exit(1)
	}

	ui.PrintTitle("Theme")
	fmt.Printf("applied %s\n", themeName)
}

func installPlugin(pluginName string) {
	repoRoot, err := os.Getwd()
	if err != nil {
		ui.PrintError("failed to determine working directory", err)
		os.Exit(1)
	}

	regRoot := filepath.Join(repoRoot, "registry")
	if err := install.ApplyPlugin(regRoot, pluginName, paths.LfConfigDir()); err != nil {
		ui.PrintError("failed to install plugin", err)
		os.Exit(1)
	}

	ui.PrintTitle("Plugin")
	fmt.Printf("installed %s\n", pluginName)
}

func uninstallPlugin(pluginName string) {
	if pluginName == "" {
		ui.PrintError("invalid arguments", fmt.Errorf("expected: lfx plugin uninstall <name>"))
		os.Exit(1)
	}
	if err := install.RemovePlugin(paths.LfConfigDir(), pluginName); err != nil {
		ui.PrintError("failed to uninstall plugin", err)
		os.Exit(1)
	}

	ui.PrintTitle("Plugin")
	fmt.Printf("uninstalled %s\n", pluginName)
}
