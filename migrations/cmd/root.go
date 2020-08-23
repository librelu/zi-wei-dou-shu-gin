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
	"os"

	"github.com/bearners-gin/configs"
	migrations "github.com/bearners-gin/migrations/clients"
	"github.com/bearners-gin/utils/utilerrors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "migrations",
	Short: "A migration command refer to golang migration package",
	Long: `It's contains the operation refer to golang migration package such as up, down, drop ... so on.
To dealing with DB schema operation`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	initDBMigration()
	if err := rootCmd.Execute(); err != nil {
		utilerrors.Wrap(err, "migration execute error")
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&ginMode, "ginMode", "e", gin.DebugMode, "set up gin mode.")
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", migrations.DefaultConfigPath, "setup config path.")
	rootCmd.PersistentFlags().StringVarP(&fileType, "fileType", "t", migrations.YamlType, "setup config file type.")
	rootCmd.PersistentFlags().StringVarP(&migrationPath, "migrationPath", "m", migrations.DefaultMigrationPath, "setup migration files directory path.")
	rootCmd.PersistentFlags().StringVarP(&migrationDriver, "migrationDriver", "d", migrations.PostgresDriverName, "setup migration driver.")
}

func initDBMigration() {
	os.Setenv("GIN_MODE", ginMode)
	config, err := configs.NewConfigs(configPath, fileType)
	if err != nil {
		panic(utilerrors.Wrap(err, "getting config error"))
	}
	dbMigration, err = migrations.NewDBMigration(
		config.DB.Host, config.DB.User, config.DB.Password, config.DB.Database, config.DB.SSLMode,
		migrationPath, migrationDriver, config.DB.Port,
	)
	if err != nil {
		panic(utilerrors.Wrap(err, "connect to db error"))
	}
}
