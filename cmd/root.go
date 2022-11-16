package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var RootCmd = &cobra.Command{
	Use:   "cmdr",
	Short: "The linux shell command instruction(template) to get task done quickly.",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
