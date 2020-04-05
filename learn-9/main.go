package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
)

type Manager struct {
	cookieName string
	lock       sync.Mutex
	provider   Provider
	maxLifeTime int64
}




func NewManager(provideName, cookieName string, maxLifetime int64) (*Manager, error) {
	provider,ok := provides[provideName]
	if !ok{
		return  nil,fmt.Errorf("session:unknown provide %q (forgotten import?)",provideName)
	}

	return &Manager{provider:provider,cookieName:cookieName,maxLifeTime:maxLifetime},nil
}

var globalSessions *Manager

func init()  {
	manager, e := NewManager("memory", "gosessionid", 3600)
	if e != nil{
		log.Fatal("初始化session错误")
	}
	globalSessions = manager
}

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

var provides = make(map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session:Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session:Register called twice for privide" + name)
	}
	provides[name] = provider
}


type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

func (manager *Manager) sessionId() string {
	b := make([]byte,32)
	if _, err := io.ReadFull(rand.Reader, b);err != nil{
		return ""
	}
	return  base64.RawURLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(w http.ResponseWriter,r *http.Request)(session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, e := r.Cookie(manager.cookieName)
	if e != nil && cookie.Value == ""{
		sid := manager.sessionId()
		session, _ := manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name:manager.cookieName,Value:url.QueryEscape(sid),Path:"/",HttpOnly:true,MaxAge:int(manager.maxLifeTime)}
		http.SetCookie(w,&cookie)
		return  session
	}else {
		sid,_ := url.QueryUnescape(cookie.Value)
		session,_ := manager.provider.SessionRead(sid)
		return  session
	}
}