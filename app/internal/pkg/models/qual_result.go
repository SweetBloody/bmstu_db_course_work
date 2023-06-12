package models

import "time"

type QualResult struct {
	ID          int       `json:"qual_id" db:"qual_id"`
	DriverPlace int       `json:"qual_driver_place" db:"qual_driver_place"`
	DriverId    int       `json:"qual_driver_id" db:"driver_id"`
	TeamId      int       `json:"qual_team_id" db:"team_id"`
	Q1time      time.Time `json:"q1_time" db:"q1_time"`
	Q2time      time.Time `json:"q2_time" db:"q2_time"`
	Q3time      time.Time `json:"q3_time" db:"q3_time"`
	GPId        int       `json:"gp_id" db:"gp_id"`
}

type QualResultView struct {
	ID          int       `json:"qual_id" db:"qual_id"`
	DriverPlace int       `json:"qual_driver_place" db:"qual_driver_place"`
	DriverName  string    `json:"driver_name" db:"driver_name"`
	TeamName    string    `json:"team_name" db:"team_name"`
	Q1time      time.Time `json:"q1_time" db:"q1_time"`
	Q2time      time.Time `json:"q2_time" db:"q2_time"`
	Q3time      time.Time `json:"q3_time" db:"q3_time"`
	GPName      string    `json:"gp_name" db:"gp_name"`
}

type QualResultUsecaseI interface {
	GetAll() ([]*QualResultView, error)
	GetAllWithId() ([]*QualResult, error)
	GetQualResultById(id int) (*QualResultView, error)
	GetQualResultByIdWithId(id int) (*QualResult, error)
	GetQualResultsOfGP(gp_id int) ([]*QualResultView, error)
	GetQualResultsOfGPWithId(gp_id int) ([]*QualResult, error)
	Create(result *QualResult) (int, error)
	Update(id int, newResult *QualResult) error
	Delete(id int) error
}

type QualResultRepositoryI interface {
	GetAll() ([]*QualResultView, error)
	GetAllWithId() ([]*QualResult, error)
	GetQualResultById(id int) (*QualResultView, error)
	GetQualResultByIdWithId(id int) (*QualResult, error)
	GetQualResultsOfGP(gp_id int) ([]*QualResultView, error)
	GetQualResultsOfGPWithId(gp_id int) ([]*QualResult, error)
	Create(result *QualResult) (int, error)
	Update(newResult *QualResult) error
	Delete(id int) error
}
