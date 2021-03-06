// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/bigflood/gostudy/todo/store"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use: "list",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := OpenStore(cmd)
		if err != nil {
			return err
		}

		filter := store.Filter{}

		switch {
		case flagDone:
			v := flagDone
			filter = store.Filter{Done: &v}
		case flagNotDone:
			v := !flagNotDone
			filter = store.Filter{Done: &v}
		}

		tasks, err := s.List(filter)
		if err != nil {
			return err
		}

		for i, task := range tasks {
			fmt.Printf("%d: ", i)
			if task.Done {
				fmt.Print("[O] ")
			} else {
				fmt.Print("[X] ")
			}
			fmt.Println(task.Desc)
		}

		return nil
	},
}

var (
	flagDone    bool
	flagNotDone bool
)

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().BoolVarP(&flagDone,
		"done", "d", false, "완료된 항목들만 리스트함")
	listCmd.Flags().BoolVarP(&flagNotDone,
		"not-done", "n", false, "완료 안된 항목들만 리스트함")
}
