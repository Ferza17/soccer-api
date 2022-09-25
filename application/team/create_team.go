package team

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/Ferza17/soccer-api/middleware"
	"github.com/Ferza17/soccer-api/model/graph/model"
)

func CreateTeam(ctx context.Context, input model.CreateTeam) (*model.Team, error) {
	var (
		teamDB = middleware.GetTeamsFromContext(ctx)
	)
	team := model.Team{
		ID:        uuid.New().String(),
		Name:      input.Name,
		City:      input.City,
		Country:   input.Country,
		CreatedAt: int(time.Now().Unix()),
		UpdatedAt: int(time.Now().Unix()),
	}
	teamDB[team.ID] = &team
	return &team, nil
}
