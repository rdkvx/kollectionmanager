package controllers

import (
	"kollectionmanager/m/models"
	"kollectionmanager/m/models/dto"
	"strings"
	"time"

	"gorm.io/gorm"
)

func GetManufacturers(db *gorm.DB) ([]dto.ManufacturerGet, error) {
	var manufacturerRaw []models.Manufacturer
	var manufacturers []dto.ManufacturerGet
	
	result := db.Find(&manufacturerRaw)
	if result.Error != nil {
		return manufacturers, result.Error
	}

	for _, manufacturer := range manufacturerRaw {
		if !manufacturer.Deleted {
			var manutemp dto.ManufacturerGet
			manutemp.Name = manufacturer.Name
			manutemp.Founded = manufacturer.Founded

			manufacturers = append(manufacturers, manutemp)
		}
	}

	return manufacturers, nil
}

func CreateManufacturer(manufacturerRaw dto.ManufacturerPost, db *gorm.DB) error{
	var manufacturer models.Manufacturer
	manufacturer.Name = strings.ToLower(manufacturerRaw.Name)
	manufacturer.Deleted = false

	//gambiarra mais feia do que bater em mãe
	manufacturerRaw.Founded = manufacturerRaw.Founded + " 05:00"
	
	dtFounded, err := time.Parse("02/01/2006 15:04", manufacturerRaw.Founded)
	if err != nil {
		return err
	}

	manufacturer.Founded = dtFounded
	result := db.Create(&manufacturer)
	if result.Error != nil {
        return result.Error
    }

	return nil
}