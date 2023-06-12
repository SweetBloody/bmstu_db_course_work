package postgresql

import (
	"app/internal/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type psqlTeamRepository struct {
	db *sqlx.DB
}

func NewPsqlTeamRepository(db *sqlx.DB) models.TeamRepositoryI {
	return &psqlTeamRepository{
		db: db,
	}
}

func (pgRepo *psqlTeamRepository) GetAll() ([]*models.Team, error) {
	teams := []*models.Team{}
	rows, err := pgRepo.db.Queryx(
		"select team_id, team_name, team_country, team_base " +
			"from teams")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		team := &models.Team{}
		err := rows.StructScan(&team)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}

func (pgRepo *psqlTeamRepository) GetTeamById(id int) (*models.Team, error) {
	teams := &models.Team{}
	err := pgRepo.db.Get(
		teams,
		"select team_id, team_name, team_country, team_base "+
			"from teams "+
			"where team_id = $1",
		id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return teams, nil
}

func (pgRepo *psqlTeamRepository) GetTeamsOfSeason(season int) ([]*models.Team, error) {
	teams := []*models.Team{}
	rows, err := pgRepo.db.Queryx(
		"select t.team_id, team_name, team_country, team_base "+
			"from teams t "+
			"join teamsdrivers td on t.team_id = td.team_id "+
			"where team_driver_season = $1 "+
			"group by t.team_id",
		season)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		team := &models.Team{}
		err = rows.StructScan(&team)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}

func (pgRepo *psqlTeamRepository) Create(team *models.Team) (int, error) {
	var id int
	err := pgRepo.db.QueryRow(
		"insert into teams (team_name, team_country, team_base) "+
			"values ($1, $2, $3) "+
			"returning team_id",
		team.Name,
		team.Country,
		team.Base,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (pgRepo *psqlTeamRepository) Update(newTeam *models.Team) error {
	_, err := pgRepo.db.Exec(
		"update teams "+
			"set team_name = $1, "+
			"team_country = $2, "+
			"team_base = $3 "+
			"where team_id = $4",
		newTeam.Name,
		newTeam.Country,
		newTeam.Base,
		newTeam.ID)
	if err != nil {
		return err
	}
	return nil
}

func (pgRepo *psqlTeamRepository) Delete(id int) error {
	_, err := pgRepo.db.Exec(
		"delete from teams "+
			"where team_id = $1",
		id)
	if err != nil {
		return err
	}
	return nil
}
