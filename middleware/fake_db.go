package middleware

import (
	"context"
	"net/http"

	"github.com/Ferza17/soccer-api/repository"
	"github.com/Ferza17/soccer-api/utils"
)

func RegisterTeamsDBHTTPContext(db repository.Teams) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			ctx = context.WithValue(ctx, utils.TeamsRepositoryContextKey, db)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func RegisterPlayersDBHTTPContext(db repository.Players) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			ctx = context.WithValue(ctx, utils.PlayersRepositoryContextKey, db)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func GetTeamsFromContext(ctx context.Context) repository.Teams {
	return ctx.Value(utils.TeamsRepositoryContextKey).(repository.Teams)
}

func GetPlayersFromContext(ctx context.Context) repository.Players {
	return ctx.Value(utils.PlayersRepositoryContextKey).(repository.Players)
}
