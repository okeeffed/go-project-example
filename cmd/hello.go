/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/okeeffed/go-project-example/pkg/greet"
	"github.com/spf13/cobra"
)

type Json struct {
	Title string `json:"title"`
	Value int32  `json:"value"`
}

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(greet.Morning)
		data, err := os.ReadFile(source)

		if err != nil {
			panic(err)
		}

		// jsonParsed, err := gabs.ParseJSON(data)
		var jsonData Json
		json.Unmarshal(data, &jsonData)

		if err != nil {
			panic(err)
		}

		// Parse the JSON string value
		// value, ok := jsonParsed.Path("title").Data().(string)

		// if !ok {
		// 	panic("Not found")
		// }

		// fmt.Println(value)

		// numberValue, numOk := jsonParsed.Path("value").Data().(float64)

		// if !numOk {
		// 	panic("Number not found")
		// }

		// fmt.Println(numberValue)

		// // Setting a new value
		// jsonParsed.Set(numberValue+2, "value")
		// jsonParsed.Set(value+" hello", "title")

		jsonData.Title = "New title"

		finalData, _ := json.MarshalIndent(jsonData, "", " ")

		// Write the file back out
		os.WriteFile(source, finalData, 0644)
	},
}

var source string

func init() {
	rootCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	helloCmd.Flags().StringVar(&source, "source", "s", "Source file to read from")
}
