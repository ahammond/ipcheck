// Copyright © 2018 NAME HERE andrew.george.hammond@gmail.com
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

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"bufio"
	"github.com/ahammond/ipcheck/pkg"
)

var cfgFile string
var outputFormat string

func mainRun(cmd *cobra.Command, args []string) {
	stdout := bufio.NewWriter(os.Stdout)
	defer stdout.Flush()
	out := ipcheck.ProcessList(outputFormat, &args)
	out.WriteTo(stdout)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ipcheck",
	Short: "Given an IP, describe it as either public, private, loopback, multicast, etc.",
	Long:  ``,
	Run:   mainRun,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ipcheck.yaml)")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "text", "output format (default is text. options include text, json)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".ipcheck" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ipcheck")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
