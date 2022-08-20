package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
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
	Home   string   `yaml:"Home"`
	Log    AlertLog `yaml:"AlertLog"`
	Web    Web      `yaml:"Web"`
	Sqlite Sqlite   `yaml:"SQLite"`
}

type AlertLog struct {
	LogPath  string `yaml:"logPath"`
	LogLevel uint8  `yaml:"logLevel"`
}

type Web struct {
	Port   string `yaml:"Port"`
	Assets string `yaml:"Assets"`
}

type Sqlite struct {
	DbDriver string `yaml:"dbDriver"`
	DbPatch  string `yaml:"dbPatch"`
	DbFile   string `yaml:"dbFile"`
}

var config Configuration
var home string

func InitConfiture() {
	initHome()
	initConfig()
}

func initHome() {
	home, _ = filepath.Abs(filepath.Dir(os.Args[0])) // определяем абсолютный путь запущенного файла
}

func initConfig() {
	_, err := os.Stat(path.Join(home, configFile))

	if os.IsNotExist(err) { // если configFile не существует
		createConfig()
	} else { // если configFile существует

	}
}

func createConfig() {
	createStructure()

	file, err := os.OpenFile(path.Join(home, configFile), os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0660)
	defer file.Close()

	//Домашний каталог
	config.Home = home

	// Параметры логирования по умолчанию
	config.Log.LogPath = path.Join(home, "log", alertLog) //назначаем переменной значение
	config.Log.LogLevel = 4
	//initRayLog()

	// Параметры Web по умолчанию
	config.Web.Port = "8000"
	config.Web.Assets = path.Join(home, "assets", "templates", "login.html")

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
	err := os.MkdirAll(path.Join(home, "log"), 0775)                //если не существует, создаем каталог "log"
	err = os.MkdirAll(path.Join(home, "assets"), 0775)              //если не существует, создаем каталог "Assets"
	err = os.MkdirAll(path.Join(home, "assets", "templates"), 0775) //если не существует, создаем каталог "Templates"
	err = os.MkdirAll(path.Join(home, "assets", "css"), 0775)       //если не существует, создаем каталог "CSS"
	err = os.MkdirAll(path.Join(home, "assets", "js"), 0775)        //если не существует, создаем каталог "JavaScripts"

	if err != nil {
		fmt.Println(err)
	}

}

func Read() (conf Configuration) {
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

func Write() {
	data, _ := yaml.Marshal(&config)
	os.WriteFile(path.Join(home, configFile), data, 0660)
}
