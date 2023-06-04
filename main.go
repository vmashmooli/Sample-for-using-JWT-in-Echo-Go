package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

var config Config

// Structure for configuration
type (
	Config struct {
		Setting struct {
			Server struct {
				Port int64 `json:"port"`
			}
			JWT struct {
				Secret string `json:"jwt"`
				Age    int64  `json:"age"`
			}
		}
	}
)

func main() {
	fmt.Println("Simple JWT with golang")

	// Getting settings from the configuration
	config = LoadConfig()
	fmt.Println("Server is running on port : ", config.Setting.Server.Port)
	fmt.Println("The secret key age is : ", config.Setting.JWT.Age)

	// Run the ECHO server on the desired port
	port := fmt.Sprintf(":%d", config.Setting.Server.Port)
	e := echo.New()
	Route(e)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(port))
}

// A function to load settings from a configuration file
func LoadConfig() Config {
	//Connecting to the config.json file
	configPath := "config.json"
	configFile, err := os.Open(configPath)
	if err != nil {
		log.Fatal("Error in loading config.json file : ", err)
	}
	defer configFile.Close()
	// Reading the setting data from JSON structure
	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		log.Fatal("Error in reading config.json file : ", err)
	}
	return config
}

// A function to routing ECHO
func Route(e *echo.Echo) {
	version := e.Group("api/v1")
	version.GET("/create_token", createToken)
	version.GET("/check_token", checkToken)
}

// host:port/api/v1/create_token
func createToken(c echo.Context) error {
	user_id := "0123456789"
	user_name := "user@company.com"
	if token, err := CreateToken(user_id, user_name); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	} else {
		return c.JSON(http.StatusOK, echo.Map{"token": token})
	}
}

// host:port/api/v1/check_token
func checkToken(c echo.Context) error {
	var token string
	token = c.Request().Header.Get("token")
	if token == "" {
		token = c.QueryParam("token")
		if token == "" {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "token is nil"})
		}
	}

	id, name, err := CheckToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "token is invalid"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "token is valid", "id": id, "name": name})
}
