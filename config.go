package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
	"path/filepath"
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

func initConfig(dirPath string) {
	//_, err := os.Stat(path.Join(home, configFile))

	//if os.IsNotExist(err) { // если configFile не существует
	//	createConfig()
	//} else { // если configFile существует

	//}
	createConfig(dirPath)

}

func createConfig(dirPath string) {
	confPath := filepath.Join(dirPath, configFile)
	file, err := os.Create(confPath) //Открыть файл для записи
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
	}

	//Домашний каталог
	config.Home = dirPath

	// Параметры логирования по умолчанию
	config.Log.LogPath = path.Join("$Home", "log", alertLog) //назначаем переменной значение
	config.Log.LogLevel = 4
	//initRayLog()

	// Параметры Web по умолчанию
	config.Web.Port = "8000"
	config.Web.Assets = path.Join("$Home", "assets")

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
