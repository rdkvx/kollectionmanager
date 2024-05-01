package utils

import (
	"errors"
	"kollectionmanager/m/models"
	"kollectionmanager/m/models/dto"
	"strings"
	"time"
)

// gambiarra para o banco setar o banco 0gmt
/*
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣤⣤⣤⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⠟⠉⠀⠀⠈⠙⢿⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⠃⠀⠀⠀⠀⠀⠀⠀⢻⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⡇⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⣇⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⡀⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣶⠾⠛⠛⠉⠉⠉⠉⠉⠉⠉⠉⠉⠛⠛⠻⢷⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⡟⠁⠀⠀⣀⣤⣤⣤⣶⣶⣶⣤⣤⣤⣀⡀⠀⠀⠙⣿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⣠⣶⠿⠿⢷⣦⣤⣶⣶⡄⣀⣀⡀⠀⠀⠀⠀⢹⣧⣶⠾⠛⠉⠉⠁⠀⠀⠀⠀⠀⠀⠉⠉⠛⠿⣶⣤⣿⠇⠀⠀⠀⠀⠀⢀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⣿⢁⠀⠀⠀⠈⠁⠀⠘⠿⠛⠛⠻⣷⡄⣠⡾⠟⠉⠀⢀⣠⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⣀⠀⠈⠛⢿⣦⡀⠀⠀⢠⣿⠛⠛⣷⣶⠿⠛⢿⣶⡀⠀⠀⠀
⠀⠀⣀⣿⣟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⠋⠀⠀⢀⡴⠋⠀⣷⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠃⠈⠳⣄⠀⠀⠙⢿⣦⣠⣼⡧⠀⠀⠙⠃⠀⠀⠀⢸⡇⠀⠀⠀
⢠⣾⠟⠋⠛⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⡿⠃⠀⠀⢀⣞⠀⠀⠀⣾⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡄⠀⠀⠘⣧⠀⠀⢠⣿⠋⠉⠀⠀⠀⠀⠀⠀⠀⢠⣿⣧⣤⡀⠀
⣿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢻⡆⠀⠀⠀⠀⠉⢉⣭⣭⣅⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣉⣭⣟⡉⠁⠀⠀⢸⣧⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠁⠀⠘⣿⡄
⢿⣇⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣤⣾⠃⠀⠀⠀⢀⣾⣿⣿⡏⠉⣳⡄⠀⠀⠀⠀⠀⠀⣰⢿⣿⡟⠉⠹⣆⠀⠀⠀⢹⡿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣹⡇
⠀⢹⡿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⣷⡀⠀⠀⠀⠀⠸⡇⠿⣿⣿⡾⢋⡟⠀⠀⠀⠀⠀⠀⣇⠸⣿⣿⣶⠏⣿⠀⠀⠀⠸⣷⣄⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⠃
⠀⢿⡇⠀⢀⠀⠀⠀⠀⠀⠀⠀⢀⢀⣿⠇⣴⠚⠉⠙⠳⡿⢦⣄⣀⣴⠞⣁⣤⡶⠶⢶⣤⣄⡘⠷⣤⣀⣤⠴⡿⠛⠉⠉⠢⣨⡿⠛⠁⠀⠀⠀⠀⠀⠀⢰⣾⠟⠁⠀
⠀⠈⠛⢿⡿⠀⠀⠀⠀⠀⠀⠀⠚⣿⡏⠰⡇⠀⠀⠀⠀⣱⠀⠈⠁⣠⣾⠛⠁⠀⠀⠀⠈⠙⢿⣦⠀⠀⠀⢸⡁⠀⠀⠀⠀⣿⣇⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⡆⠀⠀
⠀⠀⠀⢸⣧⡀⠀⠀⠀⠀⠀⠀⢀⣿⠇⠀⠓⢦⣀⣀⠴⠃⠀⠀⢠⡿⠁⠀⠀⠀⠀⠀⠀⠀⠈⢻⣦⠀⠀⠀⠳⢤⣄⣠⡜⠁⠙⣿⡆⠀⠀⢀⣄⣀⣀⣤⡿⠃⠀⠀
⠀⠀⠀⠀⠙⠻⠿⢷⣆⣀⣀⣴⣾⡋⠀⠀⠀⠀⣴⠛⠙⠓⠦⣄⣸⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣠⠶⠛⠉⠳⢦⠀⠀⠀⠀⠹⣧⣀⣠⣾⠟⠉⠉⠉⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⠋⠁⢻⣇⠀⠀⠀⠀⡇⠀⠠⡄⠀⠈⠙⣿⡄⠀⠀⠀⠀⠀⠀⠀⢀⣾⠟⠁⠀⢠⡆⠀⢈⡇⠀⠀⠀⢠⡿⠛⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⡄⠀⠀⠀⣿⠀⠀⠻⣄⠀⠀⠈⠻⣦⣄⡀⠀⢀⣀⣴⡿⠋⠀⠀⢠⠟⠀⠀⣸⠃⠀⠀⢀⣾⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⡄⠀⠀⠘⣇⠀⠀⠘⢧⡀⠀⠀⠈⠙⠛⠛⠛⠋⠉⠀⠀⢀⡴⠋⠀⠀⣰⠏⠀⠀⢀⣾⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢿⣆⠀⠀⠘⣦⡀⠀⠀⠉⠳⣄⣀⠀⠀⠀⠀⠀⠀⣀⡴⠋⠀⠀⢀⡴⠃⠀⠀⢠⡿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⣷⣄⠀⠀⠻⢦⣀⠀⠀⠀⠉⠙⠓⠒⠒⠛⠉⠀⠀⠀⢀⡴⠟⠁⠀⢀⣾⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠛⢷⣦⣀⠀⠉⠛⠶⠤⣄⣀⣀⣀⣀⣀⣀⡤⠴⠚⠋⠀⢀⣤⡾⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⠷⣦⣤⣀⣀⠀⠈⠉⠉⠉⠉⠀⣀⣀⣤⣴⡾⠛⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⠛⠛⠛⠛⠒⠚⠛⠛⠛⠋⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
*/
func ParseTimeTo0GMT(dtRaw string) (time.Time, error) {
	if dtRaw == "" {
		dtRaw = "31/12/2099"
	}
	dtRaw = dtRaw + " 04:00"	
	dt, err := time.Parse("02/01/2006 15:04", dtRaw)
	if err != nil {
		dtErr := time.Time{}
		return dtErr, err
	}

	return dt, nil
}


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
		game := dto.GameGet{
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

func GameModelToDTO(gameRaw models.Game) (dto.GameGet) {
	game := dto.GameGet{
        ID:          gameRaw.ID,
        Name:        gameRaw.Name,
        ConsoleID:   gameRaw.ConsoleID,
        DeveloperID: gameRaw.DeveloperID,
        ReleaseDate: gameRaw.ReleaseDate,
        BoughtDate:  gameRaw.BoughtDate,
    }

    return game
}

func GameDTOToModel(gameRaw dto.GamePost) (models.Game, error) {
	game := models.Game{}

    dtRelease, err := ParseTimeTo0GMT(gameRaw.ReleaseDate)
    if err != nil {
        return game, err
    }

    dtPurchase, err := ParseTimeTo0GMT(gameRaw.BoughtDate)
    if err != nil {
        return game, err
    }

    game.Name = strings.ToLower(gameRaw.Name)
    game.Deleted = false
    game.ConsoleID = gameRaw.ConsoleID
    game.DeveloperID = gameRaw.DeveloperID
    game.ReleaseDate = dtRelease
    game.BoughtDate = dtPurchase

    return game, nil
}

func DeveloperDtoToModel(developerRaw dto.DeveloperPost) (models.Developer){
	developer := models.Developer{
		Name: strings.ToLower(developerRaw.Name),
        Deleted: false,
	}

	return developer
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

	if consoleRaw.Owned && !console.Owned{
		console.Owned = true
	}

	if !consoleRaw.Owned && console.Owned{
		console.Owned = false
	}

	return console, nil
}

func ValidateGameData(gameRaw dto.GamePost, game models.Game) (models.Game, error) {
	emptyGame := models.Game{}

	if game.Deleted {
		return emptyGame, errors.New(FailedTo("find", "game", gameRaw.Name))
	}

	if gameRaw.Name != "" {
		game.Name = strings.ToLower(gameRaw.Name)
	}

	if gameRaw.ConsoleID != 0 {
		game.ConsoleID = gameRaw.ConsoleID
	}

	if gameRaw.DeveloperID != 0 {
		game.DeveloperID = gameRaw.DeveloperID
	}

	if gameRaw.BoughtDate != "" {
		boughtDate, err := ParseTimeTo0GMT(gameRaw.BoughtDate)
		if err != nil {
			return emptyGame, err
		}
		game.BoughtDate = boughtDate
	}

	if gameRaw.ReleaseDate != "" {
		releaseDate, err := ParseTimeTo0GMT(gameRaw.ReleaseDate)
		if err != nil {
			return emptyGame, err
		}
		game.ReleaseDate = releaseDate
	}

	if gameRaw.Owned && !game.Owned {
		game.Owned = true
	}

    if !gameRaw.Owned && game.Owned {
        game.Owned = false
    }

	return game, nil
}