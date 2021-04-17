package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zlobste/spotter/internal/config"
	"github.com/zlobste/spotter/internal/data"
)

const (
	votingsTable = "votings"
)

type votingStorage struct {
	db  *sql.DB
	sql sq.SelectBuilder
}

type votingsStorage interface {
	New() votingsStorage
	Get() (*data.Voting, error)
	GetVotingById(id uint64) (*data.Voting, error)
	CreateVoting(voting data.Voting) error
	UpdateVoting(id uint64, voting data.Voting) error
	DeleteVoting(id uint64) error
}

var votingsSelect = sq.Select(all).From(votingsTable).PlaceholderFormat(sq.Dollar)

func NewVotingsStorage(cfg config.Config) votingsStorage {
	return &votingStorage{
		db:  cfg.DB(),
		sql: votingsSelect.RunWith(cfg.DB()),
	}
}

func (s *votingStorage) New() votingsStorage {
	return &votingStorage{
		db:  s.db,
		sql: votingsSelect.RunWith(s.db),
	}
}

func (s *votingStorage) Get() (*data.Voting, error) {
	rowScanner := s.sql.QueryRow()
	model := data.Voting{}
	err := rowScanner.Scan(
		&model.Id,
		&model.Victim,
		&model.Type,
		&model.Title,
		&model.Description,
		&model.EndTime,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "failed to query model")
	} else if err == sql.ErrNoRows {
		return nil, nil
	}
	return &model, nil
}

func (s *votingStorage) GetVotingById(id uint64) (*data.Voting, error) {
	s.sql = s.sql.Where(sq.Eq{"id": id})
	return s.Get()
}

func (s *votingStorage) newInsert() sq.InsertBuilder {
	return sq.Insert(votingsTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *votingStorage) CreateVoting(voting data.Voting) error {
	_, err := s.newInsert().SetMap(voting.ToMap()).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to insert voting")
	}
	return nil
}

func (s *votingStorage) newUpdate() sq.UpdateBuilder {
	return sq.Update(votingsTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *votingStorage) UpdateVoting(id uint64, voting data.Voting) error {
	_, err := s.newUpdate().SetMap(voting.ToMap()).Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to update voting data")
	}
	return nil
}

func (s *votingStorage) newDelete() sq.DeleteBuilder {
	return sq.Delete(votingsTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *votingStorage) DeleteVoting(id uint64) error {
	_, err := s.newDelete().Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to delete voting")
	}
	return nil
}
