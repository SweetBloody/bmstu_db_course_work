package postgresql

import (
	"app/internal/pkg/models"
	"github.com/jmoiron/sqlx"
)

type psqlQualResultRepository struct {
	db *sqlx.DB
}

func NewPsqlQualResultRepository(db *sqlx.DB) models.QualResultRepositoryI {
	return &psqlQualResultRepository{
		db: db,
	}
}

func (pgRepo *psqlQualResultRepository) GetAll() ([]*models.QualResultView, error) {
	results := []*models.QualResultView{}
	rows, err := pgRepo.db.Queryx(
		"select qual_id, " +
			"coalesce(qual_driver_place, 0) qual_driver_place, " +
			"driver_name, team_name, " +
			"coalesce(q1_time, '00:00:00') q1_time, " +
			"coalesce(q2_time, '00:00:00') q2_time, " +
			"coalesce(q3_time, '00:00:00') q3_time, " +
			"gp_name " +
			"from qualificationresults q " +
			"join drivers d on q.driver_id = d.driver_id " +
			"join teams t on q.team_id = t.team_id " +
			"join grandprix g on g.gp_id = q.gp_id")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		res := &models.QualResultView{}
		err = rows.StructScan(&res)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}

func (pgRepo *psqlQualResultRepository) GetAllWithId() ([]*models.QualResult, error) {
	results := []*models.QualResult{}
	rows, err := pgRepo.db.Queryx(
		"select qual_id, " +
			"coalesce(qual_driver_place, 0) qual_driver_place, " +
			"driver_id, team_id, " +
			"coalesce(q1_time, '00:00:00') q1_time, " +
			"coalesce(q2_time, '00:00:00') q2_time, " +
			"coalesce(q3_time, '00:00:00') q3_time, " +
			"gp_id " +
			"from qualificationresults")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		res := &models.QualResult{}
		err = rows.StructScan(&res)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}

func (pgRepo *psqlQualResultRepository) GetQualResultById(id int) (*models.QualResultView, error) {
	results := &models.QualResultView{}
	err := pgRepo.db.Get(
		results,
		"select qual_id, "+
			"coalesce(qual_driver_place, 0) qual_driver_place, "+
			"driver_name, team_name, "+
			"coalesce(q1_time, '00:00:00') q1_time, "+
			"coalesce(q2_time, '00:00:00') q2_time, "+
			"coalesce(q3_time, '00:00:00') q3_time, "+
			"gp_name "+
			"from qualificationresults q "+
			"join drivers d on q.driver_id = d.driver_id "+
			"join teams t on q.team_id = t.team_id "+
			"join grandprix g on g.gp_id = q.gp_id"+
			"where qual_id = $1",
		id)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (pgRepo *psqlQualResultRepository) GetQualResultByIdWithId(id int) (*models.QualResult, error) {
	result := &models.QualResult{}
	err := pgRepo.db.Get(
		result,
		"select qual_id, "+
			"coalesce(qual_driver_place, 0) qual_driver_place, "+
			"driver_id, team_id, "+
			"coalesce(q1_time, '00:00:00') q1_time, "+
			"coalesce(q2_time, '00:00:00') q2_time, "+
			"coalesce(q3_time, '00:00:00') q3_time, "+
			"gp_id "+
			"from qualresults "+
			"where qual_id = $1",
		id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pgRepo *psqlQualResultRepository) GetQualResultsOfGP(gp_id int) ([]*models.QualResultView, error) {
	qual := []*models.QualResultView{}
	rows, err := pgRepo.db.Queryx(
		"select qual_id, "+
			"coalesce(qual_driver_place, 0) qual_driver_place, "+
			"driver_name, team_name, gp_name, "+
			"coalesce(q1_time, '00:00:00') q1_time, "+
			"coalesce(q2_time, '00:00:00') q2_time, "+
			"coalesce(q3_time, '00:00:00') q3_time "+
			"from qualificationresults q "+
			"join drivers d on q.driver_id = d.driver_id "+
			"join teams t on q.team_id = t.team_id "+
			"join grandprix g on g.gp_id = q.gp_id "+
			"where q.gp_id = $1",
		gp_id)
	if err != nil {

		return nil, err
	}
	for rows.Next() {
		qual_temp := &models.QualResultView{}
		err = rows.StructScan(qual_temp)
		if err != nil {
			return nil, err
		}
		qual = append(qual, qual_temp)
	}
	return qual, nil
}

func (pgRepo *psqlQualResultRepository) GetQualResultsOfGPWithId(gp_id int) ([]*models.QualResult, error) {
	qual := []*models.QualResult{}
	rows, err := pgRepo.db.Queryx(
		"select qual_id, "+
			"coalesce(qual_driver_place, 0) qual_driver_place, "+
			"driver_id, team_id, "+
			"coalesce(q1_time, '00:00:00') q1_time, "+
			"coalesce(q2_time, '00:00:00') q2_time, "+
			"coalesce(q3_time, '00:00:00') q3_time "+
			"from qualificationresults "+
			"where gp_id = $1",
		gp_id)
	if err != nil {

		return nil, err
	}
	for rows.Next() {
		qual_temp := &models.QualResult{}
		err = rows.StructScan(qual_temp)
		if err != nil {
			return nil, err
		}
		qual = append(qual, qual_temp)
	}
	return qual, nil
}

func (pgRepo *psqlQualResultRepository) Create(result *models.QualResult) (int, error) {
	var id int
	err := pgRepo.db.QueryRow(
		"insert into qualresults (qual_driver_place, driver_id, team_id, q1_tim1, q2_time, q3_time, gp_id) "+
			"values ($1, $2, $3, $4, $5, $6, $7) "+
			"returning qual_id",
		result.DriverPlace,
		result.DriverId,
		result.TeamId,
		result.Q1time,
		result.Q2time,
		result.Q3time,
		result.GPId,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (pgRepo *psqlQualResultRepository) Update(newResult *models.QualResult) error {
	_, err := pgRepo.db.Exec(
		"update qualresults "+
			"set qual_driver_place = $1, "+
			"driver_id = $2, "+
			"team_id = $3 "+
			"q1_time = $4 "+
			"q2_time = $5 "+
			"q3_time = $6 "+
			"gp_id = $7 "+
			"where qual_id = $4",
		newResult.DriverPlace,
		newResult.DriverId,
		newResult.TeamId,
		newResult.Q1time,
		newResult.Q2time,
		newResult.Q3time,
		newResult.GPId,
		newResult.ID)
	if err != nil {
		return err
	}
	return nil
}

func (pgRepo *psqlQualResultRepository) Delete(id int) error {
	_, err := pgRepo.db.Exec(
		"delete from qualresults "+
			"where qual_id = $1",
		id)
	if err != nil {
		return err
	}
	return nil
}
