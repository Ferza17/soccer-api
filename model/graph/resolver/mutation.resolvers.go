package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Ferza17/soccer-api/application/player"
	"github.com/Ferza17/soccer-api/application/team"
	"github.com/Ferza17/soccer-api/model/graph/generated"
	"github.com/Ferza17/soccer-api/model/graph/model"
)

// CreateTeam is the resolver for the createTeam field.
func (r *mutationResolver) CreateTeam(ctx context.Context, input model.CreateTeam) (*model.Team, error) {
	return team.CreateTeam(ctx, input)
}

// CreatePlayer is the resolver for the createPlayer field.
func (r *mutationResolver) CreatePlayer(ctx context.Context, input model.CreatePlayer) (*model.Player, error) {
	return player.CreatePlayer(ctx, input)
}

// AddPlayerToTeam is the resolver for the addPlayerToTeam field.
func (r *mutationResolver) AddPlayerToTeam(ctx context.Context, input model.AddPlayerToTeam) (*model.Team, error) {
	return team.AddPlayerToTeam(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
