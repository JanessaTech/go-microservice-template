package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "go-microservice-template.exe",
	Short: "A template for gin based micro-service project",
	Long: `
	This template will use the following tools:
	- gin : Http web framework
	- fx : A dependency injection system
	- gorm : Database access solution
	- koanf : a simple, extremely lightweight, extensible, configuration management library
	- cobra : A CLI application
	- zap :  A fast, structured, leveled logging
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start web server from rootCmd ...")
		fmt.Println("configFile =", configFile)
	},
}

var subCommand = &cobra.Command{
	Use:     "server",
	Short:   "start web server",
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(0), // only 0 parameter for command1
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start web server from subCommand ...")
		fmt.Println("configFile =", configFile)
	},
}

func init() {
	rootCmd.AddCommand(subCommand)
	rootCmd.PersistentFlags().StringVarP(&configFile, "conf", "c", "./config/properties.json", "config file path")
}

// .\go-microservice-template.exe server -c "./config/properties.json"
// .\go-microservice-template.exe -c "./config/properties.json"
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
