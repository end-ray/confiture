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

// var config Config
var home string

func InitConfiture(targetName string) {
	dirPath, exeName := initHome()

	//Проверка находится ли исполняемый файл в каталоге bin
	dirName := filepath.Base(dirPath)
	if dirName == "bin" {
		return
	}

	initStructure(dirPath, exeName, targetName)
	initConfig(dirPath)
}

func initHome() (string, string) {
	exePath, err := os.Executable() // Получаем абсолютный путь к текущему исполняемому файлу
	if err != nil {
		fmt.Println("Ошибка получения пути к исполняемому файлу:", err)
	}

	dirPath := filepath.Dir(exePath)  // определяем абсолютный путь запущенного файла
	exeName := filepath.Base(exePath) // Получаем имя файла из пути
	home = dirPath

	return dirPath, exeName
}

func Read() (conf Config) {
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

//func Write() {
//	data, _ := yaml.Marshal(&config)
//	os.WriteFile(path.Join(home, configFile), data, 0660)
//}
