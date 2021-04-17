package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zlobste/spotter/internal/config"
	"github.com/zlobste/spotter/internal/data"
)

const (
	votesTable = "votes"
)

type voteStorage struct {
	db  *sql.DB
	sql sq.SelectBuilder
}

type votesStorage interface {
	New() votesStorage
	Get() (*data.Vote, error)
	GetVoteById(id uint64) (*data.Vote, error)
	CreateVote(vote data.Vote) error
	UpdateVote(id uint64, vote data.Vote) error
	DeleteVote(id uint64) error
}

var votesSelect = sq.Select(all).From(votesTable).PlaceholderFormat(sq.Dollar)

func NewVotesStorage(cfg config.Config) votesStorage {
	return &voteStorage{
		db:  cfg.DB(),
		sql: votesSelect.RunWith(cfg.DB()),
	}
}

func (s *voteStorage) New() votesStorage {
	return &voteStorage{
		db:  s.db,
		sql: votesSelect.RunWith(s.db),
	}
}

func (s *voteStorage) Get() (*data.Vote, error) {
	rowScanner := s.sql.QueryRow()
	model := data.Vote{}
	err := rowScanner.Scan(
		&model.VotingId,
		&model.UserId,
		&model.Decided,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "failed to query model")
	} else if err == sql.ErrNoRows {
		return nil, nil
	}
	return &model, nil
}

func (s *voteStorage) GetVoteById(id uint64) (*data.Vote, error) {
	s.sql = s.sql.Where(sq.Eq{"id": id})
	return s.Get()
}

func (s *voteStorage) newInsert() sq.InsertBuilder {
	return sq.Insert(votesTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *voteStorage) CreateVote(vote data.Vote) error {
	_, err := s.newInsert().SetMap(vote.ToMap()).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to insert vote")
	}
	return nil
}

func (s *voteStorage) newUpdate() sq.UpdateBuilder {
	return sq.Update(votesTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *voteStorage) UpdateVote(id uint64, vote data.Vote) error {
	_, err := s.newUpdate().SetMap(vote.ToMap()).Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to update vote data")
	}
	return nil
}

func (s *voteStorage) newDelete() sq.DeleteBuilder {
	return sq.Delete(votesTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *voteStorage) DeleteVote(id uint64) error {
	_, err := s.newDelete().Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to delete vote")
	}
	return nil
}
