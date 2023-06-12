package usecase

import (
	"app/internal/pkg/models"
)

type trackUsecase struct {
	trackRepo models.TrackRepositoryI
}

func NewTrackUsecase(trackRepo models.TrackRepositoryI) models.TrackUsecaseI {
	return &trackUsecase{
		trackRepo: trackRepo,
	}
}

func (uc *trackUsecase) GetAll() ([]*models.Track, error) {
	tracks, err := uc.trackRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (uc *trackUsecase) GetTeamById(id int) (*models.Track, error) {
	tracks, err := uc.trackRepo.GetTeamById(id)
	if err != nil {
		return nil, err
	}
	return tracks, nil
}

func (uc *trackUsecase) Create(team *models.Track) (int, error) {
	id, err := uc.trackRepo.Create(team)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc *trackUsecase) Update(id int, newTrack *models.Track) error {
	newTrack.ID = id
	err := uc.trackRepo.Update(newTrack)
	if err != nil {
		return err
	}
	return nil
}

func (uc *trackUsecase) Delete(id int) error {
	err := uc.trackRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
