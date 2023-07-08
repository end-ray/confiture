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
}

type AlertLog struct {
	LogPath  string `yaml:"logPath"`
	LogLevel uint8  `yaml:"logLevel"`
}

type Sqlite struct {
	DbDriver string `yaml:"dbDriver"`
	DbFile   string `yaml:"dbFile"`
}

func initConfig(dirPath string) {
	if _, err := os.Stat(path.Join(dirPath, "conf", configFile)); os.IsNotExist(err) {
		Write(newConfig(dirPath))
	} else {
		fmt.Println("File exists!")
	}

	Read(dirPath)
}

func newConfig(dirPath string) *Config {
	var config Config

	config.Home = dirPath

	config.Server.Address = "127.0.0.1"
	config.Server.Port = "8443"
	config.Server.TLScert = path.Join(dirPath, "pki", "self-signed_cert.pem")
	config.Server.TLSkey = path.Join(dirPath, "pki", "self-signed_key.pem")

	config.Log.LogPath = path.Join("$home", "log", alertLog) //назначаем переменной значение
	config.Log.LogLevel = 4

	config.Sqlite.DbDriver = "sqlite3"
	config.Sqlite.DbFile = path.Join(dirPath, "bin", "storage.db")

	return &config
}

func Write(config *Config) {
	data, err := yaml.Marshal(&config)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path.Join(config.Home, "conf", configFile), data, 0660)
	if err != nil {
		panic(err)
	}
}

func Read(dirPath string) Config {
	var config Config
	confPath := path.Join(dirPath, "conf", configFile)

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
