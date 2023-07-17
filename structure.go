package conf

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func initStructure(dirPath string, exeName string, targetName string) {
	createStructure(dirPath)
	moveExeFile(dirPath, exeName, targetName)
	creatDbFile(dirPath)
}

func createStructure(dirPath string) {
	err := os.MkdirAll(path.Join(dirPath, "bin"), 0775) //если не существует, создаем каталог "bin"
	err = os.MkdirAll(path.Join(dirPath, "conf"), 0775) //если не существует, создаем каталог "conf"
	err = os.MkdirAll(path.Join(dirPath, "log"), 0775)  //если не существует, создаем каталог "log"
	err = os.MkdirAll(path.Join(dirPath, "pki"), 0775)  //если не существует, создаем каталог "pki"

	if err != nil {
		panic(err)
	}

}

func moveExeFile(dirPath string, exeName string, targetName string) {

	exePath := filepath.Join(dirPath, exeName)
	targetPath := filepath.Join(dirPath, "bin", targetName) // Указываем целевую директорию для перемещения файла
	err := os.Rename(exePath, targetPath)
	if err != nil {
		fmt.Println("Ошибка перемещения файла:", err)
	}
}

func creatDbFile(dirPath string) {

	filePath := path.Join(dirPath, "bin", "storage.db")

	// Проверяем наличие файла
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		// Создать или открыть файл базы данных
		dbFile, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		defer dbFile.Close()
	}

}
