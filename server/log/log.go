package log

import (
	"os"

	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func init() {
	var err error
	l, err := zap.NewDevelopment()
	Log = l.Sugar()

	if err != nil {
		os.Exit(1)
	}

	defer Log.Sync()
}
