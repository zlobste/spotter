package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zlobste/spotter/internal/data"
	"time"
)

const (
	timersTable = "timers"
)

type TimersStorage interface {
	New() TimersStorage
	Get() (*data.Timer, error)
	GetTimerById(id uint64) (*data.Timer, error)
	GetTimersByDriver(id uint64) ([]data.Timer, error)
	CreateTimer(timer data.Timer) error
	UpdateTimer(id uint64, timer data.Timer) error
	DeleteTimer(id uint64) error
	StartTimer(userId uint64) error
	StopTimer(userId uint64) error
	GetPendingTimer(userId uint64) (*data.Timer, error)
}

type timerStorage struct {
	db  *sql.DB
	sql sq.SelectBuilder
}

var timersSelect = sq.Select(all).From(timersTable).PlaceholderFormat(sq.Dollar)

func (s *timerStorage) New() TimersStorage {
	return NewTimersStorage(s.db)
}

func NewTimersStorage(db *sql.DB) TimersStorage {
	return &timerStorage{
		db:  db,
		sql: timersSelect.RunWith(db),
	}
}

func (s *timerStorage) Get() (*data.Timer, error) {
	rowScanner := s.sql.QueryRow()
	model := data.Timer{}
	err := rowScanner.Scan(
		&model.Id,
		&model.UserId,
		&model.StartTime,
		&model.EndTime,
		&model.Pending,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "failed to query model")
	} else if err == sql.ErrNoRows {
		return nil, nil
	}
	return &model, nil
}

func (s *timerStorage) Select() ([]data.Timer, error) {
	rows, err := s.sql.RunWith(s.db).Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var models []data.Timer

	for rows.Next() {
		model := data.Timer{}
		err := rows.Scan(
			&model.Id,
			&model.UserId,
			&model.StartTime,
			&model.EndTime,
			&model.Pending,
		)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}

func (s *timerStorage) GetTimerById(id uint64) (*data.Timer, error) {
	s.sql = s.sql.Where(sq.Eq{"id": id})
	return s.Get()
}

func (s *timerStorage) newInsert() sq.InsertBuilder {
	return sq.Insert(timersTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *timerStorage) CreateTimer(timer data.Timer) error {
	_, err := s.newInsert().SetMap(timer.ToMap()).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to insert timer")
	}
	return nil
}

func (s *timerStorage) newUpdate() sq.UpdateBuilder {
	return sq.Update(timersTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *timerStorage) UpdateTimer(id uint64, timer data.Timer) error {
	_, err := s.newUpdate().SetMap(timer.ToMap()).Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to update timer data")
	}
	return nil
}

func (s *timerStorage) newDelete() sq.DeleteBuilder {
	return sq.Delete(timersTable).RunWith(s.db).PlaceholderFormat(sq.Dollar)
}

func (s *timerStorage) DeleteTimer(id uint64) error {
	_, err := s.newDelete().Where(sq.Eq{"id": id}).Exec()
	if err != nil {
		return errors.Wrap(err, "failed to delete timer")
	}
	return nil
}

func (s *timerStorage) GetTimersByDriver(id uint64) ([]data.Timer, error) {
	s.sql = s.sql.Where(sq.Eq{"user_id": id})
	return s.Select()
}

func (s *timerStorage) StartTimer(userId uint64) error {
	pendingTimer, err := s.GetPendingTimer(userId)
	if err != nil {
		return err
	}
	if pendingTimer != nil {
		return errors.New("timer is pending")
	}
	newTimer := data.Timer{
		UserId:    userId,
		StartTime: time.Now(),
		Pending:   true,
	}
	_, err = s.newInsert().SetMap(newTimer.ToMap()).Exec()
	return errors.Wrap(err, "failed to insert timer")
}

func (s *timerStorage) StopTimer(userId uint64) error {
	pendingTimer, err := s.GetPendingTimer(userId)
	if err != nil {
		return err
	}
	if pendingTimer == nil {
		return errors.New("there is no pending timer")
	}
	pendingTimer.EndTime = time.Now()
	pendingTimer.Pending = false
	_, err = s.newUpdate().SetMap(pendingTimer.ToMap()).Where(sq.Eq{"user_id": userId}).Exec()
	return err
}

func (s *timerStorage) GetPendingTimer(userId uint64) (*data.Timer, error) {
	s.sql = s.sql.Where(sq.Eq{"pending": true, "user_id": userId})
	return s.Get()
}
