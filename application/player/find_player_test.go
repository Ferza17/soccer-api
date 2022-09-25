package player

import (
	"context"

	"github.com/Ferza17/soccer-api/model/graph/model"
	"github.com/Ferza17/soccer-api/utils"
)

func (t *testPlayerApplicationSuite) TestFindPlayer() {
	type (
		args struct {
			ctx   context.Context
			input model.FindPlayerInputArgs
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
			name: "ErrNotFound",
			args: args{
				ctx: mockCtx(context.Background()),
				input: model.FindPlayerInputArgs{
					ID:           "",
					IsFreePlayer: false,
				},
			},
			wantErr: wantErr{
				err: utils.ErrNotFound,
			},
		},
		{
			name: "Success in team",
			args: args{
				ctx: mockCtx(context.Background()),
				input: model.FindPlayerInputArgs{
					ID:           "c",
					IsFreePlayer: false,
				},
			},
			wantErr: wantErr{
				err: utils.ErrNotFound,
			},
		},
		{
			name: "Success no team",
			args: args{
				ctx: mockCtx(context.Background()),
				input: model.FindPlayerInputArgs{
					ID:           "a",
					IsFreePlayer: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func() {
			if _, err := FindPlayer(tt.args.ctx, tt.args.input); err != nil {
				t.Require().Error(err)
				t.Require().Equal(err, tt.wantErr.err)
			} else {
				t.Require().NoError(err)
			}
		})
	}
}
