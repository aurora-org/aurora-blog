package common

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func SetupConfig() {

	conf := flag.String("conf", "./config/runtime.config", "Config file path")
	flag.Parse()

	viper.SetConfigName(*conf)
	viper.SetConfigType("json")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Read runtime config fail: %v\n", err.Error())
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
}
