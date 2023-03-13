/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

// nsCmd represents the ns command
var nsCmd = &cobra.Command{
	Use:   "ns [hostname]",
	Short: "lookup ns name",
	Long:  `lookup hostname ip address . For example: www.siemens.com `,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {

			displayHostName(args[0])
		} else {
			cmd.Usage()
		}
	},
}

func init() {
	rootCmd.AddCommand(nsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func displayHostName(val string) {
	names, err := net.LookupHost(val)
	if err != nil {
		fmt.Printf("Err %v\n", err)
		return
	}

	for i, n := range names {
		fmt.Printf("%d - %v\n", i, n)
	}
}
