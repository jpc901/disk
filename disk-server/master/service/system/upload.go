package system

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"master/global"
	"mime/multipart"
	"os"
	"time"

	myDB "github.com/jpc901/disk-common/mapper"
	"github.com/jpc901/disk-common/model"
	"github.com/jpc901/disk-common/model/request"
	mq "github.com/jpc901/disk-common/rabbitmq"

	util "github.com/jpc901/disk-common/utils"
	log "github.com/sirupsen/logrus"
)

type UploadService struct{}

func (up *UploadService) UploadFile(uid int64, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		log.Errorf("file open failed, error: %v", err)
		return err
	}
	fileMeta := model.FileMeta{
		FileName: fileHeader.Filename,
		Location: model.UploadPath + fileHeader.Filename,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	defer file.Close()

	newFile, err := os.Create(fileMeta.Location)
	if err != nil {
		log.Errorf("Upload failed of create file, error: %v", err)
		return err
	}
	defer newFile.Close()

	fileMeta.FileSize, err = io.Copy(newFile, file)
	if err != nil {
		log.Errorf("Failed to save data into file,err: %v", err)
		return err
	}

	newFile.Seek(0, 0)
	fileMeta.FileSha1 = util.FileSha1(newFile)

	username, err := myDB.GetUsername(uid)
	if err != nil {
		log.Error("get username db failed")
		return err
	}
	// 更新文件表
	err = myDB.UpdateFile(fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize, fileMeta.Location)
	if err != nil {
		log.Error("update file db failed")
		return err
	}
	err = myDB.UpdateUserFile(username, fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize)
	if err != nil {
		log.Error("update user file db failed")
		return err
	}

	// 文件上传到rabbitmq
	msgJson, _ := json.Marshal(fileMeta)
	err = mq.Send(global.Config.RabbitMQConfig, msgJson)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("success [^_^]")
	return nil
}

func (up *UploadService) InitMultipleUploadFile(multipleInfo request.MultipleInitRequest) {
	global.RDB.GetConn().HSet(context.Background(),
		multipleInfo.UploadId,
		"fileName", multipleInfo.FileName,
		"fileSize", multipleInfo.FileSize,
		"chunkSize", multipleInfo.ChunkSize,
		"chunkCount", multipleInfo.ChunkCount,
		"fileHash", multipleInfo.FileHash,
	)
}

func (up *UploadService) MultipleUploadFile(multipleInfo request.UploadMultipleRequest) error {
	// 将chunk存入本地
	file, err := multipleInfo.File.Open()
	if err != nil {
		log.Error("open file failed, err", err)
		return err
	}
	defer file.Close()

	filePath := model.UploadPath + multipleInfo.UploadId + multipleInfo.CurChunk
	newFile, err := os.Create(filePath)
	if err != nil {
		log.Error("create file failed, err: ", err)
		return err
	}
	defer newFile.Close()
	io.Copy(newFile, file)

	// 将chunk信息存入redis
	global.RDB.GetConn().HSet(context.Background(),
		multipleInfo.UploadId,
		"curChunk"+multipleInfo.CurChunk, multipleInfo.CurChunk,
		"chunkSize"+multipleInfo.CurChunk, multipleInfo.ChunkSize,
	)

	log.Infof("[upload chunk success](^_^): uploadId %s, currentChunk %s", multipleInfo.UploadId, multipleInfo.CurChunk)
	return nil
}

func (up *UploadService) CheckChunkExist(queryInfo request.CheckChunkExistRequest) error {
	chunkNumInRedis, err := global.RDB.GetConn().HGet(context.Background(), queryInfo.UploadId, "curChunk"+queryInfo.CurChunk).Result()
	if err != nil {
		log.Error("query redis failed, err:", err)
		return err
	}
	if chunkNumInRedis != queryInfo.CurChunk {
		return err
	}
	return nil
}

func (up *UploadService) MergeAndSave(uid int64, info request.MultipleInitRequest) error {

	// 合并分块
	err := up.MergeChunk(model.UploadPath, info.UploadId, info.FileName, int(info.ChunkCount))
	if err != nil {
		log.Error(err)
		return err
	}

	// 删除redis数据
	global.RDB.GetConn().Del(context.Background(), info.UploadId)

	// 数据持久化
	fileMeta := model.FileMeta{
		FileName: info.FileName,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		FileSize: info.FileSize,
		FileSha1: info.FileHash,
	}
	fileMeta.Location = model.UploadPath + fileMeta.FileName

	err = FileMetaServiceApp.UpdateFileMetaDB(fileMeta)
	if err != nil {
		log.Error("update file db failed")
		return err
	}
	username, err := myDB.GetUsername(uid)
	if err != nil {
		log.Error("get username db failed")
		return err
	}
	err = FileMetaServiceApp.UpdateUserFileMetaDB(username, fileMeta.FileSha1, info.FileName, fileMeta.FileSize)
	if err != nil {
		log.Error("update user file db failed")
		return err
	}

	log.Info("multiple upload success uploadId", info.UploadId)
	return nil
}

// mergeChunk 合并分块
func (up *UploadService) MergeChunk(path, uploadId, fileName string, chunkCount int) error {
	// 创建一个新文件
	newFile, err := os.Create(path + fileName)
	if err != nil {
		log.Error(err)
		return err
	}

	defer newFile.Close()

	// 读取分块文件
	for i := 0; i < chunkCount; i++ {
		currentChunkPath := fmt.Sprintf("%s%s%d", path, uploadId, i)
		currentChunkFile, err := os.Open(currentChunkPath)
		if err != nil {
			log.Error(err)
			return err
		}

		defer currentChunkFile.Close()

		// 将分块文件写入新文件
		io.Copy(newFile, currentChunkFile)
	}
	return nil
}
