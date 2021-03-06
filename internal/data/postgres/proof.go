package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zlobste/spotter/internal/data"
	"time"
)

const (
	proofsTable            = "proofs"
	proofThreshold float64 = 0.3
)

type ProofsStorage interface {
	New() ProofsStorage
	Get() (*data.Proof, error)
	GetProofById(id uint64) (*data.Proof, error)
	CreateProof(proof data.Proof) error
	UpdateProof(id uint64, proof data.Proof) error
	DeleteProof(id uint64) error
	GetProofsByTimer(timerId uint64) ([]data.Proof, error)
	MakeProof(timerId uint64, percentage float64) error
}

type proofStorage struct {
	db  *sql.DB
	sql sq.SelectBuilder
}

var proofsSelect = sq.Select(all).From(proofsTable).PlaceholderFormat(sq.Dollar)

func (s *proofStorage) New() ProofsStorage {
	return NewProofsStorage(s.db)
}

func NewProofsStorage(db *sql.DB) ProofsStorage {
	return &proofStorage{
		db:  db,
		sql: proofsSelect.RunWith(db),
	}
}

func (s *proofStorage) Get() (*data.Proof, error) {
	rowScanner := s.sql.QueryRow()
	model := data.Proof{}
	err := rowScanner.Scan(
		&model.Id,
		&model.TimerId,
		&model.Time,
		&model.Percentage,
		&model.Confirmed,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "failed to query model")
	} else if err == sql.ErrNoRows {
		return nil, nil
	}
	return &model, nil
}

func (s *proofStorage) Select() ([]data.Proof, error) {
	rows, err := s.sql.RunWith(s.db).Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var models []data.Proof

	for rows.Next() {
		model := data.Proof{}
		err := rows.Scan(
			&model.Id,
			&model.TimerId,
			&model.Time,
			&model.Percentage,
			&model.Confirmed,
		)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}

func (s *proofStorage) GetProofById(id uint64) (*data.Proof, error) {
	s.sql = s.sql.Where(sq.Eq{"id": id})
	return s.Get()
}

func (s *proofStorage) newInsert() sq.InsertBuilder {
	return sq.Insert(proofsTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *proofStorage) CreateProof(proof data.Proof) error {
	_, err := s.newInsert().SetMap(proof.ToMap()).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to insert proof")
	}
	return nil
}

func (s *proofStorage) newUpdate() sq.UpdateBuilder {
	return sq.Update(proofsTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *proofStorage) UpdateProof(id uint64, proof data.Proof) error {
	_, err := s.newUpdate().SetMap(proof.ToMap()).Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to update proof data")
	}
	return nil
}

func (s *proofStorage) newDelete() sq.DeleteBuilder {
	return sq.Delete(proofsTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *proofStorage) DeleteProof(id uint64) error {
	_, err := s.newDelete().Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to delete proof")
	}
	return nil
}

func (s *proofStorage) GetProofsByTimer(id uint64) ([]data.Proof, error) {
	s.sql = s.sql.Where(sq.Eq{"timer_id": id})
	return s.Select()
}

func (s *proofStorage) MakeProof(timerId uint64, percentage float64) error {
	proof := data.Proof{
		TimerId:    timerId,
		Time:       time.Now(),
		Percentage: percentage,
		Confirmed:  percentage < proofThreshold,
	}
	_, err := s.newInsert().SetMap(proof.ToMap()).Exec()
	return err
}
