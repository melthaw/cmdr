package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var findCmd = &cobra.Command{
	Use:   "find_big_file",
	Short: "$> find / -type f -size +1G  -exec du -h {} \\;",
	Run: func(cmd *cobra.Command, args []string) {
		execCmd := exec.Command("find", " / -type f -size +1G  -exec du -h {} \\;")
		execCmd.Stdout = os.Stdout
		_ = execCmd.Run()
	},
}

func init() {
	RootCmd.AddCommand(findCmd)
}
