package repositories

import (
	"darkwilly08.com/mockapi/src/config/database"
	"darkwilly08.com/mockapi/src/entities"
)

type AppRepository struct {
	db *database.DB
}

func NewAppRepository(db *database.DB) *AppRepository {
	return &AppRepository{
		db: db,
	}
}

func (r *AppRepository) Save(appEntity entities.AppEntity) {

}
