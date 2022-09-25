package player

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/Ferza17/soccer-api/model/graph/model"
	"github.com/Ferza17/soccer-api/repository"
	"github.com/Ferza17/soccer-api/utils"
)

type testPlayerApplicationSuite struct {
	suite.Suite
}

func TestPlayerApplication(t *testing.T) {
	suite.Run(t, new(testPlayerApplicationSuite))
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

func (t *testPlayerApplicationSuite) TestCreatePlayer() {
	type (
		args struct {
			ctx   context.Context
			input model.CreatePlayer
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
			name: "Success",
			args: args{
				ctx: mockCtx(context.Background()),
				input: model.CreatePlayer{
					Name:    "Jadon Sancho",
					Number:  25,
					Country: "England",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func() {
			if _, err := CreatePlayer(tt.args.ctx, tt.args.input); err != nil {
				t.Require().Error(err)
				t.Require().Equal(err, tt.wantErr.err)
			} else {
				t.Require().NoError(err)
			}
		})
	}
}
