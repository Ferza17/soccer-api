package team

import (
	"context"

	"github.com/Ferza17/soccer-api/middleware"
	"github.com/Ferza17/soccer-api/model/graph/model"
	"github.com/Ferza17/soccer-api/utils"
)

func FindTeam(ctx context.Context, input model.FindTeamInputArgs) (*model.Team, error) {
	var (
		teamDB = middleware.GetTeamsFromContext(ctx)
	)
	response, ok := teamDB[input.ID]
	if !ok {
		return nil, utils.ErrNotFound
	}
	return response, nil
}
