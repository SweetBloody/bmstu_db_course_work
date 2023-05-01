package postgresql

import (
	"app/internal/pkg/models"
	"fmt"
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
		err := rows.StructScan(&driver)
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
		fmt.Println(err)
		return nil, err
	}
	return driver, nil
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
	fmt.Println("errrr === ", err)
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
