package databases

import (
	"healthcare-capt-america/pkg/configs"
	"healthcare-capt-america/utils"

	_ "github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewRepositories(config *configs.Config) (*Repositories, error) {
	conn := utils.GetConnection(config)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}

	return &Repositories{
		db: db,
	}, nil
}
