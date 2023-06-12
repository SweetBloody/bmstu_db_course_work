package postgresql

import (
	"app/internal/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type psqlGPRepository struct {
	db *sqlx.DB
}

func NewPsqlGPRepository(db *sqlx.DB) models.GrandPrixRepositoryI {
	return &psqlGPRepository{
		db: db,
	}
}

func (pgRepo *psqlGPRepository) GetAll() ([]*models.GrandPrix, error) {
	gp := []*models.GrandPrix{}
	rows, err := pgRepo.db.Queryx(
		"select gp_id, gp_season, gp_name, gp_date_num, gp_month, gp_place, gp_track_id " +
			"from grandprix")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		gp_temp := &models.GrandPrix{}
		err = rows.StructScan(gp_temp)
		if err != nil {
			return nil, err
		}
		gp = append(gp, gp_temp)
	}
	return gp, nil
}

func (pgRepo *psqlGPRepository) GetAllBySeason(season int) ([]*models.GrandPrix, error) {
	gp := []*models.GrandPrix{}
	rows, err := pgRepo.db.Queryx(
		"select gp_id, gp_season, gp_name, gp_date_num, gp_month, gp_place, gp_track_id "+
			"from grandprix "+
			"where gp_season = $1",
		season)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		gp_temp := &models.GrandPrix{}
		err = rows.StructScan(gp_temp)
		if err != nil {
			return nil, err
		}
		gp = append(gp, gp_temp)
	}
	return gp, nil
}

func (pgRepo *psqlGPRepository) GetAllByPlace(place string) ([]*models.GrandPrix, error) {
	gp := []*models.GrandPrix{}
	rows, err := pgRepo.db.Queryx(
		"select gp_id, gp_season, gp_name, gp_date_num, gp_month, gp_place, gp_track_id "+
			"from grandprix "+
			"where gp_place = $1",
		place)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		gp_temp := &models.GrandPrix{}
		err = rows.StructScan(gp_temp)
		if err != nil {
			return nil, err
		}
		gp = append(gp, gp_temp)
	}
	return gp, nil
}

func (pgRepo *psqlGPRepository) GetGPById(id int) (*models.GrandPrix, error) {
	gp := &models.GrandPrix{}
	err := pgRepo.db.Get(
		gp,
		"select gp_id, gp_season, gp_name, gp_date_num, gp_month, gp_place, gp_track_id "+
			"from grandprix "+
			"where gp_id = $1",
		id)
	if err != nil {
		return nil, err
	}
	return gp, nil
}

func (pgRepo *psqlGPRepository) Create(grandPrix *models.GrandPrix) (int, error) {
	var id int
	err := pgRepo.db.QueryRow(
		"insert into grandprix (gp_season, gp_name, gp_date_num, gp_month, gp_place, gp_track_id)"+
			"values ($1, $2, $3, $4, $5, $6) "+
			"returning gp_id",
		grandPrix.Season,
		grandPrix.Name,
		grandPrix.DateNum,
		grandPrix.Month,
		grandPrix.Place,
		grandPrix.TrackId,
	).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, nil
}

func (pgRepo *psqlGPRepository) Update(newGrandPrix *models.GrandPrix) error {
	_, err := pgRepo.db.Exec(
		"update grandprix "+
			"set gp_season = $1, "+
			"gp_date_num = $2, "+
			"gp_month = $3,"+
			"gp_place = $4,"+
			"gp_track_id = $5 "+
			"where driver_id = $6",
		newGrandPrix.Name,
		newGrandPrix.DateNum,
		newGrandPrix.Month,
		newGrandPrix.Place,
		newGrandPrix.TrackId,
		newGrandPrix.ID)
	if err != nil {
		return err
	}
	return nil
}

func (pgRepo *psqlGPRepository) Delete(id int) error {
	_, err := pgRepo.db.Exec(
		"delete from grandprix "+
			"where gp_id = $1",
		id)
	if err != nil {
		return err
	}
	return nil
}

//func (pgRepo *psqlGPRepository) GetRaceResultsOfGP(gp_id int) ([]*models.RaceResult, error) {
//	race := []*models.RaceResult{}
//	rows, err := pgRepo.db.Queryx(
//		"select race_id, "+
//			"coalesce(race_driver_place, 0) race_driver_place, "+
//			"driver_id, "+
//			"team_id "+
//			"from raceresults "+
//			"where gp_id = $1",
//		gp_id)
//	if err != nil {
//		fmt.Println(err)
//		return nil, err
//	}
//	for rows.Next() {
//		race_temp := &models.RaceResult{}
//		err = rows.StructScan(race_temp)
//		if err != nil {
//			fmt.Println(err)
//			return nil, err
//		}
//		race = append(race, race_temp)
//	}
//	return race, nil
//}
//
//func (pgRepo *psqlGPRepository) GetQualResultsOfGP(gp_id int) ([]*models.QualResult, error) {
//	qual := []*models.QualResult{}
//	rows, err := pgRepo.db.Queryx(
//		"select qual_id, "+
//			"coalesce(qual_driver_place, 0) qual_driver_place, "+
//			"driver_id, team_id, "+
//			"coalesce(q1_time, '00:00:00') q1_time, "+
//			"coalesce(q2_time, '00:00:00') q2_time, "+
//			"coalesce(q3_time, '00:00:00') q3_time "+
//			"from qualificationresults "+
//			"where gp_id = $1",
//		gp_id)
//	if err != nil {
//
//		return nil, err
//	}
//	for rows.Next() {
//		qual_temp := &models.QualResult{}
//		err = rows.StructScan(qual_temp)
//		if err != nil {
//			return nil, err
//		}
//		qual = append(qual, qual_temp)
//	}
//	return qual, nil
//}
