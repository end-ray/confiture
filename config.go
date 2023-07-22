package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"net"
	"os"
	"path"
)

type Config struct {
	Home   string   `yaml:"home"`
	Server Server   `yaml:"server"`
	Log    AlertLog `yaml:"AlertLog"`
	Sqlite Sqlite   `yaml:"SQLite"`
	Pgsql  Pgsql    `yaml:"PostgreSQL"`
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
	DbFile string `yaml:"dbFile"`
}

type Pgsql struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	DbName string `yaml:"dbname"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
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

	config.Server.Address = ipnet()
	config.Server.Port = "8443"
	config.Server.TLScert = path.Join(dirPath, "pki", "self-signed_cert.pem")
	config.Server.TLSkey = path.Join(dirPath, "pki", "self-signed_key.pem")

	config.Log.LogPath = path.Join(dirPath, "log", alertLog) //назначаем переменной значение
	config.Log.LogLevel = 4

	config.Sqlite.DbFile = path.Join(dirPath, "bin", "storage.db")

	config.Pgsql.Host = "localhost"
	config.Pgsql.Port = "5432"
	config.Pgsql.DbName = ""
	config.Pgsql.User = ""
	config.Pgsql.Passwd = ""

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

func ipnet() string {

	// Получаем список всех сетевых интерфейсов
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Не удалось получить список сетевых интерфейсов:", err)
		return "Не удалось получить список сетевых интерфейсов"
	}
	// Проходимся по каждому интерфейсу
	for _, iface := range interfaces {
		// Проверяем, что интерфейс активен и не является петлевым
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
			// Получаем список адресов для текущего интерфейса
			addrs, err := iface.Addrs()
			if err != nil {
				fmt.Println("Не удалось получить адреса для интерфейса", iface.Name, ":", err)
				continue
			}

			// Проходимся по каждому адресу
			for _, addr := range addrs {
				// Проверяем, является ли адрес IP-адресом
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
					return ipnet.IP.String()
				}
			}
		}
	}
	return "нет IPv4-адресов"
}
