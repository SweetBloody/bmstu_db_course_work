package models

type Driver struct {
	ID        int    `json:"driver_id" db:"driver_id"`
	Name      string `json:"driver_name" db:"driver_name"`
	Country   string `json:"driver_country" db:"driver_country"`
	BirthDate string `json:"driver_birth_date" db:"driver_birth_date"`
}

type DriverUsecaseI interface {
	GetAll() ([]*Driver, error)
	GetDriverById(id int) (*Driver, error)
	Create(driver *Driver) (int, error)
	Update(id int, newDriver *Driver) error
	Delete(id int) error
}

type DriverRepositoryI interface {
	GetAll() ([]*Driver, error)
	GetDriverById(id int) (*Driver, error)
	Create(driver *Driver) (int, error)
	Update(newDriver *Driver) error
	Delete(id int) error
}
