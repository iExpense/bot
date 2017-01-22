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
	cFilePerm         = 0777
	cLogFileMaxSize   = 50 // MB
	cLogMaxNumBackups = 5
	cLogFileMaxAge    = 30 // days
)

func Init() error {
	viper.SetConfigType(cConfType)
	viper.SetConfigName(cConfFile)
	viper.AddConfigPath(cConfPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("[WARN] no configuration file loaded")
	}

	// enable reading from environment variable
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Setup logging
	return setupLogging()
}

func setupLogging() error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if isLogsToStdout() {
		return nil
	}

	logFolder := getLogFolder()
	_, err := os.Stat(logFolder)
	if err != nil {
		if strings.Contains(err.Error(), "cannot find the") ||
			strings.Contains(err.Error(), "no such file or directory") {

			errDir := os.MkdirAll(logFolder, os.FileMode(cFilePerm))
			if errDir != nil {
				log.Println("[ERROR] unable to create log folder")
				panic(errDir)
			} else {
				log.Println("[INFO] created log directory")
			}
		} else {
			log.Println("[ERROR] unable to get log folder stats")
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

func isLogsToStdout() bool {
	logToStdout := viper.GetString("log_stdout")
	return !(logToStdout == "" || logToStdout == "false")
}

func getLogFolder() string {
	return viper.GetString("log_folder")
}
