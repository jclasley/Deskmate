package log

import (
	"os"

	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	var err error
	Log, err = zap.NewDevelopment()
	Log.Sugar()

	if err != nil {
		os.Exit(1)
	}

	defer Log.Sync()
}
