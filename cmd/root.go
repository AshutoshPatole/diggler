/*
Copyright © 2025 AshutoshPatole
*/
package cmd

import (
	"os"

	"github.com/AshutoshPatole/diggler/internal"
	"github.com/spf13/cobra"
)

var save bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "diggler",
	Short: "A tool for gathering system information for forensic analysis",
	Long:  `Diggler is a tool for gathering system information for forensic analysis. This tool will gather various system information such as network interfaces, system logs, process information, and more.`,

	Run: func(cmd *cobra.Command, args []string) {
		if save {
			file, err := os.Create("system_info.txt")
			if err != nil {
				panic(err)
			}
			defer file.Close()
			os.Stdout = file
		}
		internal.GetHostInfo()
		internal.GetCPUInfo()
		internal.GetMemoryInfo()
		internal.GetNTPInfo()
		// internal.GetOpenFiles()
		// internal.GetConnections()
		// internal.GetSecurityInfo()
		// internal.FirewallStat()
		// internal.GetDNSInfo()
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
	rootCmd.Flags().BoolVarP(&save, "save", "s", false, "Save the output to a file")
}
