// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Player struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Number    int    `json:"number"`
	Country   string `json:"country"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}

type Team struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	Players   []*Player `json:"players"`
	CreatedAt int       `json:"createdAt"`
	UpdatedAt int       `json:"updatedAt"`
}

type AddPlayerToTeam struct {
	TeamID    string   `json:"teamId"`
	PlayerIds []string `json:"playerIds"`
}

type CreatePlayer struct {
	Name    string `json:"name"`
	Number  int    `json:"number"`
	Country string `json:"country"`
}

type CreateTeam struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type FindPlayerInputArgs struct {
	ID           string `json:"id"`
	IsFreePlayer bool   `json:"isFreePlayer"`
}

type FindPlayersInputArgs struct {
	Ids          []string `json:"ids"`
	IsFreePlayer bool     `json:"isFreePlayer"`
}

type FindTeamInputArgs struct {
	ID string `json:"id"`
}

type FindTeamsInputArgs struct {
	Ids []string `json:"ids"`
}

type MapPlayer struct {
	ID *Player `json:"id"`
}
