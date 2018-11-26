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
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func IsPipe(file *os.File) bool {
	fi, err := file.Stat()
	if err != nil {
		return false
	}

	return (fi.Mode() & os.ModeNamedPipe) != 0
}

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find the JSON path to bits of the JSON blob",
	Long:  "Find the JSON path to bits of the JSON blob",
	Run: func(cmd *cobra.Command, args []string) {
		findCmdRun()
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			if IsPipe(os.Stdin) {
				options.file = os.Stdin
			}
			return errors.New("requires at least 1 file as an argument")
		} else {
			fi, err := os.Stat(args[0])
			if err != nil {
				return err
			}

			if fi.Mode().IsRegular() == false {
				return errors.Errorf("%q is not a file", args[0])
			}

			options.filename = args[0]
		}

		return nil
	},
}

type findCmdOptions struct {
	key   string
	value string
	index bool

	file     *os.File
	filename string
}

var options findCmdOptions

func init() {
	rootCmd.AddCommand(findCmd)
	findCmd.Flags().StringVarP(&options.key, "key", "k", "", "The prefix of keys you want to find")
	findCmd.Flags().StringVarP(&options.value, "value", "v", "", "The prefix of values you want to find")
	findCmd.Flags().BoolVarP(&options.index, "index", "i", false, "Provide the actual index in the JSON path")
}

func findCmdRun() {
	fmt.Println(options.key)
	fmt.Println(options.value)
	fmt.Println(options.index)

	fmt.Println(options.file == nil)
	fmt.Println(options.filename)
}
