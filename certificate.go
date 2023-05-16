package conf

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func initCertificate(config Config, targetName *string) {

	certPath := config.Server.TLScert
	keyPath := config.Server.TLSkey

	createTLSCert(certPath, keyPath, targetName)
}

func createTLSCert(certPath string, keyPath string, targetName *string) {
	// Генерировать закрытый ключ RSA 2048 бита
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Подготовить основные сведения самоподписного сертификата
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: *targetName},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
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

	err = pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if err != nil {
		panic(err)
	}

	certFile.Close()

	// кодируем приватный ключ в PEM формат
	privateKeyFile, err := os.Create(keyPath)
	if err != nil {
		panic(err)
	}

	err = pem.Encode(privateKeyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	if err != nil {
		panic(err)
	}

	privateKeyFile.Close()

	// выводим кодированный сертификат и приватный ключ в консоль
	certBytes, _ := os.ReadFile(certPath)
	keyBytes, _ := os.ReadFile(keyPath)
	certBase64 := base64.StdEncoding.EncodeToString(certBytes)
	keyBase64 := base64.StdEncoding.EncodeToString(keyBytes)
	println("CERTIFICATE:")
	println(certBase64)
	println("PRIVATE KEY:")
	println(keyBase64)
}
