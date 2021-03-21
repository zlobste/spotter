package spotter

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/zlobste/spotter/internal/config"
)

// Service ...
type Service struct {
	config config.Config
	log    *logrus.Logger
}

// New creates a service ...
func New(cfg config.Config) *Service {
	return &Service{
		config: cfg,
		log:    cfg.Logger(),
	}
}

// Run performs events listening and querying the Odin  minting module.
func (s *Service) Run(ctx context.Context) error {
	return nil
}
