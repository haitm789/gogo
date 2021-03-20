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

	"cli/adapter/cli"

	"github.com/spf13/cobra"
)

var mailCmd = &cobra.Command{
	Use:   "mail",
	Short: "collections of mail commands",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mail called")
	},
}

func loadMailCommands() []cli.Command {
	commands := cli.MailCommands()
	return commands
}

func init() {
	for _, c := range loadMailCommands() {
		mailCmd.AddCommand(newCmd(c))
	}

	rootCmd.AddCommand(mailCmd)
}
