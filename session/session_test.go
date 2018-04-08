// Copyright 2015 The Tango Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package session

import (
	"time"
	//"reflect"
	"testing"

	"github.com/VectorsOrigin/cacher"
	"github.com/VectorsOrigin/web"
)

type SessionAction struct {
	Session *TSession
	Id      string
}

func (action SessionAction) Get(hd *web.THandler) {
	//web.Warn("Get", action.Session, action.Session == nil, reflect.ValueOf(action.Session))
	//web.Warn("Get", reflect.ValueOf(action.Session).Interface().(*TMemorySession))
	//web.Warn("Get", action.Id)
	//web.Warn("Get", action.Session.Id())
	action.Session.Set("aa", action.Session.Id())
	//ss := reflect.ValueOf(action.Session).Interface().(*TMemorySession)
	hd.RespondString(action.Session.Get("aa").(string))

}

func TestSession(t *testing.T) {
	r2 := web.NewServer("")
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
