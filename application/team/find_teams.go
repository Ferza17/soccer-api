package team

import (
	"context"

	"github.com/Ferza17/soccer-api/middleware"
	"github.com/Ferza17/soccer-api/model/graph/model"
)

func FindTeams(ctx context.Context, input model.FindTeamsInputArgs) ([]*model.Team, error) {
	var (
		teamDB   = middleware.GetTeamsFromContext(ctx)
		response []*model.Team
	)

	if len(input.Ids) < 1 {
		for _, team := range teamDB {
			response = append(response, team)
		}
		return response, nil
	}

	for _, id := range input.Ids {
		response = append(response, teamDB[id])
	}
	return response, nil
}
