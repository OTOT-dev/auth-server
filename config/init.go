package config

import (
	"auth-server/common"
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

func init() {
	newFlagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	newFlagSet.StringVar(&cfgFile, "config", "", "config file (default is $EXEC/default.json)")
	newFlagSet.StringVar(&ServerHost, "host", "127.0.0.1", "server host, default is 127.0.0.1")
	newFlagSet.IntVar(&ServerPort, "port", 8000, "server port, default is 8000")
	newFlagSet.BoolVar(&DebugMode, "debug", false, "debug mode ,default is false")
	if err := newFlagSet.Parse(os.Args[1:]); err != nil {
		log.Fatalf("flag parse error:%s", err.Error())
	}
	initViper()
	initConfig()
}

func initViper() {

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(common.GetExecPath())
		viper.SetConfigName("default.json")
	}
	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}
