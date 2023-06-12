package postgresql

import (
	"app/internal/pkg/models"
	"github.com/jmoiron/sqlx"
)

type psqlUserRepository struct {
	db *sqlx.DB
}

func NewPsqlUserRepository(db *sqlx.DB) models.UserRepositoryI {
	return &psqlUserRepository{
		db: db,
	}
}

func (pgRepo *psqlUserRepository) GetUserById(id int) (*models.User, error) {
	user := &models.User{}
	err := pgRepo.db.Get(
		user,
		"select user_id, login, password, role "+
			"from users "+
			"where user_id = $1",
		id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (pgRepo *psqlUserRepository) GetUserByLogin(login string) (*models.User, error) {
	user := &models.User{}
	err := pgRepo.db.Get(
		user,
		"select user_id, login, password, role "+
			"from users "+
			"where login = $1",
		login)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (pgRepo *psqlUserRepository) Create(user *models.User) (int, error) {
	var id int
	err := pgRepo.db.QueryRow(
		"insert into users (login, password, role) "+
			"values ($1, $2, $3) "+
			"returning user_id",
		user.Login,
		user.Password,
		user.Role,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (pgRepo *psqlUserRepository) Update(newUser *models.User) error {
	_, err := pgRepo.db.Exec(
		"update users "+
			"set login = $1, "+
			"password = $2, "+
			"role = $3 "+
			"where user_id = $4",
		newUser.Login,
		newUser.Password,
		newUser.Role,
		newUser.ID)
	if err != nil {
		return err
	}
	return nil
}

func (pgRepo *psqlUserRepository) Delete(id int) error {
	_, err := pgRepo.db.Exec(
		"delete from users "+
			"where user_id = $1",
		id)
	if err != nil {
		return err
	}
	return nil
}
