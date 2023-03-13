/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}

	Client = &http.Client{Transport: transCfg}
	NormalClient = Client
}

func isHttpMethod(m string) bool {
	m = strings.ToLower(m)
	if m == "get" || m == "post" || m == "put" || m == "delete" {
		return true
	}
	return false
}

func toHeaders(values []string) (map[string]string, error) {

	out := make(map[string]string)
	if values == nil || len(values) == 0 {
		return out, nil
	}

	for _, val := range values {
		part := strings.Split(val, "=")
		if len(part) != 2 {
			return nil, fmt.Errorf("error , header %s", val)
		}
		k := strings.TrimSpace(part[0])
		v := strings.TrimSpace(part[1])
		out[k] = v
	}

	return out, nil
}

// urlCmd represents the url command
var urlCmd = &cobra.Command{
	Use:   "url [url]",
	Short: "connect url ",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			return
		}
		url := args[0]

		if !isHttpMethod(method) {
			cmd.Help()
			return
		}
		headers, err := toHeaders(header)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		RestDo(nil, method, url, headers, nil)
	},
}

var (
	debug  = false
	method = "GET"
	header = make([]string, 0)
)

func init() {
	rootCmd.AddCommand(urlCmd)
	urlCmd.Flags().BoolVarP(&debug, "debug", "d", false, " -d ")
	urlCmd.Flags().StringVarP(&method, "method", "m", "get", "http method, get | post | put | delete")
	urlCmd.Flags().StringArrayVarP(&header, "header", "H", nil, "http header , e.g. \"hedername=headervalue\" ")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// urlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// urlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var (
	// default http client
	Client       *http.Client
	NormalClient *http.Client
)

const ContentType = "Content-Type"

func RestDo(client *http.Client, method string, url string, headers map[string]string,
	data io.Reader) {

	if client == nil {
		client = Client
	}

	method = strings.ToUpper(method)

	lCase := strings.ToLower(url)
	if strings.HasPrefix(lCase, "https://") || strings.HasPrefix(lCase, "http://") {

	} else {
		url = "http://" + url
	}

	req, err := http.NewRequest(method, url, data)
	if err != nil {
		fmt.Printf("url create req error %s %s %v\n", method, url, err)
		return
	}

	if headers != nil && len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	if _, ok := req.Header[ContentType]; !ok {
		req.Header.Add("Content-Type", "application/json")
	}

	if debug {
		fmt.Printf("headers :\n")
		for k, v := range req.Header {
			fmt.Printf("%s : %s\n", k, v)
		}
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("client do error, url %s %s %v\n", method, url, err)
		return
	}
	// check err before body close
	if resp.Body != nil {
		defer resp.Body.Close()

		//
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			fmt.Printf("read all error, url %s %s %v\n", method, url, err1)
			return
		}
		fmt.Printf("%s status %s\n", method, resp.Status)
		fmt.Printf(string(body))
		return
	} else {
		fmt.Printf("%s status %s\n", method, resp.Status)

	}

}
