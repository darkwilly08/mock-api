package dtos

import "darkwilly08.com/mockapi/src/entities"

type AppDto struct {
	Name string `json:"name"`
}

func (dto *AppDto) ToEntity() entities.AppEntity {
	return entities.AppEntity{
		Name: dto.Name,
	}
}

func (dto *AppDto) FromEntity(appEntity entities.AppEntity) {
	dto.Name = appEntity.Name
}
