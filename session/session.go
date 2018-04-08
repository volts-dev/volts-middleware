package session

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	//	"strings"
	"sync"
	"time"

	"github.com/VectorsOrigin/cacher"
	"github.com/VectorsOrigin/logger"
	"github.com/VectorsOrigin/utils"
	"github.com/VectorsOrigin/web"
)

type (
	// memory session store.
	// it saved sessions in a map in memory.
	TStore struct {
		sid string //session id
		//timeAccessed time.Time                   //last access time
		value map[interface{}]interface{} //session store
		lock  sync.RWMutex
	}

	// 提供存储Session列表
	TSessionManager struct {
		CookieName        string
		EnableSetCookie   bool
		Secure            bool
		SessionIDHashFunc string
		SessionIDHashKey  string
		CookieMaxAge      int
		ProviderConfig    string
		Domain            string
		savePath          string

		default_cacher string
		cacher         map[string]cache.ICacher
		lock           sync.RWMutex
	}
	//提供API控制
	// TSession
	TSession struct {
		cacher  string
		sid     string
		manager *TSessionManager
	}
)

// set value to memory session
func (self *TStore) Write(key, value interface{}) error {
	self.lock.Lock()
	self.value[key] = value
	self.lock.Unlock()
	return nil
}

// get value from memory session by key
func (self *TStore) Read(key interface{}) interface{} {
	self.lock.RLock()
	defer self.lock.RUnlock()
	if v, ok := self.value[key]; ok {
		return v
	} else {
		return nil
	}
}

// delete in memory session by key
func (self *TStore) Delete(key interface{}) bool {
	self.lock.Lock()
	delete(self.value, key)
	self.lock.Unlock()
	return true
}

// clear all values in memory session
func (self *TStore) Clear() error {
	self.lock.Lock()
	self.value = make(map[interface{}]interface{})
	self.lock.Unlock()
	return nil
}

// get this id of memory session store
func (self *TStore) ID() string {
	return self.sid
}

// Implement method, no used.
func (self *TStore) SessionRelease(w http.ResponseWriter) {
}

/**********************************************************************/
// get memory session store by sid
func (self *TSessionManager) Get(cacher, id, key string) interface{} {
	self.lock.RLock()
	ck := self.cacher[cacher]
	self.lock.RUnlock()
	if ck != nil {
		if store, ok := ck.Get(id).(*TStore); ok && store != nil {
			return store.Read(key)
		}
	}

	return nil
}

func (self *TSessionManager) Set(cacher, id, key string, val interface{}) {
	self.lock.RLock()
	ck := self.cacher[cacher]
	self.lock.RUnlock()
	if ck != nil {
		if store, ok := ck.Get(id).(*TStore); ok && store != nil {
			store.Write(key, val)
		} else {
			store = &TStore{sid: id, value: make(map[interface{}]interface{})}
			store.Write(key, val)

			ck.Put(id, store)
		}
	}
}

func (self *TSessionManager) Del(cacher, id, key string) bool {
	self.lock.RLock()
	ck := self.cacher[cacher]
	self.lock.RUnlock()
	if ck != nil {
		if store, ok := ck.Get(id).(*TStore); ok && store != nil {
			store.Delete(key)
			return true
		}
	}

	return false
}

func (self *TSessionManager) __New(cacher, id string, store *TStore) *TStore {

	/*	self.lock.Lock() //#
		element := self.list.PushBack(lNewSess)
		self.sessions[aId] = element
		self.lock.Unlock() //#
		return lNewSess*/
	if ck, has := self.cacher[cacher]; has {
		ck.Put(id, store)
	}

	return store
}

// check session store exist in memory session by sid
func (self *TSessionManager) Contains(cacher, key string) bool {
	self.lock.RLock()
	ck := self.cacher[cacher]
	self.lock.RUnlock()
	if ck != nil {
		return ck.Contains(key)
	}

	return false

}

func (self *TSessionManager) contains(key string) bool {
	self.lock.RLock()
	defer self.lock.RUnlock()

	for _, cacher := range self.cacher {
		if cacher.Get(key) != nil {
			return true
		}
	}

	return false
}

/**********************************************************************/
/*
保留数据使用新ID
*/
/*
// generate new sid for session store in memory session
func (self *TSessionManager) Reset(oldsid, sid string) (TSessionManager, error) {
	self.lock.RLock()
	if element, ok := self.sessions[oldsid]; ok {
		go self.SessionUpdate(oldsid)
		self.lock.RUnlock()
		self.lock.Lock()
		element.Value.(*TStore).sid = sid
		self.sessions[sid] = element
		delete(self.sessions, oldsid)
		self.lock.Unlock()
		return element.Value.(*TStore), nil
	} else {
		self.lock.RUnlock()
		self.lock.Lock()
		newsess := &TStore{sid: sid, timeAccessed: time.Now(), value: make(map[interface{}]interface{})}
		element := self.list.PushBack(newsess)
		self.sessions[sid] = element
		self.lock.Unlock()
		return newsess, nil
	}
}
*/

// delete session store in memory session by id
func (self *TSessionManager) Destroy(cacher, id string) error {
	self.lock.RLock()
	ck := self.cacher[cacher]
	self.lock.RUnlock()
	if ck != nil {
		return ck.Remove(id)
	}
	return nil
}

// get count number of memory session
func (self *TSessionManager) Count(cacher string) int {
	self.lock.RLock()
	ck := self.cacher[cacher]
	self.lock.RUnlock()
	if ck != nil {
		return ck.Len()
	}

	return 0
}

func (self *TSessionManager) Cacher(name string) bool {
	self.lock.RLock()
	ck := self.cacher[name]
	self.lock.RUnlock()
	if ck != nil {
		return true
	}

	return false
}

// expand time of session store by id in memory session
func (self *TSessionManager) UpdateTime(cacher, id string) error {
	self.lock.RLock()
	ck := self.cacher[cacher]
	self.lock.RUnlock()
	if ck != nil {
		if ck.Get(id) != nil {
			return nil
		}
	}

	return errors.New("id is not in sessions.")

}

/*
生成新session唯一id
组合：随机字符+时间毫秒+IP
*/
// generate session id with rand string, unix nano time, remote addr by hash function.
func (self *TSessionManager) new_id(r *http.Request) (RSid string) {
	bs := make([]byte, 32)
	if n, err := io.ReadFull(rand.Reader, bs); n != 32 || err != nil {
		bs = utils.RandomCreateBytes(32)
	}
	RSid = fmt.Sprintf("%s%d%s", r.RemoteAddr, time.Now().UnixNano(), bs)
	if self.SessionIDHashFunc == "md5" {
		h := md5.New()
		h.Write([]byte(RSid))
		RSid = hex.EncodeToString(h.Sum(nil))
	} else if self.SessionIDHashFunc == "sha1" {
		h := hmac.New(sha1.New, []byte(self.SessionIDHashKey))
		fmt.Fprintf(h, "%s", RSid)
		RSid = hex.EncodeToString(h.Sum(nil))
	} else {
		h := hmac.New(sha1.New, []byte(self.SessionIDHashKey))
		fmt.Fprintf(h, "%s", RSid)
		RSid = hex.EncodeToString(h.Sum(nil))
	}
	return
}

/*
链接设置Http请求
# 判断是否有HasSession
# 创建Session
# 返回Session
// Start session. generate or read the session id from http request.
// if session id exists, return SessionStore with this id.
*/
func (self *TSession) Request(act interface{}, hd *web.THandler) {
	var lSessionId string
	lCookie, err := hd.Request.Cookie(self.manager.CookieName)
	//fmt.Println("config.CookieName", lCookie, err)
	// 当无session的cookie新建一个
	if err != nil || lCookie.Value == "" {
		lSessionId = self.manager.new_session(hd.Request, hd.Response)
	} else {
		//fmt.Println("config.CookieName", lCookie.Name, lCookie.Value)
		lSessionId, _ = url.QueryUnescape(lCookie.Value)

		if !self.manager.contains(lSessionId) {
			lSessionId = self.manager.new_session(hd.Request, hd.Response)
		}
	}

	self.sid = lSessionId //update ID
	self.cacher = self.manager.default_cacher
	return
}

func (self *TSession) Response(act interface{}, hd *web.THandler) {

}

func (self *TSession) Panic(act interface{}, hd *web.THandler) {

}

func (self *TSession) Id() string {
	return self.sid
}

func (self *TSession) SetId(id string) {
	self.sid = id
}

func (self *TSession) CookieName() string {
	return self.manager.CookieName
}

func (self *TSession) Cacher(name string) *TSession {
	if self.manager.Cacher(name) {
		self.cacher = name
		return self
	}

	return nil
}

// check session store exist in memory session by sid
func (self *TSession) Contains(id string) bool {
	return self.manager.Contains(self.cacher, id)
}

func (self *TSession) Delete(id string) bool {
	return self.manager.Destroy(self.cacher, id) == nil
}

func (self *TSession) Get(key string) interface{} {
	return self.manager.Get(self.cacher, self.sid, key)
}

func (self *TSession) Set(key string, value interface{}) {
	self.manager.Set(self.cacher, self.sid, key, value)
}

func (self *TSession) Del(key string) bool {
	return self.manager.Del(self.cacher, self.sid, key)
}

/*
func (self *TSession) Release() {
	session.manager.Invalidate(session.rw, session)
}


func (self *TSession) IsValid() bool {
	return session.manager.Generator.IsValid(session.id)
}

func (self *TSession) SetMaxAge(maxAge time.Duration) {
	//self.maxAge = maxAge
}
*/

func (self *TSessionManager) new_session(lReq *http.Request, lRsp http.ResponseWriter) string {
	lSessionId := self.new_id(lReq) // 生成新ID
	//self.New(lSessionId)            //根据新ID创建新Sess
	lCookie := &http.Cookie{
		Name:     self.CookieName,
		Value:    url.QueryEscape(lSessionId),
		Path:     "/",
		HttpOnly: true,
		Secure:   self.Secure,
		Domain:   self.Domain}
	if self.CookieMaxAge >= 0 {
		lCookie.MaxAge = self.CookieMaxAge
	}
	if self.EnableSetCookie {
		http.SetCookie(lRsp, lCookie)
	}
	lReq.AddCookie(lCookie)

	return lSessionId
}

/*
func (self *TSession) Sessions() *TSessionManager {
	return self.store
}

// 重置
func (self *TSession) SetSession(s *TSession) {
	self.sid = s.sid
	//self.maxAge = s.maxAge
	self.manager = s.manager
	//self.rw = s.rw
}
*/

/*
//返回必须是TSession
	TAction struct {
		TSession
	}
*/
func NewSession(config string, cacher ...cache.ICacher) *TSession {
	cfg := make(map[string]interface{})
	if config != "" {
		err := json.Unmarshal([]byte(config), &cfg)
		if err != nil {
			logger.Dbg("NewSession Unmarshal", err, config)
			return nil
		}
	}

	lSessions := &TSession{
		manager: &TSessionManager{
			cacher:            make(map[string]cache.ICacher),
			CookieName:        "Session",
			SessionIDHashFunc: "sha1",
			EnableSetCookie:   true,
		},
	}

	if len(cacher) > 0 {

		// # 载入缓存器
		for idx, ck := range cacher {
			// #默认Cacher
			if idx == 0 {
				lSessions.manager.default_cacher = ck.Type()
			}

			lSessions.manager.cacher[ck.Type()] = ck
		}

	} else {
		ck, err := cache.NewCacher("memory", config)
		if err != nil {
			panic("NewCacher error:" + err.Error())
		}

		lSessions.manager.default_cacher = ck.Type()
		lSessions.manager.cacher[ck.Type()] = ck
	}

	if cookie_name, ok := cfg["cookie_name"].(string); ok {
		lSessions.manager.CookieName = cookie_name
	}
	if enable_set_cookie, ok := cfg["enable_set_cookie"].(bool); ok {
		lSessions.manager.EnableSetCookie = enable_set_cookie
	}
	if secure, ok := cfg["secure"].(bool); ok {
		lSessions.manager.Secure = secure
	}
	if session_id_hash_func, ok := cfg["session_id_hash_func"].(string); ok {
		lSessions.manager.SessionIDHashFunc = session_id_hash_func
	}
	if session_id_hash_key, ok := cfg["session_id_hash_key"].(string); ok {
		lSessions.manager.SessionIDHashKey = session_id_hash_key
	}
	if cookie_max_age, ok := cfg["cookie_max_age"].(int); ok {
		lSessions.manager.CookieMaxAge = cookie_max_age
	}

	return lSessions
}

func NewStore() *TStore {
	return &TStore{
		value: make(map[interface{}]interface{}),
	}
}
