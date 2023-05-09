package conf

import (
	"gopkg.in/yaml.v3"
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
	//_, err := os.Stat(path.Join(home, configFile))

	//if os.IsNotExist(err) { // если configFile не существует
	//	createConfig()
	//} else { // если configFile существует

	//}
	//createConfig(dirPath)

	Write(newConfig(dirPath))
}

func newConfig(dirPath string) *Config {
	var config Config

	config.Home = dirPath

	config.Server.Address = "0.0.0.0"
	config.Server.Port = "8000"
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

//func createConfig(dirPath string) {
//	confPath := filepath.Join(dirPath, configFile)
//	file, err := os.Create(confPath) //Открыть файл для записи
//	if err != nil {
//		fmt.Println("Ошибка создания файла:", err)
//	}
//
//	//Домашний каталог
//	config.Home = dirPath
//
//	// Параметры логирования по умолчанию
//	config.Log.LogPath = path.Join("$Home", "log", alertLog) //назначаем переменной значение
//	config.Log.LogLevel = 4
//	//initRayLog()
//
//	// Параметры Web по умолчанию
//	config.Server.Port = "8000"
//	config.Server.Assets = path.Join("$Home", "assets")
//
//	// Параметры SQLite по умолчанию
//	config.Sqlite.DbDriver = "sqlite3"
//	config.Sqlite.DbPatch = home
//	config.Sqlite.DbFile = "storage.db"
//
//	data, err := yaml.Marshal(&config)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	file.Write(data)
//}
