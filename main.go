package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Character struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Team       string  `json:"team"`
	Position   string  `json:"position"`
	BattingAvg float64 `json:"battingAvg"`
}

var (
	characters = make(map[int]*Character)
	nextID     = 1
)

func getAllCharacters(c echo.Context) error {
	charList := []*Character{}
	for _, char := range characters {
		charList = append(charList, char)
	}
	return c.JSON(http.StatusOK, charList)
}

func getCharacter(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "ID inválido")
	}

	char, ok := characters[id]
	if !ok {
		return c.String(http.StatusNotFound, "Personaje no encontrado")
	}
	return c.JSON(http.StatusOK, char)
}

func createCharacter(c echo.Context) error {
	char := new(Character)
	if err := c.Bind(char); err != nil {
		return c.String(http.StatusBadRequest, "Datos de personaje inválidos")
	}

	char.ID = nextID
	characters[char.ID] = char
	nextID++

	return c.JSON(http.StatusCreated, char)
}

func updateCharacter(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "ID inválido")
	}

	char, ok := characters[id]
	if !ok {
		return c.String(http.StatusNotFound, "Personaje no encontrado")
	}

	updatedChar := new(Character)
	if err := c.Bind(updatedChar); err != nil {
		return c.String(http.StatusBadRequest, "Datos de personaje inválidos")
	}

	if updatedChar.Name != "" {
		char.Name = updatedChar.Name
	}
	if updatedChar.Team != "" {
		char.Team = updatedChar.Team
	}
	if updatedChar.Position != "" {
		char.Position = updatedChar.Position
	}
	if updatedChar.BattingAvg != 0 {
		char.BattingAvg = updatedChar.BattingAvg
	}

	return c.JSON(http.StatusOK, char)
}

func deleteCharacter(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "ID inválido")
	}

	_, ok := characters[id]
	if !ok {
		return c.String(http.StatusNotFound, "Personaje no encontrado")
	}

	delete(characters, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Rutas de la API
	e.GET("/characters", getAllCharacters)
	e.GET("/characters/:id", getCharacter)
	e.POST("/characters", createCharacter)
	e.PUT("/characters/:id", updateCharacter)
	e.DELETE("/characters/:id", deleteCharacter)

	e.Logger.Fatal(e.Start(":8080"))
}
