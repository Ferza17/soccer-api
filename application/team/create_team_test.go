package team

import (
	"context"

	"github.com/Ferza17/soccer-api/model/graph/model"
)

func (t *testTeamApplicationSuite) TestCreateTeam() {
	type (
		args struct {
			ctx   context.Context
			input model.CreateTeam
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
				input: model.CreateTeam{
					Name:    "bang dodo",
					City:    "tangerang",
					Country: "indonesia",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func() {
			if _, err := CreateTeam(tt.args.ctx, tt.args.input); err != nil {
				t.Require().Error(err)
				t.Require().Equal(err, tt.wantErr.err)
			} else {
				t.Require().NoError(err)
			}
		})
	}
}
