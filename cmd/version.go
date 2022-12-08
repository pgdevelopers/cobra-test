package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(NewABCCMD())
	rootCmd.AddCommand(NewVersionCMD("1.0"))
}

func NewABCCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "abc",
		Short: "abc123",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
			fmt.Fprintf(cmd.OutOrStdout(), "abc test")
		},
	}
}

func NewVersionCMD(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Hugo",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
			fmt.Fprintf(cmd.OutOrStdout(), version)
		},
	}
}
