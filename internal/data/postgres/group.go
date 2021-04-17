package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zlobste/spotter/internal/config"
	"github.com/zlobste/spotter/internal/data"
)

const (
	groupsTable = "groups"
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
}

var groupsSelect = sq.Select(all).From(groupsTable).PlaceholderFormat(sq.Dollar)

func NewGroupsStorage(cfg config.Config) GroupsStorage {
	return &groupStorage{
		db:  cfg.DB(),
		sql: groupsSelect.RunWith(cfg.DB()),
	}
}

func (s *groupStorage) New() GroupsStorage {
	return &groupStorage{
		db:  s.db,
		sql: groupsSelect.RunWith(s.db),
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
