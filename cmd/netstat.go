/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/drael/GOnetstat"
	"github.com/spf13/cobra"
)

const (
	LineFormat  = "%-24s %-24s %-16s %-40s"
	LineTitle   = "Proto " + LineFormat + "\n"
	LineContent = "%5s " + LineFormat + "\n"
)

var types = []string{"tcp", "tcp6", "udp", "udp6"}

// netstatCmd represents the netstat command
var netstatCmd = &cobra.Command{
	Use:   "netstat",
	Short: "list network stats, like netstat command",

	Run: func(cmd *cobra.Command, args []string) {

		// format header
		fmt.Printf(LineTitle, "Local Adress", "Foregin Adress",
			"State", "Pid/Program")
		for _, t := range types {
			showNetStat(t)
			fmt.Println("")
		}

	},
}

func init() {
	rootCmd.AddCommand(netstatCmd)

}

func showNetStat(t string) {

	var d []GOnetstat.Process
	if t == "tcp" {
		d = GOnetstat.Tcp()
	} else if t == "tcp" {
		d = GOnetstat.Tcp6()
	} else if t == "udp" {
		d = GOnetstat.Udp()
	} else if t == "udp6" {
		d = GOnetstat.Udp6()
	} else {
		d = GOnetstat.Tcp()
	}

	displayNetStat(t, d, "LISTEN", true)
	displayNetStat(t, d, "LISTEN", false)
}

func displayNetStat(name string, d []GOnetstat.Process, wanted string, w bool) {
	for _, p := range d {

		if w {
			if p.State != wanted {
				continue
			}
		} else {
			if p.State == wanted {
				continue
			}
		}

		// format data like netstat output
		ip_port := fmt.Sprintf("%v:%v", p.Ip, p.Port)
		fip_port := fmt.Sprintf("%v:%v", p.ForeignIp, p.ForeignPort)
		pid_program := fmt.Sprintf("%v/%v", p.Pid, p.Name)

		fmt.Printf(LineContent, name, ip_port, fip_port,
			p.State, pid_program)

	}
}
