package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zlobste/spotter/internal/data"
)

const (
	confirmationsTable = "confirmations"
)

type confirmationStorage struct {
	db  *sql.DB
	sql sq.SelectBuilder
}

type confirmationsStorage interface {
	New() confirmationsStorage
	Get() (*data.Confirmation, error)
	GetConfirmationById(id uint64) (*data.Confirmation, error)
	CreateConfirmation(confirmation data.Confirmation) error
	UpdateConfirmation(id uint64, confirmation data.Confirmation) error
	DeleteConfirmation(id uint64) error
}

var confirmationsSelect = sq.Select(all).From(confirmationsTable).PlaceholderFormat(sq.Dollar)

func (s *confirmationStorage) New() confirmationsStorage {
	return NewConfirmationsStorage(s.db)
}

func NewConfirmationsStorage(db *sql.DB) confirmationsStorage {
	return &confirmationStorage{
		db:  db,
		sql: confirmationsSelect.RunWith(db),
	}
}

func (s *confirmationStorage) Get() (*data.Confirmation, error) {
	rowScanner := s.sql.QueryRow()
	model := data.Confirmation{}
	err := rowScanner.Scan(
		&model.UserId,
		&model.TimerId,
		&model.Date,
		&model.Confirmed,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "failed to query model")
	} else if err == sql.ErrNoRows {
		return nil, nil
	}
	return &model, nil
}

func (s *confirmationStorage) GetConfirmationById(id uint64) (*data.Confirmation, error) {
	s.sql = s.sql.Where(sq.Eq{"id": id})
	return s.Get()
}

func (s *confirmationStorage) newInsert() sq.InsertBuilder {
	return sq.Insert(confirmationsTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *confirmationStorage) CreateConfirmation(confirmation data.Confirmation) error {
	_, err := s.newInsert().SetMap(confirmation.ToMap()).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to insert confirmation")
	}
	return nil
}

func (s *confirmationStorage) newUpdate() sq.UpdateBuilder {
	return sq.Update(confirmationsTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *confirmationStorage) UpdateConfirmation(id uint64, confirmation data.Confirmation) error {
	_, err := s.newUpdate().SetMap(confirmation.ToMap()).Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to update confirmation data")
	}
	return nil
}

func (s *confirmationStorage) newDelete() sq.DeleteBuilder {
	return sq.Delete(confirmationsTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *confirmationStorage) DeleteConfirmation(id uint64) error {
	_, err := s.newDelete().Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to delete confirmation")
	}
	return nil
}
