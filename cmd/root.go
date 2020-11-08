// Copyright © 2020 Dmitry Stoletov <info@imega.ru>
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
	"regexp"

	"github.com/imega/luaformatter/formatter"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	config  = formatter.DefaultConfig()
	write   bool

	rootCmd = &cobra.Command{
		Use:     "luaformatter",
		Version: "1.0.0",
		Short:   "A brief description of your application",
		Long:    `A longer description `,
		Example: `    Formatting file on stdout:

        $ luaformatter path/to/script.lua

  Overwrite file with formatting

    $ luaformatter -w path/to/script.lua
  `,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			w := os.Stdout

			// if write {
			// 	w = os.Stdout
			// }

			fs := afero.NewRegexpFs(
				afero.NewOsFs(),
				regexp.MustCompile(`\.lua$`),
			)

			for _, arg := range args {
				isDir, _ := afero.IsDir(fs, arg)
				// if err != nil {
				// 	fmt.Printf("is dir %s\n", err)
				// 	continue
				// }

				if isDir {
					files, err := afero.ReadDir(fs, arg)
					if err != nil {
						fmt.Printf("read_dir %s\n", err)

						continue
					}
					for _, f := range files {
						b, err := afero.ReadFile(fs, f.Name())
						if err != nil {
							fmt.Printf("read file in dir %s\n", err)

							continue
						}
						err = formatter.Format(formatter.Config{}, b, w)
						fmt.Printf("%s\n", err)
					}
				}

				b, err := afero.ReadFile(fs, arg)
				if err != nil {
					fmt.Printf("read file %s\n", err)

					continue
				}

				if err := formatter.Format(config, b, w); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		},
	}
)

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

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.luaformatter.yaml)",
	)

	rootCmd.Flags().BoolVarP(
		&write,
		"write",
		"w",
		false,
		"write result to (source) file instead of stdout",
	)
}

func initConfig() {
	v := viper.New()
	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		v.AddConfigPath(home)
		v.SetConfigName(".luaformatter")
	}

	if err := v.ReadInConfig(); err == nil {
		if err := v.Unmarshal(&config); err != nil {
			fmt.Printf("Error decoding config %s\n%s\n", v.ConfigFileUsed(), err)
			os.Exit(1)
		}
	}
}
