package team

import (
	"context"

	"github.com/Ferza17/soccer-api/model/graph/model"
)

func (t *testTeamApplicationSuite) TestFindTeams() {
	type (
		args struct {
			ctx   context.Context
			input model.FindTeamsInputArgs
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
			name: "success no input ids",
			args: args{
				ctx: mockCtx(context.Background()),
			},
		},
		{
			name: "success with input ids",
			args: args{
				ctx: mockCtx(context.Background()),
				input: model.FindTeamsInputArgs{
					Ids: []string{"b"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func() {
			if _, err := FindTeams(tt.args.ctx, tt.args.input); err != nil {
				t.Require().Error(err)
				t.Require().Equal(err, tt.wantErr.err)
			} else {
				t.Require().NoError(err)
			}
		})
	}
}
