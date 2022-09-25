package team

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/Ferza17/soccer-api/model/graph/model"
	"github.com/Ferza17/soccer-api/repository"
	"github.com/Ferza17/soccer-api/utils"
)

type testTeamApplicationSuite struct {
	suite.Suite
}

func TestTeamApplication(t *testing.T) {
	suite.Run(t, new(testTeamApplicationSuite))
}

func mockCtx(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.PlayersRepositoryContextKey, repository.Players{
		"a": {
			ID:        "a",
			Name:      "a",
			Number:    1,
			Country:   "a",
			CreatedAt: 1,
			UpdatedAt: 1,
		},
	})
	ctx = context.WithValue(ctx, utils.TeamsRepositoryContextKey, repository.Teams{
		"b": {
			ID:      "b",
			Name:    "b",
			City:    "b",
			Country: "b",
			Players: []*model.Player{
				{
					ID:        "c",
					Name:      "c",
					Number:    1,
					Country:   "c",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			CreatedAt: 1,
			UpdatedAt: 1,
		},
	})
	return ctx
}

func (t *testTeamApplicationSuite) TestAddPlayerToTeam() {
	type (
		args struct {
			ctx   context.Context
			input model.AddPlayerToTeam
		}
		wantResponse struct {
		}
		wantErr struct {
			err error
		}
	)
	tests := []struct {
		name         string
		args         args
		wantResponse wantResponse
		wantErr      wantErr
	}{
		{
			name: "Err Team Not Found",
			args: args{
				ctx: mockCtx(context.Background()),
				input: model.AddPlayerToTeam{
					TeamID:    "hehe",
					PlayerIds: nil,
				},
			},
			wantErr: wantErr{
				err: utils.ErrNotFound,
			},
		},
		{
			name: "Err Player Not Found",
			args: args{
				ctx: mockCtx(context.Background()),
				input: model.AddPlayerToTeam{
					TeamID:    "b",
					PlayerIds: []string{"hehe"},
				},
			},
			wantErr: wantErr{
				err: utils.ErrNotFound,
			},
		},
		{
			name: "Success",
			args: args{
				ctx: mockCtx(context.Background()),
				input: model.AddPlayerToTeam{
					TeamID:    "b",
					PlayerIds: []string{"a"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func() {
			if _, err := AddPlayerToTeam(tt.args.ctx, tt.args.input); err != nil {
				t.Require().Error(err)
				t.Require().Equal(err, tt.wantErr.err)
			} else {
				t.Require().NoError(err)
			}
		})
	}
}
