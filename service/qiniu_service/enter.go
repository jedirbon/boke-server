package qiniu_service

import (
	"boke-server/global"
	"boke-server/utils/file"
	"boke-server/utils/hash"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"github.com/qiniu/go-sdk/v7/storagev2/uptoken"
	"io"
	"time"
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

// 客户端上传
func GetQiNiuToken() (upToken string, err error) {
	mac := credentials.NewCredentials(global.Config.QiNiu.AccessKey, global.Config.QiNiu.SecretKey)
	bucket := global.Config.QiNiu.Bucket
	putPolicy, err := uptoken.NewPutPolicy(bucket, time.Now().Add(time.Duration(global.Config.QiNiu.Expiry)*time.Second))
	if err != nil {
		return
	}
	upToken, err = uptoken.NewSigner(putPolicy, mac).GetUpToken(context.Background())
	return
}
