/**
* @Author: 18209
* @Description:
* @File:  minioInit
* @Version: 1.0.0
* @Date: 2022/5/31 0:51
 */

package model

import (
	"fmt"
	_ "github.com/minio/minio-go/pkg/encrypt"
	"github.com/minio/minio-go/pkg/policy"
	"github.com/minio/minio-go/v6"
	"go.uber.org/zap"
	"io"
	"log"
	"net/url"
	"time"
	"vedio/conf"
)

var MinioClient *minio.Client

func InitMinIo() {
	// 初使化 minio client对象。false是关闭https证书校验
	minioClient, err := minio.New(conf.ConfigYaml.Minio.Endpoint, conf.ConfigYaml.Minio.AccessKeyID, conf.ConfigYaml.Minio.SecretAccessKey, false)
	if err != nil {
		log.Fatalln(err)
	}
	//客户端注册到全局变量中
	MinioClient = minioClient
	//创建一个叫userheader的存储桶。
	CreateMinIoBuket("vedio")
}

//创建minIo桶
func CreateMinIoBuket(bucketName string) {
	err := MinioClient.MakeBucket(bucketName, "us-east-1")
	if err != nil {
		//检查存储桶是否已经存在
		exists, err := MinioClient.BucketExists(bucketName)
		if err != nil {
			if exists {
				log.Println(bucketName + "exists")
			} else {
				log.Println(err)
			}
		}
		return
	}
	err = MinioClient.SetBucketPolicy(bucketName, policy.BucketPolicyReadWrite)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("successfully created " + bucketName)
}

// UploadFile 上传文件给minio指定的桶中,传入io.Reader
func UploadFile(bucketName, objectName string, reader io.Reader, objectSize int64) (ok bool) {
	n, err := MinioClient.PutObject(bucketName, objectName, reader, objectSize, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println("Successfully uploaded bytes: ", n)
	return true
}

//  GetFileUrl 获取文件url
func GetFileUrl(bucketName string, fileName string, expires time.Duration) string {
	//time.Second*24*60*60
	reqParams := make(url.Values)
	presignedURL, err := MinioClient.PresignedGetObject(bucketName, fileName, expires, reqParams)
	if err != nil {
		zap.L().Error(err.Error())
		return ""
	}
	return fmt.Sprintf("%s", presignedURL)
}
