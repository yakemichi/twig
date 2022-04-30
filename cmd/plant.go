/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// plantCmd represents the plant command
var plantCmd = &cobra.Command{
	Use:   "plant",
	Short: "Summon a new seedling",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("planting the seed....")

		createDataDirectory()

		fmt.Println("succesfully planted your little seedling!")
	},
}


func getDataDirectory() string {
	dirname, err := os.UserHomeDir()
	
	if err != nil {
		log.Fatal( err )
    }
	
	return path.Join(dirname, ".twig-cli")
}


func createDataDirectory() {
	pathname := getDataDirectory()
	
    fmt.Println(pathname)

	if _, err := os.Stat(pathname); os.IsNotExist(err) {
		os.Mkdir(pathname, os.ModeDir|0755)
	}
}


func init() {
	rootCmd.AddCommand(plantCmd)
}