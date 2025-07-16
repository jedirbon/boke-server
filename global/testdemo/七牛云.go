package main

import (
	"boke-server/core"
	"boke-server/flags"
	"boke-server/global"
	"boke-server/service/qiniu_service"
	"boke-server/utils/file"
	"boke-server/utils/hash"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"io"
)

func SendFile(fileName string) (url string, err error) {
	fmt.Println(global.Config.QiNiu.AccessKey)
	fmt.Println(global.Config.QiNiu.SecretKey)
	hashString, err := hash.FileMd5(fileName)
	if err != nil {
		return
	}

	_, suffix := file.ImageSuffixJudge(fileName)
	filename := fmt.Sprintf("%s.%s", hashString, suffix)
	mac := credentials.NewCredentials(global.Config.QiNiu.AccessKey, global.Config.QiNiu.SecretKey)
	key := fmt.Sprintf("%s/%s", global.Config.QiNiu.Prefix, filename)
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})
	err = uploadManager.UploadFile(context.Background(), fileName, &uploader.ObjectOptions{
		BucketName: global.Config.QiNiu.Bucket,
		ObjectName: &key,
		FileName:   filename,
	}, nil)
	return fmt.Sprintf("%s/%s", global.Config.QiNiu.Uri, key), err
}

func SendReader(reader io.Reader) (url string, err error) {
	accessKey := "your access key"
	secretKey := "your secret key"
	mac := credentials.NewCredentials(accessKey, secretKey)
	bucket := "if-pbl"
	key := "github-x.png"
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})
	err = uploadManager.UploadReader(context.Background(), reader, &uploader.ObjectOptions{
		BucketName: bucket,
		ObjectName: &key,
		CustomVars: map[string]string{
			"name": "github logo",
		},
		FileName: "",
	}, nil)
	return fmt.Sprintf("%s/%s", global.Config.QiNiu.Uri, key), err
}

func main() {
	flags.Parse()
	//读取配置信息
	global.Config = core.ReadConf()
	//初始话日志
	core.InitLogrus()
	/*	url, err := SendFile("uploads/images/123.jpg")
		fmt.Println("**********")
		fmt.Println(err, url)*/
	token, err := qiniu_service.GetQiNiuToken()
	if err != nil {
		fmt.Println("获取七牛云token失败")
		return
	}
	fmt.Println("**********")
	fmt.Println(token)
	fmt.Println("**********")
}
