package conf

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"
	"path"
	"time"
)

func initCertificate(config Config, targetName *string) {

	certPath := config.Server.TLScert
	keyPath := config.Server.TLSkey
	dirHome := config.Home

	if _, err := os.Stat(path.Join(dirHome, "pki", "self-signed_cert.pem")); os.IsNotExist(err) {
		createTLSCert(certPath, keyPath, targetName)
	} else {
		fmt.Println("File exists!")
	}
}

func createTLSCert(certPath string, keyPath string, targetName *string) {
	// Генерировать закрытый ключ RSA 2048 бита
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Заполняем данные о владельце сертификата
	subject := pkix.Name{
		CommonName:   *targetName,
		Organization: []string{"Self-Signed Inc."},
	}

	// Подготовить основные сведения самоподписного сертификата
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	// Создать самоподписанный сертификат
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		panic(err)
	}

	// кодируем сертификат в PEM формат
	certFile, err := os.Create(certPath)
	if err != nil {
		panic(err)
	}
	//defer certFile.Close()
	defer func() {
		err := certFile.Close()
		if err != nil {
			log.Printf("Error closing cert file: %v", err)
		}
	}()

	err = pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if err != nil {
		panic(err)
	}

	// кодируем приватный ключ в PEM формат
	privateKeyFile, err := os.Create(keyPath)
	if err != nil {
		panic(err)
	}

	//defer privateKeyFile.Close()
	defer func() {
		err := privateKeyFile.Close()
		if err != nil {
			log.Printf("Error closing cert file: %v", err)
		}
	}()

	err = pem.Encode(privateKeyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	if err != nil {
		panic(err)
	}
}
