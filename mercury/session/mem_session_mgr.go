package session

import (
	"sync"
)
import "github.com/satori/go.uuid"

type MemSessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

func NewMemorySessionMgr() SessionMgr {
	sr := &MemSessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}

	return sr
}

func (s *MemSessionMgr) Init(addr string, options ...string) (err error) {
	return
}

func (s *MemSessionMgr) Get(sessionId string) (session Session, err error) {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	session, ok := s.sessionMap[sessionId]
	if !ok {
		err = ErrSessionNotExist
		return
	}

	return
}

func (s *MemSessionMgr) CreateSession() (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	id := uuid.NewV4()
	sessionId := id.String()
	session = NewMemorySession(sessionId)

	s.sessionMap[sessionId] = session
	return
}
