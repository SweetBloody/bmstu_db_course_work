package postgresql

import (
	"app/internal/pkg/models"
	"github.com/jmoiron/sqlx"
)

type psqlDriverRepository struct {
	db *sqlx.DB
}

func NewPsqlDriverRepository(db *sqlx.DB) models.DriverRepositoryI {
	return &psqlDriverRepository{
		db: db,
	}
}

func (pgRepo *psqlDriverRepository) GetAll() ([]*models.Driver, error) {
	drivers := []*models.Driver{}
	rows, err := pgRepo.db.Queryx(
		"select driver_id, driver_name, driver_country, driver_birth_date " +
			"from drivers")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		driver := &models.Driver{}
		err = rows.StructScan(&driver)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func (pgRepo *psqlDriverRepository) GetDriverById(id int) (*models.Driver, error) {
	driver := &models.Driver{}
	err := pgRepo.db.Get(
		driver,
		"select driver_id, driver_name, driver_country, driver_birth_date "+
			"from drivers "+
			"where driver_id = $1",
		id)
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func (pgRepo *psqlDriverRepository) GetDriversOfSeason(season int) ([]*models.Driver, error) {
	drivers := []*models.Driver{}
	rows, err := pgRepo.db.Queryx(
		"select d.driver_id, driver_name, driver_country, driver_birth_date "+
			"from drivers d "+
			"join teamsdrivers t on d.driver_id = t.driver_id "+
			"where team_driver_season = $1",
		season)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		driver := &models.Driver{}
		err = rows.StructScan(&driver)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func (pgRepo *psqlDriverRepository) GetDriversStanding() ([]*models.Standings, error) {
	standings := []*models.Standings{}
	rows, err := pgRepo.db.Queryx(
		"select st_id, season, driver_name, team_name, score " +
			"from season_standings s " +
			"join drivers d on s.driver_id = d.driver_id " +
			"join teams t on s.team_id = t.team_id " +
			"order by score desc")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := &models.Standings{}
		err = rows.StructScan(&temp)
		if err != nil {
			return nil, err
		}
		standings = append(standings, temp)
	}
	return standings, nil
}

func (pgRepo *psqlDriverRepository) Create(driver *models.Driver) (int, error) {
	var id int
	err := pgRepo.db.QueryRow(
		"insert into drivers (driver_name, driver_country, driver_birth_date) "+
			"values ($1, $2, $3) "+
			"returning driver_id",
		driver.Name,
		driver.Country,
		driver.BirthDate,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (pgRepo *psqlDriverRepository) Update(newDriver *models.Driver) error {
	_, err := pgRepo.db.Exec(
		"update drivers "+
			"set driver_name = $1, "+
			"driver_country = $2, "+
			"driver_birth_date = $3 "+
			"where driver_id = $4",
		newDriver.Name,
		newDriver.Country,
		newDriver.BirthDate,
		newDriver.ID)
	if err != nil {
		return err
	}
	return nil
}

func (pgRepo *psqlDriverRepository) Delete(id int) error {
	_, err := pgRepo.db.Exec(
		"delete from drivers "+
			"where driver_id = $1",
		id)
	if err != nil {
		return err
	}
	return nil
}

func (pgRepo *psqlDriverRepository) LinkDriverTeam(new *models.DriversTeams) error {
	_, err := pgRepo.db.Exec(
		"insert into teamsdrivers (driver_id, team_id, team_driver_season) "+
			"values ($1, $2, $3)",
		new.DriverId,
		new.TeamId,
		new.Season,
	)
	if err != nil {
		return err
	}
	return nil
}
