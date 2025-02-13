package config

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var DB *DBConfig

func locaEnv() {
	env := os.Getenv("GO_ENV")
	var envFileName string

	// 環境に応じて .env ファイルを切り替える
	switch env {
	case "test":
		envFileName = ".env.test"
	default:
		envFileName = ".env"
	}
	envFilePath := filepath.Join(GetProjectRoot(), envFileName)

	// .env ファイルをロード
	if err := godotenv.Load(envFilePath); err != nil {
		log.Printf("Failed to load %s file: %v\n", envFilePath, err)
	} else {
		log.Printf("Loaded %s file\n", envFilePath)
	}
}


func GetProjectRoot() string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to get project root: %v", err)
	}
	return strings.TrimSpace(out.String())
}


func newDBConfig() *DBConfig {
	return &DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}
}

func InitConfig() {
	locaEnv()
	DB = newDBConfig()
}

