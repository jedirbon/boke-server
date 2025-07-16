package user_api

import (
	"boke-server/common/res"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func saveKeyToFile(filePath string, data []byte) {
	if err := ensureFileRemoved(filePath); err != nil {
		log.Fatal("删除文件失败: ", err)
		return
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_EXCL, 0)
	if err != nil {
		log.Fatal("打开文件失败：", err.Error())
		return
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
		return
	}
}

// 确保文件被删除
func ensureFileRemoved(filePath string) error {
	// 检查文件是否存在
	if _, err := os.Stat(filePath); err == nil {
		// 文件存在，尝试删除
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("无法删除文件 %s: %w", filePath, err)
		}
	} else if !os.IsNotExist(err) {
		// 不是"文件不存在"的其他错误
		return fmt.Errorf("检查文件状态失败: %w", err)
	}
	return nil
}
func CreateKeys() {
	// 生成 2048 位 RSA 密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// 将私钥编码为 PEM 格式
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// 将公钥编码为 PEM 格式
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	// 保存到文件
	go saveKeyToFile("key/private.pem", pem.EncodeToMemory(privateKeyPEM))
	go saveKeyToFile("key/public.pem", pem.EncodeToMemory(publicKeyPEM))
}
func InitKeys() {
	// 读取私钥
	privKeyBytes, err := ioutil.ReadFile("key/private.pem")
	if err != nil {
		log.Fatal("读取私钥失败:", err)
	}

	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privKeyBytes)
	if err != nil {
		log.Fatal("解析私钥失败:", err)
	}

	// 读取公钥
	pubKeyBytes, err := ioutil.ReadFile("key/public.pem")
	if err != nil {
		log.Fatal("读取公钥失败:", err)
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(pubKeyBytes)
	if err != nil {
		log.Fatal("解析公钥失败:", err)
	}
}

func (UserApi) GetPublicKey(c *gin.Context) {
	InitKeys()
	pubKeyBytes, err := ioutil.ReadFile("key/public.pem")
	if err != nil {
		res.FailedMsg("无法读取公钥", c)
		return
	}
	// 去除所有换行符
	pemStr := string(pubKeyBytes)
	pemStr = strings.ReplaceAll(pemStr, "\n", "")
	res.Ok(pemStr, c)
}

// 解密函数
func DecryptPassword(encrypted string) (string, error) {
	encryptedData, err := base64.StdEncoding.DecodeString(encrypted)
	fmt.Println("密码:", encryptedData)
	fmt.Println("私钥:", PrivateKey)
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, PrivateKey, []byte(encryptedData))
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}
