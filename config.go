package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
)

type Config struct {
	Home   string   `yaml:"home"`
	Server Server   `yaml:"server"`
	Log    AlertLog `yaml:"AlertLog"`
	Sqlite Sqlite   `yaml:"SQLite"`
}

type Server struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
	TLScert string `yaml:"tls_cert"`
	TLSkey  string `yaml:"tls_key"`
	Assets  string `yaml:"assets"`
}

type AlertLog struct {
	LogPath  string `yaml:"logPath"`
	LogLevel uint8  `yaml:"logLevel"`
}

type Sqlite struct {
	DbDriver string `yaml:"dbDriver"`
	DbPatch  string `yaml:"dbPatch"`
	DbFile   string `yaml:"dbFile"`
}

func initConfig(dirPath string) {
	Write(newConfig(dirPath))
	Read(dirPath)
}

func newConfig(dirPath string) *Config {
	var config Config

	config.Home = dirPath

	config.Server.Address = "0.0.0.0"
	config.Server.Port = "8443"
	config.Server.TLScert = path.Join(dirPath, "pki", "self-signed_cert.pem")
	config.Server.TLSkey = path.Join(dirPath, "pki", "self-signed_key.pem")
	config.Server.Assets = path.Join("$home", "assets")

	config.Log.LogPath = path.Join("$home", "log", alertLog) //назначаем переменной значение
	config.Log.LogLevel = 4

	config.Sqlite.DbDriver = "sqlite3"
	config.Sqlite.DbPatch = "$home"
	config.Sqlite.DbFile = "storage.db"

	return &config
}

func Write(config *Config) {
	data, err := yaml.Marshal(&config)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path.Join(config.Home, configFile), data, 0660)
	if err != nil {
		panic(err)
	}
}

func Read(dirPath string) Config {
	var config Config
	confPath := path.Join(dirPath, configFile)

	configData, err := os.ReadFile(confPath)
	if err != nil {
		log.Fatal(err)
	}

	// parse the YAML stored in the byte slice into the struct
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal config data: %w", err))
	}
	return config
}
