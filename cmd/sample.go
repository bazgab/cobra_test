package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "This is a sample command",
	Long: `This is a longer description for the sample command. Example usage:
			"go run . --sample [name to output]"`,

	Run: sampleCmdFunc,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print sample version",
	Long:  `Print sample version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1.0")
	},
}

func init() {
	rootCmd.AddCommand(sampleCmd)
	rootCmd.AddCommand(versionCmd)
	i := os.Getenv("GOPATH")
	sampleCmd.Flags().StringP("name", "N", i, "Input name to be displayed")

}

func sampleCmdFunc(cmd *cobra.Command, args []string) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		c := color.New(color.FgRed, color.Bold).SprintFunc()
		p := c("[ERROR] ")
		fmt.Println(p, "Command failed to execute")
	}

	c := color.New(color.FgBlue, color.Bold).SprintFunc()
	p := c("[INFO] ")
	fmt.Println(p, "Command successfully executed.")
	fmt.Printf("%s Name that was entered: %s\n", p, name)

}
