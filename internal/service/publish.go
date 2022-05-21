package service

import (
	"TikTok/internal/log"
	"TikTok/internal/model"
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func Publish(video *model.Video) (err error) {
	result := db.Create(&video)
	err = result.Error
	if err != nil {
		log.Errorf(result.Error)
		return
	}
	return

}

//https://www.cnblogs.com/jssyjam/p/11428683.html

// SaveUploadedFile uploads the form file to specific dst.
func SaveUploadedFile(file *multipart.FileHeader, dst string) (err error, filename string) {
	src, err := file.Open()
	if err != nil {
		log.Errorf(err)
		return
	}
	defer src.Close()
	exists, err := PathExists(dst) //判断文件或者文件夹是否存在
	names := strings.Split(dst, "/")
	if err != nil {
		log.Errorf(err)
		return
	}
	if !exists {
		err = os.MkdirAll(dst, 0777)
	}
	if err != nil {
		log.Errorf(err)
		return
	}

	filename = dst + "/" + names[len(names)-1]
	out, err := os.Create(filename)
	if err != nil {
		log.Errorf(err)
		return
	}
	defer out.Close()
	// 拷贝文件
	_, err = io.Copy(out, src)
	if err != nil {
		log.Errorf(err)
		return
	}
	return
}

// GetSnapshot 生成视频缩略图并保存（作为封面）
func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Errorf("生成缩略图失败：", err)
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		log.Errorf("生成缩略图失败：", err)
	}
	exists, err := PathExists(snapshotPath)

	if err != nil {
		log.Errorf("发生了异常", err)
		return
	}
	if !exists {
		os.MkdirAll(snapshotPath, 0777)
	}
	err = imaging.Save(img, snapshotPath+".jpeg")
	if err != nil {
		log.Errorf("生成缩略图失败：", err)
	}

	// 成功则返回生成的缩略图名
	names := strings.Split(snapshotPath, "/")
	snapshotName = names[len(names)-1] + ".jpeg"
	return
}

// 判断所给路径文件/文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	//isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}
