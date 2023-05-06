package conf

import (
	"fmt"
	"os"
	"path"
)

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
