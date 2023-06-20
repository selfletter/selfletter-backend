package initialization

import (
	"selfletter-backend/config"
	"selfletter-backend/db"
)

func InitializeConfigAndDatabase() error {
	err := config.ParseConfig()
	if err != nil {
		return err
	}
	cfg := config.GetConfig()

	err = db.Open()
	if err != nil {
		return err
	}

	if cfg.FirstRun {
		FirstRun()
	}

	return nil
}
