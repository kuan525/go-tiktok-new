package ffmpeg

import (
	"bytes"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func GetFirstFrame(videoPath string) (buf *bytes.Buffer, err error) {
	buf = bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Output("pipe:", ffmpeg.KwArgs{
			"vframes": 1,        // 只输出一帧
			"format":  "image2", // 输出一系列图像文件
			"vcodec":  "mjpeg",  // 视频编码格式 Motion JPEG（MJPEG）
		}).WithOutput(buf).
		Run()

	return buf, nil
}
