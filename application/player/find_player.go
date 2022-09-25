package player

import (
	"context"

	"github.com/Ferza17/soccer-api/middleware"
	"github.com/Ferza17/soccer-api/model/graph/model"
	"github.com/Ferza17/soccer-api/utils"
)

func FindPlayer(ctx context.Context, input model.FindPlayerInputArgs) (*model.Player, error) {
	var (
		playerDB = middleware.GetPlayersFromContext(ctx)
		teamDB   = middleware.GetTeamsFromContext(ctx)
		response *model.Player
		ok       bool
	)
	if !input.IsFreePlayer {
		for _, team := range teamDB {
			for _, player := range team.Players {
				if player.ID == input.ID {
					return player, nil
				}
			}
		}
	}

	response, ok = playerDB[input.ID]
	if !ok {
		return nil, utils.ErrNotFound
	}
	return response, nil
}
