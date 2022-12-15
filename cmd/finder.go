package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var findBigFileCmd = &cobra.Command{
	Use:   "find_big_file",
	Short: "$> find / -type f -size +1G  -exec du -h {} \\;",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("#############################################")
		fmt.Println("# To find big file over 1G under the root path :")
		fmt.Println("# sudo find / -type f -size +1G")
		fmt.Println("# or ")
		fmt.Println("# sudo find / -type f -size +1G -exec du -h {} \\;")
		fmt.Println("#############################################")
		execCmd := exec.Command("find", "/", "-type", "f", "-size", "+1G")

		stdout, err := execCmd.StdoutPipe()
		if err != nil {
			log.Fatalln(err)
		}
		execCmd.Stderr = execCmd.Stdout
		if err = execCmd.Start(); err != nil {
			log.Fatalln(err)
		}
		for {
			tmp := make([]byte, 1024)
			_, err := stdout.Read(tmp)
			if err != nil {
				break
			}
			fmt.Print(string(tmp))
		}
	},
}

func init() {
	RootCmd.AddCommand(findBigFileCmd)
}
