package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/scottgreenup/json-util/pkg/json"
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
			} else {
				return errors.New("requires at least 1 file as an argument")
			}
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
	findCmd.Flags().BoolVarP(&options.index, "index", "i", false, "Provide the actual index in the JSON path")
}

func findCmdRun() error {
	if options.key == "" {
		return errors.New("No 'key' to search for was provided")
	}

	var file *os.File

	if options.file != nil {
		file = options.file
	} else if len(options.filename) > 0 {
		var err error
		file, err = os.Open(options.filename)
		if err != nil {
			return err
		}
	} else {
		return errors.New("no file provided")
	}

	input, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	obj, err := json.NewObject(input)
	if err != nil {
		return err
	}

	getSuffix := func(x string) string {
		splat := strings.Split(x, ".")
		return splat[len(splat)-1]
	}

	removeIndex := func(x string) string {
		m := make([]rune, 0)
		inside := false
		for _, c := range x {
			if c == '[' {
				m = append(m, c)
				inside = true
			} else if c == ']' {
				m = append(m, c)
				inside = false
			} else if !inside {
				m = append(m, c)
			}
		}
		return string(m)
	}

	withIndex := func(path string, value *json.Value) {
		key := getSuffix(path)
		if key == options.key {
			fmt.Println(path)
		}
	}

	withoutIndexCache := make([]string, 0)
	withoutIndex := func(path string, value *json.Value) {
		key := getSuffix(path)
		if key == options.key {
			indexless := removeIndex(path)
			for _, cacheValue := range withoutIndexCache {
				if cacheValue == indexless {
					return
				}
			}
			withoutIndexCache = append(withoutIndexCache, indexless)
		}
	}

	if !options.index {
		obj.Walk(withoutIndex)
		for _, cacheValue := range withoutIndexCache {
			fmt.Println(cacheValue)
		}
	} else {
		obj.Walk(withIndex)
	}

	return nil
}
