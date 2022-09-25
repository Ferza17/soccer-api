package cmd

import (
	"context"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Ferza17/soccer-api/repository"
)

var (
	logger    *zap.Logger
	playersDB repository.Players
	teamsDB   repository.Teams
)

func init() {
	logger = NewLogger()
	playersDB = NewPlayersDB()
	teamsDB = NewTeamsDB()
}

func NewLogger() (logger *zap.Logger) {
	var err error
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	if logger, err = logConfig.Build(); err != nil {
		panic(err)
	}
	return
}

func NewTeamsDB() repository.Teams {
	db := make(repository.Teams)
	return db
}

func NewPlayersDB() repository.Players {
	db := make(repository.Players)
	return db
}

func Shutdown(ctx context.Context) (err error) {
	log.Println("Shutdown...")
	return
}
