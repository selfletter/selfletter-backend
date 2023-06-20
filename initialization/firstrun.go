package initialization

import (
	"fmt"
	"os"
	"selfletter-backend/config"
	"selfletter-backend/db"
	"selfletter-backend/secureToken"
)

func FirstRun() {
	file, err := os.Create("admin_keys.txt")
	dbHandle := db.GetDatabaseHandle()
	if err != nil {
		panic("first run: failed to create file containing frontend admin keys")
	}
	key := secureToken.GenerateSecureToken()

	_, err = file.Write([]byte(key))
	if err != nil {
		panic("first run: failed to write file containing frontend admin keys")
	}

	err = file.Sync()
	if err != nil {
		panic("first run: failed to write file containing frontend admin keys")
	}

	if err := dbHandle.Create(&db.AdminKey{Key: key}).Error; err != nil {
		panic("first run: failed to add admin key to database")
	}

	fmt.Println("Your frontend admin key will be in file admin_keys.txt\nYou can also access it from database table admin_keys")

	cfg := config.GetConfig()
	cfg.FirstRun = false
	err = config.WriteConfig(cfg)
	if err != nil {
		panic("first run: failed to write config file")
	}
}
