package models

type RaceResult struct {
	ID          int `json:"race_id" db:"race_id"`
	DriverPlace int `json:"race_driver_place" db:"race_driver_place"`
	DriverId    int `json:"race_driver_id" db:"driver_id"`
	TeamId      int `json:"race_team_id" db:"team_id"`
	GPId        int `json:"gp_id" db:"gp_id"`
}

type RaceResultView struct {
	ID          int    `json:"race_id" db:"race_id"`
	DriverPlace int    `json:"race_driver_place" db:"race_driver_place"`
	DriverName  string `json:"driver_name" db:"driver_name"`
	TeamName    string `json:"team_name" db:"team_name"`
	GPName      string `json:"gp_name" db:"gp_name"`
}

type RaceResultUsecaseI interface {
	GetAll() ([]*RaceResultView, error)
	GetAllWithId() ([]*RaceResult, error)
	GetRaceResultById(id int) (*RaceResultView, error)
	GetRaceResultByIdWithId(id int) (*RaceResult, error)
	GetRaceResultsOfGP(gp_id int) ([]*RaceResultView, error)
	GetRaceResultsOfGPWithId(gp_id int) ([]*RaceResult, error)
	GetRaceWinnerOfGP(gp_id int) (*RaceResultView, error)
	Create(result *RaceResult) (int, error)
	Update(id int, newResult *RaceResult) error
	Delete(id int) error
}

type RaceResultRepositoryI interface {
	GetAll() ([]*RaceResultView, error)
	GetAllWithId() ([]*RaceResult, error)
	GetRaceResultById(id int) (*RaceResultView, error)
	GetRaceResultByIdWithId(id int) (*RaceResult, error)
	GetRaceResultsOfGP(gp_id int) ([]*RaceResultView, error)
	GetRaceResultsOfGPWithId(gp_id int) ([]*RaceResult, error)
	GetRaceWinnerOfGP(gp_id int) (*RaceResultView, error)
	Create(result *RaceResult) (int, error)
	Update(newResult *RaceResult) error
	Delete(id int) error
}
