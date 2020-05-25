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

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create", "make", "new"},
	Short:   "Add new todo. use '@' to add with tags",
	Long: `You can add new todo by this command. Use quotes to wrap your todo.
To add todo with tags put '@' before tag name.
e.g.  todo add 'call John til 4pm' @work`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("you need to type what to do!")
			return
		}

		todos, err := ReadData()
		if err != nil {
			ErrorMsg(err)
			return
		}

		var todo string
		var tags []string

		for _, v := range args {
			if todo == "" && string(v[0]) != "@" {
				todo = v
			} else if string(v[0]) == "@" {
				tags = append(tags, v)
			}
		}

		todos = append(todos, Todo{0, false, todo, tags})

		err = WriteData(todos)
		if err != nil {
			ErrorMsg(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
