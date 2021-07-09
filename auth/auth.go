package auth

import (
	"time"

	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/basic"
	"github.com/shaj13/go-guardian/v2/auth/strategies/token"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
	"github.com/shaj13/libcache"
	_ "github.com/shaj13/libcache/fifo"
)

var Strategy union.Union
var cacheObj libcache.Cache
var TokenStrategy auth.Strategy

func InitAuth() {
	cacheObj = libcache.FIFO.New(0)
	cacheObj.SetTTL(time.Minute * 5)
	cacheObj.RegisterOnExpired(func(key, _ interface{}) {
		cacheObj.Peek(key)
	})

	basicStrategy := basic.NewCached(ValidateUser, cacheObj)
	TokenStrategy = token.New(token.NoOpAuthenticate, cacheObj)
	Strategy = union.New(TokenStrategy, basicStrategy)
}
