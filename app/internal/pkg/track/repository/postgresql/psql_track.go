package postgresql

import (
	"app/internal/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type psqlTrackRepository struct {
	db *sqlx.DB
}

func NewPsqlTrackRepository(db *sqlx.DB) models.TrackRepositoryI {
	return &psqlTrackRepository{
		db: db,
	}
}

func (pgRepo *psqlTrackRepository) GetAll() ([]*models.Track, error) {
	tracks := []*models.Track{}
	rows, err := pgRepo.db.Queryx(
		"select track_id, track_name, track_country, track_town " +
			"from tracks")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		track := &models.Track{}
		err := rows.StructScan(&track)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	return tracks, nil
}

func (pgRepo *psqlTrackRepository) GetTeamById(id int) (*models.Track, error) {
	tracks := &models.Track{}
	err := pgRepo.db.Get(
		tracks,
		"select track_id, track_name, track_country, track_town "+
			"from tracks "+
			"where track_id = $1",
		id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return tracks, nil
}

func (pgRepo *psqlTrackRepository) Create(track *models.Track) (int, error) {
	var id int
	err := pgRepo.db.QueryRow(
		"insert into tracks (track_name, track_country, track_town) "+
			"values ($1, $2, $3) "+
			"returning track_id",
		track.Name,
		track.Country,
		track.Town,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (pgRepo *psqlTrackRepository) Update(newTrack *models.Track) error {
	_, err := pgRepo.db.Exec(
		"update tracks "+
			"set track_name = $1, "+
			"track_country = $2, "+
			"track_base = $3 "+
			"where track_id = $4",
		newTrack.Name,
		newTrack.Country,
		newTrack.Town,
		newTrack.ID)
	if err != nil {
		return err
	}
	return nil
}

func (pgRepo *psqlTrackRepository) Delete(id int) error {
	_, err := pgRepo.db.Exec(
		"delete from tracks "+
			"where track_id = $1",
		id)
	if err != nil {
		return err
	}
	return nil
}
