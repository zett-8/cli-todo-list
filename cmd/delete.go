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
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete todo",
	Long: `You can delete your todo by this command.
	Put numbers you want to delete.
	e.g.  todo delete 2 12 4
	`,
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
			ErrorMsg(err)
			return
		}

		// delete by args
		if len(nums) > 0 {
			var filtered []Todo
			for i, todo := range todos {
				if !nums.contains(i + 1) {
					filtered = append(filtered, todo)
				}
			}
			todos = filtered

			err = WriteData(todos)
			if err != nil {
				ErrorMsg(err)
				return
			}

			ConsoleList(todos, "", false)

			return
		}

		// delete by flags
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
			ErrorMsg(err)
			return
		}

		ConsoleList(todos, "", false)

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
	deleteCmd.Flags().StringP("tag", "t", "", "Filter by tag")
	deleteCmd.Flags().BoolP("all", "a", false, "Delete all(or tag-filtered) todos")
	deleteCmd.Flags().BoolP("done", "d", false, "Delete all(or tag-filtered) finished todos")
}
