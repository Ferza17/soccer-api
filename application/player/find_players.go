package player

import (
	"context"

	"github.com/Ferza17/soccer-api/middleware"
	"github.com/Ferza17/soccer-api/model/graph/model"
)

func FindPlayers(ctx context.Context, input model.FindPlayersInputArgs) ([]*model.Player, error) {
	var (
		playerDB = middleware.GetPlayersFromContext(ctx)
		teamDB   = middleware.GetTeamsFromContext(ctx)
		response []*model.Player
	)

	if !input.IsFreePlayer {
		if len(input.Ids) > 0 {
			for _, team := range teamDB {
				for _, player := range team.Players {
					for _, id := range input.Ids {
						if player.ID == id {
							response = append(response, player)
						}
					}
				}
			}
			return response, nil
		}

		for _, team := range teamDB {
			for _, player := range team.Players {
				response = append(response, player)
			}
		}
		return response, nil
	}

	if len(input.Ids) < 1 {
		for _, player := range playerDB {
			response = append(response, player)
		}
		return response, nil
	}

	for _, id := range input.Ids {
		response = append(response, playerDB[id])
	}
	return response, nil
}
