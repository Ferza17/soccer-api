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

// FindTeams is the resolver for the findTeams field.
func (r *queryResolver) FindTeams(ctx context.Context, input model.FindTeamsInputArgs) ([]*model.Team, error) {
	return team.FindTeams(ctx, input)
}

// FindTeam is the resolver for the findTeam field.
func (r *queryResolver) FindTeam(ctx context.Context, input model.FindTeamInputArgs) (*model.Team, error) {
	return team.FindTeam(ctx, input)
}

// FindPlayers is the resolver for the findPlayers field.
func (r *queryResolver) FindPlayers(ctx context.Context, input model.FindPlayersInputArgs) ([]*model.Player, error) {
	return player.FindPlayers(ctx, input)
}

// FindPlayer is the resolver for the findPlayer field.
func (r *queryResolver) FindPlayer(ctx context.Context, input model.FindPlayerInputArgs) (*model.Player, error) {
	return player.FindPlayer(ctx, input)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
