package models

type Track struct {
	ID      int    `json:"track_id" db:"track_id"`
	Name    string `json:"track_name" db:"track_name"`
	Country string `json:"track_country" db:"track_country"`
	Town    string `json:"track_town" db:"track_town"`
}

type TrackUsecaseI interface {
	GetAll() ([]*Track, error)
	GetTeamById(id int) (*Track, error)
	Create(track *Track) (int, error)
	Update(id int, newTrack *Track) error
	Delete(id int) error
}

type TrackRepositoryI interface {
	GetAll() ([]*Track, error)
	GetTeamById(id int) (*Track, error)
	Create(track *Track) (int, error)
	Update(newTrack *Track) error
	Delete(id int) error
}
