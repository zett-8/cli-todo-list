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

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"dn", "fin", "finished", "did"},
	Short:   "Set todo as done",
	Long: `You can set your todo done by this command.
Put numbers you want to set it done.
e.g.  todo done 2 12 4 `,
	Run: func(cmd *cobra.Command, args []string) {
		var nums IntList

		for _, n := range args {
			idx, _ := strconv.Atoi(n)
			nums = append(nums, idx)
		}

		if len(nums) == 0 {
			return
		}

		todos, err := ReadData()
		if err != nil {
			ErrorMsg(err)
			return
		}

		for i, v := range todos {
			if nums.contains(v.id) {
				todos[i].done = true
			}
		}

		err = WriteData(todos)
		if err != nil {
			ErrorMsg(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
