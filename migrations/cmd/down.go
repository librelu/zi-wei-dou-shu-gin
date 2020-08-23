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
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Reverse one or multiple steps of migrations",
	Run: func(cmd *cobra.Command, args []string) {
		if steps != 0 && steps > 0 {
			steps = -steps
			if err := dbMigration.Steps(steps); err != nil {
				panic(err)
			}
		} else if steps == 0 {
			if err := dbMigration.Down(); err != nil {
				if errors.Is(err, migrate.ErrNoChange) {
					fmt.Println(migrate.ErrNoChange)
				} else {
					panic(err)
				}
			}
		} else {
			panic("steps can't be negative")
		}
		fmt.Println("Down command completed!!")
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
	rootCmd.PersistentFlags().IntVarP(&steps, "steps", "s", 0, "set up steps")
}
