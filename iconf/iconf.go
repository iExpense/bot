package iconf

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	cConfType         = "yaml"
	cConfFile         = "botconf"
	cConfPath         = "/etc/bot"
	cLogFile          = "bot.log"
	cFilePerm         = 640
	cLogFileMaxSize   = 50 // MB
	cLogMaxNumBackups = 5
	cLogFileMaxAge    = 30 // days
)

func Init() error {
	// Read configuration file
	viper.SetConfigType(cConfType)
	viper.SetConfigName(cConfFile)
	viper.AddConfigPath(cConfPath)
	viper.AddConfigPath(".") // for development
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("[ERROR] unable to read configuration file")
		panic(err)
	}

	// Setup logging
	return setupLogging()
}

func setupLogging() error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logFolder := getLogFolder()
	_, err := os.Stat(logFolder)
	if err != nil {
		if strings.Contains(err.Error(), "cannot find the") ||
			strings.Contains(err.Error(), "no such file or directory") {

			errDir := os.MkdirAll(logFolder, os.FileMode(cFilePerm))
			if errDir != nil {
				log.Println("[ERROR] unable to create app folder")
				panic(errDir)
			} else {
				log.Println("[INFO] created log directory")
			}
		} else {
			log.Println("[ERROR] unable to get app folder stats")
			panic(err)
		}
	} else {
		log.Println("[INFO] log directory already exists")
	}

	logFilePath := filepath.Join(logFolder, cLogFile)
	log.SetOutput(&lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    cLogFileMaxSize,
		MaxBackups: cLogMaxNumBackups,
		MaxAge:     cLogFileMaxAge,
		LocalTime:  true,
	})

	log.Println("#################### BEGIN OF LOG ##########################")
	return nil
}

func getLogFolder() string {
	return viper.GetString("logfolder")
}
