package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"api/models"

	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var readDB *gorm.DB
var writeDB *gorm.DB

type DatabaseConfig struct {
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	DBName     string `yaml:"dbname"`
	ReaderHost string `yaml:"reader_host"`
	WriterHost string `yaml:"writer_host"`
	Params     string `yaml:"params"`
}

func dsnFromConfig(config DatabaseConfig, host string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", config.Username, config.Password, host, config.DBName, config.Params)
}

type Configurations struct {
	Database map[string]DatabaseConfig `yaml:"database"`
}

func init() {
	configFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal("Failed to open config file:", err)
	}
	defer configFile.Close()

	config, err := ioutil.ReadAll(configFile)

	var configurations Configurations
	err = yaml.Unmarshal(config, &configurations)
	if err != nil {
		log.Fatal("Failed to deserialize config:", err)
	}

	stage := os.Getenv("STAGE")
	dbConfig, exists := configurations.Database[stage]
	if !exists {
		log.Fatal("Specified stage does not exist in config:", stage)
	}

	readDSN := dsnFromConfig(dbConfig, dbConfig.ReaderHost)
	writeDSN := dsnFromConfig(dbConfig, dbConfig.WriterHost)

	readDB, err = gorm.Open(mysql.Open(readDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open config file:", err)
	}

	writeDB, err = gorm.Open(mysql.Open(writeDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open config file:", err)
	}

	// 両方のDBインスタンスでマイグレーションを行う
	readDB.AutoMigrate(&models.ToDo{})
	writeDB.AutoMigrate(&models.ToDo{})
}

// Get all todos
func GetAllToDos(c echo.Context) error {
	var todos []models.ToDo
	result := readDB.Find(&todos)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	if len(todos) == 0 {
		return c.String(http.StatusOK, "No record!!!")
	}

	return c.JSON(http.StatusOK, todos)
}

// GetToDoByID handles GET request to fetch a todo by ID
func GetToDoByID(c echo.Context) error {
	id := c.Param("id")
	todoID, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid todo ID")
	}

	var todo models.ToDo
	result := readDB.First(&todo, todoID)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Todo not found")
	}

	return c.JSON(http.StatusOK, todo)
}

// AddToDo handles POST request to add a new todo
func AddToDo(c echo.Context) error {
	var newToDo models.ToDo
	if err := c.Bind(&newToDo); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request data")
	}

	writeDB.Create(&newToDo)
	return c.JSON(http.StatusCreated, newToDo)
}
