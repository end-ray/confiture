package conf

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

const (
	configFile = "config.yaml"
	alertLog   = "alert.log"
)

type Configuration struct {
	Home   string `yaml:"Home"`
	Log    RayLog `yaml:"Log"`
	Sqlite Sqlite `yaml:"SQLite"`
}

type Sqlite struct {
	DbDriver string `yaml:"dbDriver"`
	DbPatch  string `yaml:"dbPatch"`
	DbFile   string `yaml:"dbFile"`
}

type RayLog struct {
	LogPath  string `yaml:"logPath"`
	LogLevel uint8  `yaml:"logLevel"`
}

var config Configuration
var home string

func init() {
	initHome()
	initConfig()
}

func initHome() {
	home, _ = filepath.Abs(filepath.Dir(os.Args[0])) // определяем абсолютный путь запущенного файла
}

func initConfig() {
	_, err := os.Stat(path.Join(home, configFile))

	if err == nil { // если configFile не существует
		createConfig()
	} else { // если configFile существует

	}
}

func createConfig() {
	createStructure()

	file, err := os.OpenFile(path.Join(home, configFile), os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0660)
	defer file.Close()

	// Параметры логирования по умолчанию
	config.Log.LogPath = path.Join(home, "log", alertLog) //назначаем переменной значение
	config.Log.LogLevel = 4
	//initRayLog()

	// Параметры SQLite по умолчанию
	config.Sqlite.DbDriver = "sqlite3"
	config.Sqlite.DbPatch = home
	config.Sqlite.DbFile = "storage.db"

	data, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	file.Write(data)
}

func createStructure() {
	err := os.MkdirAll(path.Join(home, "log"), 0775) //если не существует, создаем каталог "log"
	if err != nil {
		fmt.Println(err)
	}
}

func ReadConfig() (conf Configuration) {
	file, err := os.ReadFile(path.Join(home, configFile))
	if err != nil {
		log.Fatal(err)
	}

	// parse the YAML stored in the byte slice into the struct
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}
