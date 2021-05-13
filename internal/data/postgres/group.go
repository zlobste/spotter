package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zlobste/spotter/internal/data"
)

const (
	groupsTable      = "groups"
	usersGroupsTable = "users_groups"
)

type groupStorage struct {
	db  *sql.DB
	sql sq.SelectBuilder
}

type GroupsStorage interface {
	New() GroupsStorage
	Get() (*data.Group, error)
	GetGroupById(id uint64) (*data.Group, error)
	CreateGroup(group data.Group) error
	UpdateGroup(id uint64, group data.Group) error
	DeleteGroup(id uint64) error
	GetGroupUsers(id uint64) ([]data.User, error)
	AddUserToGroup(groupUser data.UserGroup) error
}

var groupsSelect = sq.Select(all).From(groupsTable).PlaceholderFormat(sq.Dollar)

func (s *groupStorage) New() GroupsStorage {
	return NewGroupsStorage(s.db)
}

func NewGroupsStorage(db *sql.DB) GroupsStorage {
	return &groupStorage{
		db:  db,
		sql: groupsSelect.RunWith(db),
	}
}

func (s *groupStorage) Get() (*data.Group, error) {
	rowScanner := s.sql.QueryRow()
	model := data.Group{}
	err := rowScanner.Scan(
		&model.Id,
		&model.Title,
		&model.Description,
		&model.Level,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "failed to query model")
	} else if err == sql.ErrNoRows {
		return nil, nil
	}
	return &model, nil
}

func (s *groupStorage) Select() ([]data.Group, error) {
	rows, err := s.sql.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var models []data.Group

	for rows.Next() {
		model := data.Group{}
		err := rows.Scan(&model.Id, &model.Title, &model.Description, &model.Level)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}

func (s *groupStorage) GetGroupById(id uint64) (*data.Group, error) {
	s.sql = s.sql.Where(sq.Eq{"id": id})
	return s.Get()
}

func (s *groupStorage) newInsert() sq.InsertBuilder {
	return sq.Insert(groupsTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *groupStorage) CreateGroup(group data.Group) error {
	_, err := s.newInsert().SetMap(group.ToMap()).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to insert group")
	}
	return nil
}

func (s *groupStorage) newUpdate() sq.UpdateBuilder {
	return sq.Update(groupsTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *groupStorage) UpdateGroup(id uint64, group data.Group) error {
	_, err := s.newUpdate().SetMap(group.ToMap()).Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to update group data")
	}
	return nil
}

func (s *groupStorage) newDelete() sq.DeleteBuilder {
	return sq.Delete(groupsTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *groupStorage) DeleteGroup(id uint64) error {
	_, err := s.newDelete().Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to delete group")
	}
	return nil
}

func (s *groupStorage) GetGroupUsers(id uint64) ([]data.User, error) {
	rows, err := s.db.Query("select * from users where id in (select user_id from users_groups where group_id = $1)", id)
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
			&model.Balance,
			&model.Salary,
		)
		if err != nil {
			return nil, err
		}
		model.Password = ""
		models = append(models, model)
	}

	return models, nil
}

func (s *groupStorage) AddUserToGroup(groupUser data.UserGroup) error {
	_, err := sq.Insert(usersGroupsTable).RunWith(s.db).PlaceholderFormat(sq.Dollar).SetMap(groupUser.ToMap()).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to add user to group")
	}
	return nil
}
