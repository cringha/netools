/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"time"
)

// connCmd represents the conn command
var connCmd = &cobra.Command{
	Use: "conn host port",

	Short: "connect to remote server host port ",
	Long:  `Connect to remote server . For example: 	192.168.72.140 8080 `,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			cmd.Help()
			return
		}
		tcpConnect(args[0], args[1])
	},
}

var (
	connHost = ""
	connPort = 0
)

func init() {
	rootCmd.AddCommand(connCmd)

}

func tcpConnect(host, port string) {
	host1 := fmt.Sprintf("%s:%s", host, port)

	//tcpAddr, err1 := net.ResolveTCPAddr("tcp4", host1)
	//if err1 != nil {
	//	fmt.Printf("Err ResolveTCPAddr %v\n", err1)
	//	return
	//}

	conn, err := net.DialTimeout("tcp", host1, time.Second*2)
	if err != nil {
		fmt.Printf("Err %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Printf("%s connected\n", host)

	conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	var response = make([]byte, 255)

	n, err := conn.Read(response[0:])
	if err != nil {

		fmt.Printf("Err read %v\n", err)
		return
	}

	if n > 0 {
		fmt.Printf("%v\n", string(response[:n]))
	}

}
