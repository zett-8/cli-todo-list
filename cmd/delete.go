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
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var nums IntList
		tag, _ := cmd.Flags().GetString("tag")
		all, _ := cmd.Flags().GetBool("all")
		done, _ := cmd.Flags().GetBool("done")

		for _, n := range args {
			idx, _ := strconv.Atoi(n)
			nums = append(nums, idx)
		}

		todos, err := ReadData()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(args)

		// delete by flags
		if tag != "" || all || done {

			var deleteTodos []Todo
			if all {
				deleteTodos = todos
			}

			if done {
				var tmp []Todo
				for _, v := range todos {
					if v.done {
						tmp = append(tmp, v)
					}
				}
				deleteTodos = tmp
			}

			if tag != "" && len(deleteTodos) > 0 {
				var tmp []Todo
				for _, v := range deleteTodos {
					if v.containsTag(tag) {
						tmp = append(tmp, v)
					}
				}

				deleteTodos = tmp
			}

			var deleteIds IntList
			for _, v := range deleteTodos {
				deleteIds = append(deleteIds, v.id)
			}

			var filtered []Todo
			for _, v := range todos {
				if !deleteIds.contains(v.id) {
					filtered = append(filtered, v)
				}
			}

			err = WriteData(filtered)
			if err != nil {
				fmt.Println(err)
				return
			}

			return
		}

		// delete by args
		var filtered []Todo
		for i, todo := range todos {
			if !nums.contains(i + 1) {
				filtered = append(filtered, todo)
			}
		}
		todos = filtered

		err = WriteData(todos)
		if err != nil {
			fmt.Println(err)
			return
		}

		return
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	deleteCmd.Flags().StringP("tag", "t", "", "Help message for toggle")
	deleteCmd.Flags().BoolP("all", "a", false, "Help message for toggle")
	deleteCmd.Flags().BoolP("done", "d", false, "Help message for toggle")
}
