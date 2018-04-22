package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Verbose controls the logging level, when enabled will set level to debug
var Verbose bool

var rootCmd = &cobra.Command{
	Use:   "svcbndl",
	Short: "svcbndl is a tool to manage ServiceBundle images",
	Long: `ServiceBundles are images that represent lifecycle components
in that they contain all of the orchestration logic to manage
an application through out it's lifecycle, i.e. install, uninstall,
bind, unbind, etc.  ServiceBundles are intended to be invoked and run
as a short job to execute the intended work, example I want to deploy a
postgres database to my kubernetes cluster.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if Verbose {
			log.SetLevel(log.DebugLevel)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run has been invoked.")
	},
}

func init() {
	log.SetLevel(log.WarnLevel)
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

// Execute invokes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
