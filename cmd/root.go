/*
Copyright © 2020 Kris Nova <kris@nivenly.com>

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
	"os"

	"github.com/kris-nova/logger"
	"github.com/kris-nova/tz0rk/bot"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tzork",
	Short: "A Zork like Twitter Bot",
	Long:  `A twitter bot that responds to all tweets containing /tz0rk`,
	Run: func(cmd *cobra.Command, args []string) {

		logger.Always("Welcome to Tz0rk!")
		b := bot.New(devMode)
		err := b.Auth()
		if err != nil {
			logger.Critical(err.Error())
			os.Exit(1)
		}
		errch := make(chan error)
		go b.Run(errch)
		for {
			err := <-errch
			if err != nil {
				logger.Warning(err.Error())
			}
		}
		logger.Always("Bye!")
		os.Exit(0)
	},
}

var devMode bool

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&devMode, "dev", "d", false, "enable dev mode which bypasses the cache autoloading")
	rootCmd.Flags().IntVarP(&logger.Level, "verbosity", "v", 4, "verbosity level 0 (low) 4 (high)")
}
