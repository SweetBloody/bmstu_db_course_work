package usecase

import (
	"app/internal/pkg/models"
)

type driverUsecase struct {
	driverRepo models.DriverRepositoryI
}

func NewDriverUsecase(driverRepo models.DriverRepositoryI) models.DriverUsecaseI {
	return &driverUsecase{
		driverRepo: driverRepo,
	}
}

func (uc *driverUsecase) GetAll() ([]*models.Driver, error) {
	drivers, err := uc.driverRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return drivers, nil
}

func (uc *driverUsecase) GetDriverById(id int) (*models.Driver, error) {
	driver, err := uc.driverRepo.GetDriverById(id)
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func (uc *driverUsecase) GetDriversOfSeason(season int) ([]*models.Driver, error) {
	drivers, err := uc.driverRepo.GetDriversOfSeason(season)
	if err != nil {
		return nil, err
	}
	return drivers, nil
}

func (uc *driverUsecase) GetDriversStanding() ([]*models.Standings, error) {
	standings, err := uc.driverRepo.GetDriversStanding()
	if err != nil {
		return nil, err
	}
	return standings, nil
}

func (uc *driverUsecase) Create(driver *models.Driver) (int, error) {
	id, err := uc.driverRepo.Create(driver)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc *driverUsecase) Update(id int, newDriver *models.Driver) error {
	newDriver.ID = id
	err := uc.driverRepo.Update(newDriver)
	if err != nil {
		return err
	}
	return nil
}

func (uc *driverUsecase) Delete(id int) error {
	err := uc.driverRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *driverUsecase) LinkDriverTeam(new *models.DriversTeams) error {
	err := uc.driverRepo.LinkDriverTeam(new)
	if err != nil {
		return err
	}
	return nil
}
