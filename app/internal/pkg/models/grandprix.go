package models

type GrandPrix struct {
	ID      int    `json:"gp_id" db:"gp_id"`
	Season  int    `json:"gp_season" db:"gp_season"`
	Name    string `json:"gp_name" db:"gp_name"`
	DateNum int    `json:"gp_date_num" db:"gp_date_num"`
	Month   string `json:"gp_month" db:"gp_month"`
	Place   string `json:"gp_place" db:"gp_place"`
	TrackId int    `json:"gp_track_id" db:"gp_track_id"`
}

type GrandPrixUsecaseI interface {
	GetAll() ([]*GrandPrix, error)
	GetGPById(id int) (*GrandPrix, error)
	GetAllBySeason(season int) ([]*GrandPrix, error)
	GetAllByPlace(place string) ([]*GrandPrix, error)
	Create(grandPrix *GrandPrix) (int, error)
	Update(id int, newGrandPrix *GrandPrix) error
	Delete(id int) error
}

type GrandPrixRepositoryI interface {
	GetAll() ([]*GrandPrix, error)
	GetGPById(id int) (*GrandPrix, error)
	GetAllBySeason(season int) ([]*GrandPrix, error)
	GetAllByPlace(place string) ([]*GrandPrix, error)
	Create(grandPrix *GrandPrix) (int, error)
	Update(newGrandPrix *GrandPrix) error
	Delete(id int) error
}
