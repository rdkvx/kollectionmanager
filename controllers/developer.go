package controllers

import (
	"errors"
	"fmt"
	"kollectionmanager/m/models"
	"kollectionmanager/m/models/dto"
	"strings"
	"kollectionmanager/m/utils"

	"gorm.io/gorm"
)

func CreateDeveloper(developer models.Developer, db *gorm.DB) error {
	developer.Name = strings.ToLower(developer.Name)
	developer.Deleted = false
	result := db.Create(&developer)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetDevelopers(db *gorm.DB) ([]dto.Developer, error) {
	var devsRaw []models.Developer

	result := db.Find(&devsRaw)
	if result.Error != nil {
		return nil, result.Error
	}

	var devs []dto.Developer

	for _, dev := range devsRaw {
		if !dev.Deleted {
			var devtemp dto.Developer
			devtemp.Name = dev.Name
			devtemp.CreatedAt = dev.CreatedAt
			devtemp.UpdatedAt = dev.UpdatedAt

			devs = append(devs, devtemp)
		}
	}

	return devs, nil
}

func GetDeveloperByName(name string, db *gorm.DB) (dto.Developer, error) {
	var devRaw models.Developer

	result := db.Where(utils.FilterByName, strings.ToLower(name)).First(&devRaw)
	if result.Error != nil {
		return dto.Developer{}, result.Error
	}

	if devRaw.Deleted {
		return dto.Developer{}, errors.New(utils.FailedTo("find", "developer", name))
	}

	var dev dto.Developer
	dev.Name = devRaw.Name
	dev.CreatedAt = devRaw.CreatedAt
	dev.UpdatedAt = devRaw.UpdatedAt

	return dev, nil
}

func UpdateDeveloperByName(name string, developer dto.Developer, db *gorm.DB) error {

	name = strings.ToLower(name)
	developer.Name = strings.ToLower(developer.Name)

	result := db.Model(&models.Developer{}).Where(utils.FilterByName, name).Updates(models.Developer{Name: developer.Name})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		fmt.Println(utils.FailedTo("update", "developer", name))
		return errors.New(utils.FailedTo("update", "developer", name))
	}

	return nil
}

func SoftDeleteDeveloperByName(name string, db *gorm.DB) error {

	name = strings.ToLower(name)

	result := db.Model(&models.Developer{}).Where(utils.FilterByName, name).Where(utils.FilterByDeleted, false).Updates(models.Developer{Deleted: true})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		fmt.Println(utils.FailedTo("delete", "developer", name))
		return errors.New(utils.FailedTo("delete", "developer", name))
	}

	return nil
}
