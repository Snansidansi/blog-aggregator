package state

import (
	"github.com/snansidansi/blog-aggregator/internal/config"
	"github.com/snansidansi/blog-aggregator/internal/database"
)

type State struct {
	Config *config.Config
	Db     *database.Queries
}
