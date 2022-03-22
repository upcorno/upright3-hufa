package utils

import (
	"law/conf"
	"fmt"
	"math/rand"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

func FileUploadAuthSTS(basePath string) (*FileUploadConfInfo, error) {
	fileUploadConfInfo := &FileUploadConfInfo{}
	//构建一个阿里云客户端, 用于发起请求。
	client, err := sts.NewClientWithAccessKey(conf.App.Oss.RegionId, conf.App.Oss.AccessKeyId, conf.App.Oss.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"
	//设置参数。关于参数含义和设置方法，请参见API参考。
	request.RoleArn = conf.App.Oss.RoleArn
	request.RoleSessionName = "device-001"
	authPath := genAuthPath(basePath)
	request.Policy = genPolicyDoc(conf.App.Oss.BucketName, authPath)
	//发起请求，并得到响应。
	response, err := client.AssumeRole(request)
	if err != nil {
		return nil, err
	}
	fileUploadConfInfo.AccessKeyId = response.Credentials.AccessKeyId
	fileUploadConfInfo.AccessKeySecret = response.Credentials.AccessKeySecret
	fileUploadConfInfo.StsToken = response.Credentials.SecurityToken
	fileUploadConfInfo.Host = conf.App.Oss.Host
	fileUploadConfInfo.Expire = response.Credentials.Expiration
	fileUploadConfInfo.Bucket = conf.App.Oss.BucketName
	fileUploadConfInfo.Region = conf.App.Oss.RegionId
	fileUploadConfInfo.Dir = authPath
	return fileUploadConfInfo, err
}

func genPolicyDoc(bucketName string, path string) (policy string) {
	policy = fmt.Sprintf(`{"Statement":[
        {"Effect":"Allow","Action":["oss:ListObjects"],"Resource":
        ["acs:oss:*:*:%s"],"Condition":{"StringLike":{"oss:Prefix":"%s*"}}},
        {"Effect":"Allow","Action":"oss:*","Resource":"acs:oss:*:*:%s/%s*"}
        ],"Version":"1"}`, bucketName, path, bucketName, path)
	return
}

//返回 basePath="path1" return path1/202101/01/21312
func genAuthPath(basePath string) (authPath string) {
	nowTime := time.Now()
	authPath = fmt.Sprintf("%s/%d%d/%d/%d", basePath, nowTime.Year(), nowTime.Month(), nowTime.Day(), rand.Intn(99999))
	return
}

type FileUploadConfInfo struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	StsToken        string `json:"stsToken"`
	Host            string `json:"host"`
	Expire          string `json:"expire"`
	Bucket          string `json:"bucket"`
	Dir             string `json:"dir"`
	Region          string `json:"region"`
}
