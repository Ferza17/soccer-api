package player

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/Ferza17/soccer-api/middleware"
	"github.com/Ferza17/soccer-api/model/graph/model"
)

func CreatePlayer(ctx context.Context, input model.CreatePlayer) (*model.Player, error) {
	var (
		playerDB = middleware.GetPlayersFromContext(ctx)
	)
	player := model.Player{
		ID:        uuid.New().String(),
		Name:      input.Name,
		Number:    input.Number,
		Country:   input.Country,
		CreatedAt: int(time.Now().Unix()),
		UpdatedAt: int(time.Now().Unix()),
	}
	playerDB[player.ID] = &player
	return &player, nil
}
