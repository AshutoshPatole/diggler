/*
Copyright © 2025 AshutoshPatole
*/
package cmd

import (
	"os"

	"github.com/AshutoshPatole/diggler/internal"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "diggler",
	Short: "A tool for gathering system information for forensic analysis",
	Long:  `Diggler is a tool for gathering system information for forensic analysis. This tool will gather various system information such as network interfaces, system logs, process information, and more.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		internal.GetHostInfo()
		internal.GetCPUInfo()
		internal.GetMemoryInfo()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.diggler.yaml)")
	rootCmd.Flags().BoolP("help", "h", false, "Help message for toggle")
}
