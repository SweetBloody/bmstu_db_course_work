package postgresql

import (
	"app/internal/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type psqlRaceResultRepository struct {
	db *sqlx.DB
}

func NewPsqlRaceResultRepository(db *sqlx.DB) models.RaceResultRepositoryI {
	return &psqlRaceResultRepository{
		db: db,
	}
}

func (pgRepo *psqlRaceResultRepository) GetAll() ([]*models.RaceResultView, error) {
	results := []*models.RaceResultView{}
	rows, err := pgRepo.db.Queryx(
		"select race_id, " +
			"coalesce(race_driver_place, 0) race_driver_place, " +
			"driver_name, team_name, gp_name " +
			"from raceresults r " +
			"join drivers d on r.driver_id = d.driver_id " +
			"join grandprix g on g.gp_id = r.gp_id " +
			"join teams t on r.team_id = t.team_id")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		res := &models.RaceResultView{}
		err = rows.StructScan(&res)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}

func (pgRepo *psqlRaceResultRepository) GetAllWithId() ([]*models.RaceResult, error) {
	results := []*models.RaceResult{}
	rows, err := pgRepo.db.Queryx(
		"select race_id, " +
			"coalesce(race_driver_place, 0) race_driver_place, " +
			"driver_id, " +
			"team_id, " +
			"gp_id " +
			"from raceresults")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		res := &models.RaceResult{}
		err = rows.StructScan(&res)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}

func (pgRepo *psqlRaceResultRepository) GetRaceResultById(id int) (*models.RaceResultView, error) {
	result := &models.RaceResultView{}
	err := pgRepo.db.Get(
		result,
		"select race_id, "+
			"coalesce(race_driver_place, 0) race_driver_place, "+
			"driver_name, team_name, gp_name "+
			"from raceresults r "+
			"join drivers d on r.driver_id = d.driver_id "+
			"join grandprix g on g.gp_id = r.gp_id "+
			"join teams t on r.team_id = t.team_id "+
			"where race_id = $1",
		id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

func (pgRepo *psqlRaceResultRepository) GetRaceResultByIdWithId(id int) (*models.RaceResult, error) {
	results := &models.RaceResult{}
	err := pgRepo.db.Get(
		results,
		"select race_id, "+
			"coalesce(race_driver_place, 0) race_driver_place, "+
			"driver_id, "+
			"team_id, "+
			"gp_id "+
			"from raceresults "+
			"where race_id = $1",
		id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return results, nil
}

func (pgRepo *psqlRaceResultRepository) GetRaceResultsOfGP(gp_id int) ([]*models.RaceResultView, error) {
	race := []*models.RaceResultView{}
	rows, err := pgRepo.db.Queryx(
		"select race_id, "+
			"coalesce(race_driver_place, 0) race_driver_place, "+
			"driver_name, team_name, gp_name "+
			"from raceresults r "+
			"join drivers d on r.driver_id = d.driver_id "+
			"join grandprix g on g.gp_id = r.gp_id "+
			"join teams t on r.team_id = t.team_id "+
			"where r.gp_id = $1",
		gp_id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for rows.Next() {
		race_temp := &models.RaceResultView{}
		err = rows.StructScan(race_temp)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		race = append(race, race_temp)
	}
	return race, nil
}

func (pgRepo *psqlRaceResultRepository) GetRaceResultsOfGPWithId(gp_id int) ([]*models.RaceResult, error) {
	race := []*models.RaceResult{}
	rows, err := pgRepo.db.Queryx(
		"select race_id, "+
			"coalesce(race_driver_place, 0) race_driver_place, "+
			"driver_id, "+
			"team_id "+
			"from raceresults "+
			"where gp_id = $1",
		gp_id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for rows.Next() {
		race_temp := &models.RaceResult{}
		err = rows.StructScan(race_temp)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		race = append(race, race_temp)
	}
	return race, nil
}

func (pgRepo *psqlRaceResultRepository) GetRaceWinnerOfGP(gp_id int) (*models.RaceResultView, error) {
	result := &models.RaceResultView{}
	err := pgRepo.db.Get(
		result,
		"select race_id, "+
			"race_driver_place, "+
			"driver_name, team_name, gp_name "+
			"from raceresults r "+
			"join drivers d on r.driver_id = d.driver_id "+
			"join grandprix g on g.gp_id = r.gp_id "+
			"join teams t on r.team_id = t.team_id "+
			"where r.gp_id = $1 and race_driver_place = 1",
		gp_id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

func (pgRepo *psqlRaceResultRepository) Create(result *models.RaceResult) (int, error) {
	var id int
	err := pgRepo.db.QueryRow(
		"insert into raceresults (race_driver_place, driver_id, team_id, gp_id) "+
			"values ($1, $2, $3, $4) "+
			"returning race_id",
		result.DriverPlace,
		result.DriverId,
		result.TeamId,
		result.GPId,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (pgRepo *psqlRaceResultRepository) Update(newResult *models.RaceResult) error {
	_, err := pgRepo.db.Exec(
		"update raceresults "+
			"set race_driver_place = $1, "+
			"driver_id = $2, "+
			"team_id = $3 "+
			"gp_id = $4 "+
			"where race_id = $4",
		newResult.DriverPlace,
		newResult.DriverId,
		newResult.TeamId,
		newResult.GPId,
		newResult.ID)
	if err != nil {
		return err
	}
	return nil
}

func (pgRepo *psqlRaceResultRepository) Delete(id int) error {
	_, err := pgRepo.db.Exec(
		"delete from raceresults "+
			"where race_id = $1",
		id)
	if err != nil {
		return err
	}
	return nil
}
