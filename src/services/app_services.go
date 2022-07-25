package services

import (
	"darkwilly08.com/mockapi/src/config/database"
	"darkwilly08.com/mockapi/src/dtos"
	"darkwilly08.com/mockapi/src/repositories"
)

type AppService struct {
	appRepository repositories.AppRepository
}

func NewAppService(db *database.DB) *AppService {
	return &AppService{
		appRepository: *repositories.NewAppRepository(db),
	}
}

func (s AppService) CreateOrUpdate(dto *dtos.AppDto) *dtos.AppDto {
	appEntity := dto.ToEntity()

	dto.FromEntity(appEntity)
	// return appRepository.Save(appEntity)
	return dto
}
