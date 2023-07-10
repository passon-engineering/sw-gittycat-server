package application

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
	"github.com/passon-engineering/sw-go-utility-lib/networking"
)

type Application struct {
	ServerPath string
	SystemIP   string
	Logger     *logger.Logger
}

func Init(app *Application) error {
	startTime := time.Now()
	var err error

	//capture global server path
	app.ServerPath = filepath.Dir(os.Args[0])

	appLogger, err := logger.NewLogger(
		[]logger.LogFormat{
			logger.FORMAT_TIMESTAMP,
			logger.FORMAT_STATUS,
			logger.FORMAT_INFO,
			logger.FORMAT_PRE_TEXT,
			logger.FORMAT_HTTP_REQUEST,
			logger.FORMAT_ID,
			logger.FORMAT_SOURCE,
			logger.FORMAT_DATA,
			logger.FORMAT_ERROR,
			logger.FORMAT_PROCESSING_TIME,
		}, logger.Options{
			OutputToStdout:   true,
			OutputToFile:     true,
			OutputFolderPath: "/var/log/gittycat-server/",
		}, logger.Container{
			Status: logger.STATUS_INFO,
			Info:   "System Logger succesfully started! Awaiting logger tasks...",
		})
	if err != nil {
		log.Fatalf(err.Error())
	}

	ip, err := networking.GetNetworkExternalIP()
	if err != nil {
		appLogger.Entry(logger.Container{
			Status: logger.STATUS_ERROR,
			Error:  "Could not get network external IP: " + err.Error(),
		})
		log.Fatalf("Could not get network external IP: %v", err)
	}
	appLogger.Entry(logger.Container{
		Status: logger.STATUS_INFO,
		Info:   "Network external IP: " + ip,
	})

	app.SystemIP = ip

	appLogger.Entry(logger.Container{
		Status:         logger.STATUS_INFO,
		Info:           "Basic app framework sucessfully initialized",
		ProcessingTime: time.Since(startTime),
	})
	return nil
}
