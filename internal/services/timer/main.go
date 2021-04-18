package timer

import (
	"github.com/sirupsen/logrus"
	"github.com/zlobste/spotter/internal/config"
)

type Service struct {
	log  *logrus.Logger
}

func New(cfg config.Config) *Service {
	return &Service{
		log:  cfg.Logging(),
	}
}

func (s *Service) Run() error {



	return nil
}

func (s *Service) CreateVoting() error {



	return nil
}
