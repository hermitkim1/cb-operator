/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/cloud-barista/cb-operator/src/common"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get information of Cloud-Barista System",
	Long:  `Get information of Cloud-Barista System. Information about containers and container images`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n[Get info for Cloud-Barista runtimes]\n")

		if common.FileStr == "" {
			fmt.Println("file is required")
		} else {
			fmt.Println("[v]Status of Cloud-Barista runtimes")
			cmdStr := "sudo docker-compose -f " + common.FileStr + " ps"
			//fmt.Println(cmdStr)
			common.SysCall(cmdStr)

			fmt.Println("")
			fmt.Println("[v]Status of Cloud-Barista runtime images")
			cmdStr = "sudo docker-compose -f " + common.FileStr + " images"
			//fmt.Println(cmdStr)
			common.SysCall(cmdStr)
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	pf := infoCmd.PersistentFlags()
	pf.StringVarP(&common.FileStr, "file", "f", "../docker-compose-mode-files/docker-compose.yaml", "Path to Cloud-Barista Docker-compose file")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
