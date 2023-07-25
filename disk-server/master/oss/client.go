package oss

import (
	"master/global"
	"sync"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	once   sync.Once
	client *OssClient
)

type OssClient struct {
	client *oss.Client
	bucket *oss.Bucket
}

func GetClient() *OssClient {
	once.Do(func() {
		client = &OssClient{}
	})
	return client
}

func (o *OssClient) Init() {
	var err error
	o.client, err = oss.New(global.Config.OSSEndpoint, global.Config.OSSAccessKeyId, global.Config.OSSAccessKeySecret)
	if err != nil {
		panic(err)
	}

	o.bucket, err = o.client.Bucket(global.Config.OSSBucket)
	if err != nil {
		panic(err)
	}
}
