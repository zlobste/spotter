package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zlobste/spotter/internal/config"
	"github.com/zlobste/spotter/internal/data"
)

const (
	usersTable = "users"
	all        = "*"
)

type userStorage struct {
	db  *sql.DB
	sql sq.SelectBuilder
}

type UsersStorage interface {
	New() UsersStorage
	Get() (*data.User, error)
	GetUser(username string) (*data.User, error)
	GetUserById(id int64) (*data.User, error)
	CreateUser(user data.User) error
	UpdateUser(oldUsername string, user data.User) error
	DeleteUser(username string) error
}

var usersSelect = sq.Select(all).From(usersTable).PlaceholderFormat(sq.Dollar)

func NewUsersStorage(cfg config.Config) UsersStorage {
	return &userStorage{
		db:  cfg.DB(),
		sql: usersSelect.RunWith(cfg.DB()),
	}
}

func (s *userStorage) New() UsersStorage {
	return &userStorage{
		db:  s.db,
		sql: usersSelect.RunWith(s.db),
	}
}

func (s *userStorage) Get() (*data.User, error) {
	rowScanner := s.sql.QueryRow()
	user := data.User{}
	err := rowScanner.Scan(
		&user.Id,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.Password,
		&user.Balance,
		&user.Role,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "failed to query user")
	} else if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, nil
}

func (s *userStorage) GetUser(email string) (*data.User, error) {
	s.sql = s.sql.Where(sq.Eq{"email": email})
	return s.Get()
}

func (s *userStorage) GetUserById(id int64) (*data.User, error) {
	s.sql = s.sql.Where(sq.Eq{"id": id})
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

func (s *userStorage) UpdateUser(email string, user data.User) error {
	_, err := s.newUpdate().SetMap(user.ToMap()).Where(sq.Eq{"email": email}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to update user data")
	}
	return nil
}

func (s *userStorage) newDelete() sq.DeleteBuilder {
	return sq.Delete(usersTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *userStorage) DeleteUser(email string) error {
	_, err := s.newDelete().Where(sq.Eq{"email": email}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}
	return nil
}
