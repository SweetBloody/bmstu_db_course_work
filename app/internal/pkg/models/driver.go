package models

type Driver struct {
	ID        int    `json:"driver_id" db:"driver_id"`
	Name      string `json:"driver_name" db:"driver_name"`
	Country   string `json:"driver_country" db:"driver_country"`
	BirthDate string `json:"driver_birth_date" db:"driver_birth_date"`
}

type Standings struct {
	ID         int    `json:"st_id" db:"st_id"`
	Season     int    `json:"season" db:"season"`
	DriverName string `json:"driver_name" db:"driver_name"`
	TeamName   string `json:"team_name" db:"team_name"`
	Score      int    `json:"score" db:"score"`
}

type DriversTeams struct {
	DriverId int `json:"driver_id" db:"driver_id"`
	TeamId   int `json:"team_id" db:"team_id"`
	Season   int `json:"team_driver_season" db:"team_driver_season"`
}

type DriverUsecaseI interface {
	GetAll() ([]*Driver, error)
	GetDriverById(id int) (*Driver, error)
	GetDriversOfSeason(season int) ([]*Driver, error)
	GetDriversStanding() ([]*Standings, error)
	Create(driver *Driver) (int, error)
	Update(id int, newDriver *Driver) error
	Delete(id int) error
	LinkDriverTeam(new *DriversTeams) error
}

type DriverRepositoryI interface {
	GetAll() ([]*Driver, error)
	GetDriverById(id int) (*Driver, error)
	GetDriversOfSeason(season int) ([]*Driver, error)
	GetDriversStanding() ([]*Standings, error)
	Create(driver *Driver) (int, error)
	Update(newDriver *Driver) error
	Delete(id int) error
	LinkDriverTeam(new *DriversTeams) error
}
