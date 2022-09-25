package team

import (
	"context"

	"github.com/Ferza17/soccer-api/middleware"
	"github.com/Ferza17/soccer-api/model/graph/model"
	"github.com/Ferza17/soccer-api/utils"
)

func AddPlayerToTeam(ctx context.Context, input model.AddPlayerToTeam) (*model.Team, error) {
	var (
		teamDB   = middleware.GetTeamsFromContext(ctx)
		playerDB = middleware.GetPlayersFromContext(ctx)
		team     = &model.Team{}
		ok       bool
	)
	if team, ok = teamDB[input.TeamID]; !ok {
		return nil, utils.ErrNotFound
	}
	for _, id := range input.PlayerIds {
		if _, ok = playerDB[id]; !ok {
			return nil, utils.ErrNotFound
		}
		team.Players = append(team.Players, playerDB[id])
		delete(playerDB, id)
	}
	return team, nil
}
