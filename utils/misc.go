package utils

import (
	"errors"
	"kollectionmanager/m/models"
	"kollectionmanager/m/models/dto"
	"strings"
	"time"
)

// Pega uma model que veio do banco e parsea pra um DTO pra retornar no endpoint
func ConsoleModelToDTO(consoleRaw models.Console) dto.ConsoleGet {
	console := dto.ConsoleGet{
		ID:             consoleRaw.ID,
		Name:           consoleRaw.Name,
		ManufacturerID: consoleRaw.ManufacturerID,
		DtRelease:      consoleRaw.DtRelease,
		DtPurchase:     consoleRaw.DtPurchase,
		Owned:          consoleRaw.Owned,
	}

	for _, g := range consoleRaw.Games {
		game := dto.Game{
			ID:          g.ID,
			Name:        g.Name,
			ConsoleID:   g.ConsoleID,
			DeveloperID: g.DeveloperID,
			ReleaseDate: g.ReleaseDate,
			BoughtDate:  g.BoughtDate,
		}
		console.Games = append(console.Games, game)
	}

	return console
}

// Pega um DTO que veio pelo endpoint e parsea pra model pra ser inserido no banco
func ConsoleDTOToModel(consoleRaw dto.ConsolePost) (models.Console, error) {
	console := models.Console{}

	dtRelease, err := ParseTimeTo0GMT(consoleRaw.DtRelease)
	if err != nil {
		return console, err
	}

	dtPurchase, err := ParseTimeTo0GMT(consoleRaw.DtPurchase)
	if err != nil {
		return console, err
	}

	console.Name = strings.ToLower(consoleRaw.Name)
	console.Deleted = false
	console.ManufacturerID = consoleRaw.ManufacturerID
	console.DtRelease = dtRelease
	console.DtPurchase = dtPurchase
	console.Owned = consoleRaw.Owned

	return console, nil
}

// gambiarra para o banco setado -3gmt
func ParseTimeTo0GMT(dtRaw string) (time.Time, error) {
	dtRaw = dtRaw + " 04:00"
	dt, err := time.Parse("02/01/2006 15:04", dtRaw)
	if err != nil {
		dtErr := time.Time{}
		return dtErr, err
	}

	return dt, nil
}

func ValidateConsoleData(consoleRaw dto.ConsolePost, console models.Console) (models.Console, error) {
	emptyConsole := models.Console{}
	
	if console.Deleted {
		return emptyConsole, errors.New(FailedTo("find", "console", consoleRaw.Name))
	}

	if consoleRaw.Name != "" {
		console.Name = strings.ToLower(consoleRaw.Name)
	}

	if consoleRaw.ManufacturerID != 0 {
		console.ManufacturerID = consoleRaw.ManufacturerID
	}

	if consoleRaw.DtRelease != "" {
		dtRelease, err := ParseTimeTo0GMT(consoleRaw.DtRelease)
		if err != nil {
			return emptyConsole, err
		}
		console.DtPurchase = dtRelease
	}

	if consoleRaw.DtPurchase != "" {
		dtPurchase, err := ParseTimeTo0GMT(consoleRaw.DtPurchase)
		if err != nil {
			return emptyConsole, err
		}
		console.DtPurchase = dtPurchase
	}

	console.Owned = false
	if consoleRaw.Owned {
		console.Owned = true
	}

	return console, nil
}
