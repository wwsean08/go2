// Copyright © 2017 Sean Smith <sean@wwsean08.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"database/sql"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wwsean08/go2/handler"
	"log"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go2",
	Short: "A webservice for redirecting based on simple words",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := getDBConn()
		if err != nil {
			log.Panicf("Error when connecting to database, aborting startup: %v", err)
		}
		router := handler.SetupHandlers(conn)
		address := fmt.Sprintf("0.0.0.0:%d", viper.GetInt("port"))
		err = router.Start(address)
		if err != nil {
			errorMsg := fmt.Sprintf("Got the following error starting the server: %v", err.Error())
			fmt.Println(errorMsg)
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/go2/config.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")                                                        // name of config file (without extension)
	viper.AddConfigPath(fmt.Sprintf("$HOME/%s", viper.GetString("application.pkgName"))) // adding home directory as first search path
	viper.AddConfigPath(fmt.Sprintf("/etc/%s", viper.GetString("application.pkgName")))
	viper.AutomaticEnv() // read in environment variables that match

	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func getDBConn() (*sql.DB, error) {
	viper.SetDefault("postgres.sslmode", "disable")
	viper.SetDefault("postgres.port", "5432")

	//Get needed variables
	database := viper.GetString("postgres.database")
	user := viper.GetString("postgres.user")
	pass := viper.GetString("postgres.password")
	host := viper.GetString("postgres.host")
	port := viper.GetInt("postgres.port")
	SSLMode := viper.GetString("postgres.sslmode")

	//Initialize and setup connection
	connString := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%d sslmode=%s",
		database, user, pass, host, port, SSLMode)
	var err error
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}
