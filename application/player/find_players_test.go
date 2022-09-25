package player

import (
	"context"

	"github.com/Ferza17/soccer-api/model/graph/model"
)

func (t *testPlayerApplicationSuite) TestFindPlayers() {
	type (
		args struct {
			ctx   context.Context
			input model.FindPlayersInputArgs
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
			name: "Success in team",
			args: args{
				ctx: mockCtx(context.Background()),
				input: model.FindPlayersInputArgs{
					Ids:          nil,
					IsFreePlayer: false,
				},
			},
		},
		{
			name: "Success in team with ids",
			args: args{
				ctx: mockCtx(context.Background()),
				input: model.FindPlayersInputArgs{
					Ids:          []string{"c"},
					IsFreePlayer: false,
				},
			},
		},
		{
			name: "Success no team",
			args: args{
				ctx: mockCtx(context.Background()),
				input: model.FindPlayersInputArgs{
					Ids:          nil,
					IsFreePlayer: true,
				},
			},
		},
		{
			name: "Success no team with ids",
			args: args{
				ctx: mockCtx(context.Background()),
				input: model.FindPlayersInputArgs{
					Ids:          []string{"a"},
					IsFreePlayer: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func() {
			if _, err := FindPlayers(tt.args.ctx, tt.args.input); err != nil {
				t.Require().Error(err)
				t.Require().Equal(err, tt.wantErr.err)
			} else {
				t.Require().NoError(err)
			}
		})
	}
}
