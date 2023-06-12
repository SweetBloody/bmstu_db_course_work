package usecase

import (
	"app/internal/pkg/models"
)

type raceResultUsecase struct {
	raceResultRepo models.RaceResultRepositoryI
}

func NewRaceResultUsecase(raceResultRepo models.RaceResultRepositoryI) models.RaceResultUsecaseI {
	return &raceResultUsecase{
		raceResultRepo: raceResultRepo,
	}
}

func (uc *raceResultUsecase) GetAll() ([]*models.RaceResultView, error) {
	results, err := uc.raceResultRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (uc *raceResultUsecase) GetAllWithId() ([]*models.RaceResult, error) {
	results, err := uc.raceResultRepo.GetAllWithId()
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (uc *raceResultUsecase) GetRaceResultById(id int) (*models.RaceResultView, error) {
	result, err := uc.raceResultRepo.GetRaceResultById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *raceResultUsecase) GetRaceResultByIdWithId(id int) (*models.RaceResult, error) {
	result, err := uc.raceResultRepo.GetRaceResultByIdWithId(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *raceResultUsecase) GetRaceResultsOfGP(gp_id int) ([]*models.RaceResultView, error) {
	race_results, err := uc.raceResultRepo.GetRaceResultsOfGP(gp_id)
	if err != nil {
		return nil, err
	}
	return race_results, nil
}

func (uc *raceResultUsecase) GetRaceResultsOfGPWithId(gp_id int) ([]*models.RaceResult, error) {
	race_results, err := uc.raceResultRepo.GetRaceResultsOfGPWithId(gp_id)
	if err != nil {
		return nil, err
	}
	return race_results, nil
}

func (uc *raceResultUsecase) GetRaceWinnerOfGP(gp_id int) (*models.RaceResultView, error) {
	race_result, err := uc.raceResultRepo.GetRaceWinnerOfGP(gp_id)
	if err != nil {
		return nil, err
	}
	return race_result, nil
}

func (uc *raceResultUsecase) Create(driver *models.RaceResult) (int, error) {
	id, err := uc.raceResultRepo.Create(driver)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc *raceResultUsecase) Update(id int, newResult *models.RaceResult) error {
	newResult.ID = id
	err := uc.raceResultRepo.Update(newResult)
	if err != nil {
		return err
	}
	return nil
}

func (uc *raceResultUsecase) Delete(id int) error {
	err := uc.raceResultRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
