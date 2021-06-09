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
