/**
* @Author: 18209
* @Description:
* @File:  minioInit_test
* @Version: 1.0.0
* @Date: 2022/6/1 9:22
 */

package model

import (
	"fmt"
	"os"
	"testing"
	"vedio/conf"
)

func TestUploadFile(t *testing.T) {
	conf.LoadConfig()
	InitMinIo()
	//UploadFile(bucketName, objectName string, reader io.Reader, objectSize int64) (ok bool)
	filePath := "D:\\Goprogram\\minIOClientTest\\object\\test.mp4"
	open, _ := os.Open(filePath)
	defer open.Close()
	stat, _ := open.Stat()
	file := UploadFile("vedio", "testUpload.mp4", open, stat.Size())
	fmt.Println(file)
}
