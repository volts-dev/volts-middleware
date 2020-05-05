package event

import (
	"os"
	"runtime/pprof"
	"testing"

	"github.com/volts-dev/volts/server"
)

type TAction struct {
	TEvent
}

func (action TAction) index(hd *server.TWebHandler) {
	hd.Respond([]byte("index"))
}

func (action TAction) Before(hd *server.TWebHandler) {
	hd.Info("Before")
	hd.Respond([]byte("Before"))
}

func (action TAction) After(hd *server.TWebHandler) {
	hd.Info("After")
	hd.Respond([]byte("After"))
}

func (action TAction) Panic(hd *server.TWebHandler) {
	hd.Info("Panic")
	hd.Respond([]byte("Panic"))
}
func TestSession(t *testing.T) {
	f, _ := os.Create("profile_file")
	pprof.StartCPUProfile(f)     // 开始cpu profile，结果写到文件f中
	defer pprof.StopCPUProfile() // 结束profile
	srv := server.NewServer()
	srv.RegisterMiddleware(NewEvent())
	srv.Url("GET", "/", TAction.index)

	// serve as a http server
	//go func() {
	err := srv.Listen("http")
	if err != nil {
		t.Fatal(err)
	}
	//}()
	//<-make(chan int)
}
