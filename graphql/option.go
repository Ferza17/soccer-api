package graphql

import (
	"go.uber.org/zap"

	"github.com/Ferza17/soccer-api/repository"
)

func NewLogger(logger *zap.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func NewPlayersDB(db repository.Players) Option {
	return func(s *Server) {
		s.playersDB = db
	}
}

func NewTeamsDB(db repository.Teams) Option {
	return func(s *Server) {
		s.teamsDB = db
	}
}
