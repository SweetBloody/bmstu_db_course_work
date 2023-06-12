package usecase

import (
	"app/internal/pkg/models"
)

type teamUsecase struct {
	teamRepo models.TeamRepositoryI
}

func NewTeamUsecase(teamRepo models.TeamRepositoryI) models.TeamUsecaseI {
	return &teamUsecase{
		teamRepo: teamRepo,
	}
}

func (uc *teamUsecase) GetAll() ([]*models.Team, error) {
	teams, err := uc.teamRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (uc *teamUsecase) GetTeamById(id int) (*models.Team, error) {
	teams, err := uc.teamRepo.GetTeamById(id)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (uc *teamUsecase) GetTeamsOfSeason(season int) ([]*models.Team, error) {
	drivers, err := uc.teamRepo.GetTeamsOfSeason(season)
	if err != nil {
		return nil, err
	}
	return drivers, nil
}

func (uc *teamUsecase) Create(team *models.Team) (int, error) {
	id, err := uc.teamRepo.Create(team)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc *teamUsecase) Update(id int, newTeam *models.Team) error {
	newTeam.ID = id
	err := uc.teamRepo.Update(newTeam)
	if err != nil {
		return err
	}
	return nil
}

func (uc *teamUsecase) Delete(id int) error {
	err := uc.teamRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
