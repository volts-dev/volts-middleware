// Copyright 2015 The Tango Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package session

import (
	"time"
	//	"bytes"
	//	"net/http"
	//	"net/http/httptest"
	//"reflect"
	"testing"
	"webgo"
	"webgo/cache"
)

type SessionAction struct {
	Session *TSession
	Id      string
}

func (action SessionAction) Get(hd *webgo.THandler) {
	//webgo.Warn("Get", action.Session, action.Session == nil, reflect.ValueOf(action.Session))
	//webgo.Warn("Get", reflect.ValueOf(action.Session).Interface().(*TMemorySession))
	//webgo.Warn("Get", action.Id)
	//webgo.Warn("Get", action.Session.Id())
	action.Session.Set("aa", action.Session.Id())
	//ss := reflect.ValueOf(action.Session).Interface().(*TMemorySession)
	hd.RespondString(action.Session.Get("aa").(string))

}

func TestSession(t *testing.T) {
	r2 := webgo.NewServer("")
	r2.Url("/", SessionAction.Get)
	ck, _ := cache.NewCacher("memory", `{"interval":5,"expired":30}`)
	r2.RegisterMiddleware(
		NewSession(
			`{"interval":5,"expired":10}`,
			ck,
		))
	go r2.Listen()

	for {
		<-time.After(10 * time.Second)
		ck.Clear()
	}
	<-make(chan int)
}
