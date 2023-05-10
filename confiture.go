package conf

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	configFile = "config.yaml"
	alertLog   = "alert.log"
)

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

	return dirPath, exeName
}
