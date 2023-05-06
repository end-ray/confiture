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
}

func createStructure(dirPath string) {
	err := os.MkdirAll(path.Join(dirPath, "bin"), 0775)                //если не существует, создаем каталог "bin"
	err = os.MkdirAll(path.Join(dirPath, "log"), 0775)                 //если не существует, создаем каталог "log"
	err = os.MkdirAll(path.Join(dirPath, "assets"), 0775)              //если не существует, создаем каталог "Assets"
	err = os.MkdirAll(path.Join(dirPath, "assets", "templates"), 0775) //если не существует, создаем каталог "Templates"
	err = os.MkdirAll(path.Join(dirPath, "assets", "css"), 0775)       //если не существует, создаем каталог "CSS"
	err = os.MkdirAll(path.Join(dirPath, "assets", "js"), 0775)        //если не существует, создаем каталог "JavaScripts"

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
