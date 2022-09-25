package utils

type ContextKey string

const (
	HeadersContextKey           ContextKey = "headers_context_key"
	HostInfoContextKey          ContextKey = "host_info_context_key"
	TeamsRepositoryContextKey   ContextKey = "teams_repository_context_key"
	PlayersRepositoryContextKey ContextKey = "players_repository_context_key"
)
