package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zlobste/spotter/internal/data"
)

const (
	usersTable = "users"
	all        = "*"
)

type UsersStorage interface {
	New() UsersStorage
	Get() (*data.User, error)
	GetUserById(id uint64) (*data.User, error)
	GetUserByEmail(email string) (*data.User, error)
	CreateUser(user data.User) error
	UpdateUser(id uint64, user data.User) error
	DeleteUser(id uint64) error
	GetAllManagers() ([]data.User, error)
	GetAllDrivers() ([]data.User, error)
	SetManager(id uint64) error
	BlockUser(id uint64) error
	UnblockUser(id uint64) error
}

type userStorage struct {
	db  *sql.DB
	sql sq.SelectBuilder
}

var usersSelect = sq.Select(all).From(usersTable).PlaceholderFormat(sq.Dollar)

func (s *userStorage) New() UsersStorage {
	return NewUsersStorage(s.db)
}

func NewUsersStorage(db *sql.DB) UsersStorage {
	return &userStorage{
		db:  db,
		sql: usersSelect.RunWith(db),
	}
}

func (s *userStorage) Get() (*data.User, error) {
	rowScanner := s.sql.QueryRow()
	model := data.User{}
	err := rowScanner.Scan(
		&model.Id,
		&model.Name,
		&model.Surname,
		&model.Email,
		&model.Password,
		&model.Role,
		&model.Blocked,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "failed to query model")
	} else if err == sql.ErrNoRows {
		return nil, nil
	}
	return &model, nil
}

func (s *userStorage) Select() ([]data.User, error) {
	rows, err := s.sql.RunWith(s.db).Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var models []data.User

	for rows.Next() {
		model := data.User{}
		err := rows.Scan(
			&model.Id,
			&model.Name,
			&model.Surname,
			&model.Email,
			&model.Password,
			&model.Role,
			&model.Blocked,
		)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}

func (s *userStorage) GetUserById(id uint64) (*data.User, error) {
	s.sql = s.sql.Where(sq.Eq{"id": id})
	return s.Get()
}

func (s *userStorage) GetUserByEmail(email string) (*data.User, error) {
	s.sql = s.sql.Where(sq.Eq{"email": email})
	return s.Get()
}

func (s *userStorage) newInsert() sq.InsertBuilder {
	return sq.Insert(usersTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *userStorage) CreateUser(user data.User) error {
	_, err := s.newInsert().SetMap(user.ToMap()).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to insert user")
	}
	return nil
}

func (s *userStorage) newUpdate() sq.UpdateBuilder {
	return sq.Update(usersTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *userStorage) UpdateUser(id uint64, user data.User) error {
	_, err := s.newUpdate().SetMap(user.ToMap()).Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to update user data")
	}
	return nil
}

func (s *userStorage) newDelete() sq.DeleteBuilder {
	return sq.Delete(usersTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *userStorage) DeleteUser(id uint64) error {
	_, err := s.newDelete().Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}
	return nil
}

func (s *userStorage) GetAllManagers() ([]data.User, error) {
	s.sql = s.sql.Where(sq.Eq{"role": data.RoleTypeManager})
	return s.Select()
}

func (s *userStorage) GetAllDrivers() ([]data.User, error) {
	s.sql = s.sql.Where(sq.Eq{"role": data.RoleTypeDriver})
	return s.Select()
}

func (s *userStorage) SetManager(id uint64) error {
	_, err := s.newUpdate().Set("role", data.RoleTypeManager).Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to update user data")
	}
	return nil
}

func (s *userStorage) BlockUser(id uint64) error {
	_, err := s.newUpdate().Set("blocked", true).Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to block user")
	}
	return nil
}

func (s *userStorage) UnblockUser(id uint64) error {
	_, err := s.newUpdate().Set("blocked", false).Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to unblock user")
	}
	return nil
}
