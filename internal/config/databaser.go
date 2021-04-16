package config

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zlobste/spotter/internal/data/migration"
	"sync"
)

type Databaser interface {
	DB() *sql.DB
}

type databaser struct {
	url    string
	method string

	cache struct {
		db *sql.DB
	}

	log *logrus.Logger
	sync.Once
}

func NewDatabaser(url, method string, log *logrus.Logger) Databaser {
	return &databaser{
		url:    url,
		method: method,
		log:    log,
	}
}

func (d *databaser) DB() *sql.DB {
	d.Once.Do(func() {
		var err error
		d.cache.db, err = sql.Open("postgres", d.url)
		if err != nil {
			panic(err)
		}

		switch d.method {
		case migration.Up:
			applied, err := migration.MigrateUp(d.cache.db)
			if err != nil {
				panic(err)
			}
			d.log.WithField("applied", applied).Info("Migrations up applied")
		case migration.Down:
			applied, err := migration.MigrateDown(d.cache.db)
			if err != nil {
				panic(err)
			}
			d.log.WithField("applied", applied).Info("Migrations down applied")
		default:
			panic("Unknown migration method")
		}

		if err := d.cache.db.Ping(); err != nil {
			panic(errors.Wrap(err, "database unavailable"))
		}
	})
	return d.cache.db
}
