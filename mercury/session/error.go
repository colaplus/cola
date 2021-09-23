package session

import "errors"

var ErrSessionNotExist = errors.New("session not exists")
var ErrKeyNotExist = errors.New("key not exists in session")
