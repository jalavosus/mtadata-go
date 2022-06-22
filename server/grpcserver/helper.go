package grpcserver

import (
	"gorm.io/gorm"

	"github.com/jalavosus/mtadata/models"
	"github.com/jalavosus/mtadata/models/apimethods"
)

func dbError(err error, entityName, entityId string, method apimethods.ApiMethod) models.ApiError {
	if err == gorm.ErrRecordNotFound {
		return models.EntityNotFoundError(entityName, entityId, method)
	}

	return models.InternalError(err, method)
}
