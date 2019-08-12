package mp

//音乐播放模块
import (
	"fmt"
	"strings"
)

//设计一个简单的接口避免将MusicEntry中多余的信息传入
type Player interface {
	Play(Source string)
}

//播放实现，再此也可以在添加其他格式
func Play(Source, mtype string) {
	var p Player
	mtype = strings.ToUpper(mtype)
	switch mtype {
	case "MP3": //MP3格式播放
		p = &MP3Player{}
	case "WAV": //WAV格式播放
		p = &WAVPlayer{}
	default:
		fmt.Println("Unsupported music type", mtype)
		return
	}
	p.Play(Source)
}
