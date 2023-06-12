package usecase

import (
	"app/internal/pkg/models"
)

type grandPrixUsecase struct {
	grandPrixRepo models.GrandPrixRepositoryI
}

func NewGrandPrixUsecase(grandPrix models.GrandPrixRepositoryI) models.GrandPrixUsecaseI {
	return &grandPrixUsecase{
		grandPrixRepo: grandPrix,
	}
}

func (uc *grandPrixUsecase) GetAll() ([]*models.GrandPrix, error) {
	gp, err := uc.grandPrixRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return gp, nil
}

func (uc *grandPrixUsecase) GetAllBySeason(season int) ([]*models.GrandPrix, error) {
	gp, err := uc.grandPrixRepo.GetAllBySeason(season)
	if err != nil {
		return nil, err
	}
	return gp, nil
}

func (uc *grandPrixUsecase) GetAllByPlace(place string) ([]*models.GrandPrix, error) {
	gp, err := uc.grandPrixRepo.GetAllByPlace(place)
	if err != nil {
		return nil, err
	}
	return gp, nil
}

func (uc *grandPrixUsecase) GetGPById(id int) (*models.GrandPrix, error) {
	gp, err := uc.grandPrixRepo.GetGPById(id)
	if err != nil {
		return nil, err
	}
	return gp, nil
}

func (uc *grandPrixUsecase) Create(grandPrix *models.GrandPrix) (int, error) {
	id, err := uc.grandPrixRepo.Create(grandPrix)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc *grandPrixUsecase) Update(id int, newGrandPrix *models.GrandPrix) error {
	newGrandPrix.ID = id
	err := uc.grandPrixRepo.Update(newGrandPrix)
	if err != nil {
		return err
	}
	return nil
}

func (uc *grandPrixUsecase) Delete(id int) error {
	err := uc.grandPrixRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

//func (uc *grandPrixUsecase) GetRaceResultsOfGP(gp_id int) ([]*models.RaceResult, error) {
//	race_results, err := uc.grandPrixRepo.GetRaceResultsOfGP(gp_id)
//	if err != nil {
//		return nil, err
//	}
//	return race_results, nil
//}
//
//func (uc *grandPrixUsecase) GetQualResultsOfGP(gp_id int) ([]*models.QualResult, error) {
//	qual_results, err := uc.grandPrixRepo.GetQualResultsOfGP(gp_id)
//	if err != nil {
//		return nil, err
//	}
//	return qual_results, nil
//}
