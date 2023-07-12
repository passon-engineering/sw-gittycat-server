package application

import (
	"io/ioutil"
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

func Init() *Application {
	startTime := time.Now()
	var err error

	app := &Application{
		ServerPath: filepath.Dir(os.Args[0]),
		SystemIP:   "",
		Logger:     &logger.Logger{},
	}

	//capture global server path
	app.ServerPath = filepath.Dir(os.Args[0])

	app.Logger, err = logger.NewLogger(
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
		log.Fatalf("Could not initialize logger: %v", err)
	}

	app.Logger.Entry(logger.Container{
		Status:         logger.STATUS_INFO,
		Info:           "Server path: " + app.ServerPath,
		ProcessingTime: time.Since(startTime),
	})

	ip, err := networking.GetNetworkExternalIP()
	if err != nil {
		app.Logger.Entry(logger.Container{
			Status: logger.STATUS_ERROR,
			Error:  "Could not get network external IP: " + err.Error(),
		})
		os.Exit(1)
	}
	app.Logger.Entry(logger.Container{
		Status: logger.STATUS_INFO,
		Info:   "Network external IP: " + ip,
	})

	app.SystemIP = ip

	publicSshKey, err := getUserSshPublicKey()
	if err != nil {
		app.Logger.Entry(logger.Container{
			Status: logger.STATUS_ERROR,
			Error:  "Could not read user public SSH key: " + err.Error(),
		})
		os.Exit(1)
	} else {
		app.Logger.Entry(logger.Container{
			Status: logger.STATUS_INFO,
			Info:   "User public SSH key:",
		})
		app.Logger.Entry(logger.Container{
			Status: logger.STATUS_INFO,
			Info:   string(publicSshKey),
		})
	}

	app.Logger.Entry(logger.Container{
		Status:         logger.STATUS_INFO,
		Info:           "Basic app framework sucessfully initialized",
		ProcessingTime: time.Since(startTime),
	})

	if err != nil {
		log.Fatalf("Could not initialize app! %v", err)
	}

	return app
}

func getUserSshPublicKey() ([]byte, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		// Unable to get user's home directory
		return nil, err
	}

	keyPath := filepath.Join(homeDir, ".ssh", "id_rsa.pub")

	keyData, err := ioutil.ReadFile(keyPath)
	if err != nil {
		// Unable to read file
		return nil, err
	}

	return keyData, nil
}
