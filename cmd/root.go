// Copyright © 2018 MTRNord <info@nordgedanken.de>
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

	"github.com/Nordgedanken/Morpheusv2/pkg/app"
	"github.com/Nordgedanken/Morpheusv2/pkg/db/sqlite"
	"github.com/Nordgedanken/Morpheusv2/pkg/util"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/shibukawa/configdir"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"path/filepath"
	"time"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "Morpheusv2",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Init Logs and folders
		// configdir.New gets the Config Paths that are available on this system
		configDirs := configdir.New("Nordgedanken", "Morpheusv2")
		// path holds the global configdir of the system with the log folder appended
		path := filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path + "/log/")
		logFilePath := filepath.ToSlash(path + "main.log")

		// This checks if the directory is missing and if that's the case to generate the directory
		if _, StatErr := os.Stat(path); os.IsNotExist(StatErr) {
			MkdirErr := os.MkdirAll(path, os.ModeDir)
			if MkdirErr != nil {
				return MkdirErr
			}
		}

		writer, err := rotatelogs.New(
			logFilePath+".%Y%m%d%H%M",
			rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
			rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
		)
		if err != nil {
			return err
		}

		log.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.InfoLevel:  writer,
				logrus.ErrorLevel: writer,
			},
			&logrus.TextFormatter{},
		))

		// dbImpl.Init() generates the needed tables if needed before the app starts
		util.DB = &sqlite.SQLite{}
		err = util.DB.Init()
		if err != nil {
			return err
		}

		log.Println("DB Set Up")
		return nil
	},
	// TODO add descriptions
	Short: "",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.Start(args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
