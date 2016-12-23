// Copyright 2015 The Tango Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package event

import (
	//	"bytes"
	//	"net/http"
	//	"net/http/httptest"
	"os"
	"runtime/pprof"
	"testing"
	"webgo"
)

type TAction struct {
	TEvent
	//Id string
}

func (action TAction) Get(hd *webgo.THandler) {
	hd.RespondString("Get")
}

func (action TAction) Before(hd *webgo.THandler) {
	hd.RespondString("Before")
}

func (action TAction) After(hd *webgo.THandler) {
	hd.Logger.Info("After")
}

func TestSession(t *testing.T) {
	f, _ := os.Create("profile_file")
	pprof.StartCPUProfile(f)     // 开始cpu profile，结果写到文件f中
	defer pprof.StopCPUProfile() // 结束profile
	r2 := webgo.NewServer("")
	r2.Url("/", TAction.Get)
	r2.RegisterMiddleware(NewEvent())
	go r2.Listen()

	<-make(chan int)
}
