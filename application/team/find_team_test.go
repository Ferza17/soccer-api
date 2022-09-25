package team

import (
	"context"

	"github.com/Ferza17/soccer-api/model/graph/model"
	"github.com/Ferza17/soccer-api/utils"
)

func (t *testTeamApplicationSuite) TestFindTeam() {
	type (
		args struct {
			ctx   context.Context
			input model.FindTeamInputArgs
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
			name: "Err Not Found",
			args: args{
				ctx: mockCtx(context.Background()),
				input: model.FindTeamInputArgs{
					ID: "hehe",
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
				input: model.FindTeamInputArgs{
					ID: "b",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func() {
			if _, err := FindTeam(tt.args.ctx, tt.args.input); err != nil {
				t.Require().Error(err)
				t.Require().Equal(err, tt.wantErr.err)
			} else {
				t.Require().NoError(err)
			}
		})
	}
}
