/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

// addrCmd represents the addr command
var addrCmd = &cobra.Command{
	Use:   "addr",
	Short: "get local interface & addr",

	Run: func(cmd *cobra.Command, args []string) {
		listAddrs()
	},
}

func init() {
	rootCmd.AddCommand(addrCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addrCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Gethardw(inf net.Interface) string {
	val := inf.HardwareAddr.String()
	if val != "" {
		return val
	}
	return "N"
}
func listAddrs() {
	infs, err := net.Interfaces()
	if err != nil {
		fmt.Printf("interface error %v\n", err)
		return
	}

	for _, inf := range infs {
		fmt.Printf("%-20s %-30s %-30s \n", Gethardw(inf), inf.Flags.String(), inf.Name)
		addrs, e := inf.Addrs()
		if e != nil {
			fmt.Printf(" get addr error %v\n", err)
			continue
		}
		for _, addr := range addrs {
			fmt.Printf("    %-20s %-20s\n", addr.String(), addr.Network())
		}

		addrs, e = inf.MulticastAddrs()
		if e != nil {
			fmt.Printf(" get multicase addr error %v\n", err)
			continue
		}
		for _, addr := range addrs {
			fmt.Printf("    %-20s %-20s\n", addr.String(), addr.Network())
		}

		fmt.Printf("\n")
	}

}
