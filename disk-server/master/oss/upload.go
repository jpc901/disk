package oss

import (
	"encoding/json"
	"fmt"
	"master/global"

	"github.com/jpc901/disk-common/model"
	mq "github.com/jpc901/disk-common/rabbitmq"
	log "github.com/sirupsen/logrus"
)

func (o *OssClient) UploadFile() {
	chanMsg := make(chan []byte, 100)

	go mq.Receive(global.Config.RabbitMQConfig, chanMsg)

	for {
		select {
		case msg := <-chanMsg:
			fileModel := &model.FileMeta{}
			err := json.Unmarshal(msg, fileModel)
			if err != nil {
				log.Error(fmt.Printf("oss client upload file: [%s]", string(msg)))
			} else {
				o.uploadToOss(fileModel)
			}
			fmt.Println("oss client upload file: [" + string(msg) + "]")
		}
	}
}

func (o *OssClient) uploadToOss(fileModel *model.FileMeta) {
	err := o.bucket.PutObjectFromFile(fileModel.FileSha1, fileModel.Location)
	if err != nil {
		log.Error(err)
	}
}
