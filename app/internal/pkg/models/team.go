package models

type Team struct {
	ID      int    `json:"team_id" db:"team_id"`
	Name    string `json:"team_name" db:"team_name"`
	Country string `json:"team_country" db:"team_country"`
	Base    string `json:"team_base" db:"team_base"`
}

type TeamUsecaseI interface {
	GetAll() ([]*Team, error)
	GetTeamById(id int) (*Team, error)
	GetTeamsOfSeason(season int) ([]*Team, error)
	Create(team *Team) (int, error)
	Update(id int, newTeam *Team) error
	Delete(id int) error
}

type TeamRepositoryI interface {
	GetAll() ([]*Team, error)
	GetTeamById(id int) (*Team, error)
	GetTeamsOfSeason(season int) ([]*Team, error)
	Create(team *Team) (int, error)
	Update(newTeam *Team) error
	Delete(id int) error
}
