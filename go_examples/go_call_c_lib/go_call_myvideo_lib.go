package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lmyvideo
#include "myvideo.h"
*/
import "C"

//import "fmt"

func main() {
   cmd := C.CString("ffmpeg -i ./xxx/*.png ./xxx/yyy.mp4")
   C.exeFFmpegCmd(cmd)
}

// 作者：林冠宏
// 链接：https://juejin.im/post/5a62f7cff265da3e4c07e0ab
// 来源：掘金
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
