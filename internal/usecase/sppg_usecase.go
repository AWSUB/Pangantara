package usecase

import (
	"sppg-backend/internal/entity"
	"sppg-backend/internal/model"
	"sppg-backend/internal/repository"

	"github.com/google/uuid"
)

func CreateSPPG(req model.CreateSPPGRequest) (*entity.SPPG, error) {
	sppg := &entity.SPPG{
		SPPGID:      uuid.New(),
		UserID:      req.UserID,
		NameSPPG:    req.NameSPPG,
		LocationURL: req.LocationURL,
		Contact:     req.Contact,
	}
	return sppg, repository.CreateSPPG(sppg)
}

func GetAllSPPG() ([]entity.SPPG, error) {
	return repository.GetAllSPPG()
}

func GetSPPGByID(id uuid.UUID) (*entity.SPPG, error) {
	return repository.GetSPPGByID(id)
}

func GetSPPGByUserID(userID uuid.UUID) ([]entity.SPPG, error) {
	return repository.GetSPPGByUserID(userID)
}

func UpdateSPPG(id uuid.UUID, req model.UpdateSPPGRequest) error {
	data := map[string]interface{}{}
	if req.NameSPPG != "" {
		data["name_sppg"] = req.NameSPPG
	}
	if req.LocationURL != nil {
		data["location_url"] = req.LocationURL
	}
	if req.Contact != nil {
		data["contact"] = req.Contact
	}
	return repository.UpdateSPPG(id, data)
}

func DeleteSPPG(id uuid.UUID) error {
	return repository.DeleteSPPG(id)
}