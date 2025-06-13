package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "This is a sample command",
	Long: `This is a longer description for the sample command. Example usage:
			"go run . --sample [name to output]"`,

	Run: sampleCmdFunc,
}

func init() {
	rootCmd.AddCommand(sampleCmd)

	sampleCmd.Flags().BoolP("help", "h", false, "Display command help")
	sampleCmd.Flags().BoolP("version", "v", false, "Print sample version")
	sampleCmd.Flags().StringP("name", "n", "gabriel", "Input name to be displayed")

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
