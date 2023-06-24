package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rohanchauhan02/clean/common/util"
)

type (
	RateLimitConfig struct {
		Skipper    middleware.Skipper
		Duration   time.Duration
		MaxCounter int64
	}
)

var DefaultRateLimitConfig = RateLimitConfig{
	Skipper:    middleware.DefaultSkipper,
	Duration:   time.Duration(30) * time.Minute,
	MaxCounter: int64(10),
}

// Default rate limit is 10 hits in 30 minutes per endpoint per user Id.
func RateLimitWithDefaultRedis() echo.MiddlewareFunc {
	return RateLimitWithRedis(DefaultRateLimitConfig)
}

// Use this to use custom config for the rate limit duration and counter.
func RateLimitWithRedis(config RateLimitConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultRateLimitConfig.Skipper
	}

	rateLimitCacheNamespaceFormat := "RL:%s:%s"

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			ac := c.(*util.CustomApplicationContext)

			if ac.RedisSession == nil || ac.RedisSession.Client == nil {
				return next(c)
			}

			userData := getUserDataFromCtx(c)
			if userData == nil {
				return next(c)
			}

			endpoint := fmt.Sprintf("[%s]%s", ac.Context.Request().Method, ac.Context.Request().RequestURI)
			cacheKey := fmt.Sprintf(rateLimitCacheNamespaceFormat, endpoint, userData.Data.User.ID)

			counter := ac.RedisSession.Client.Get(cacheKey)

			if counter.Val() == "" {
				ac.RedisSession.Set(cacheKey, 1, config.Duration)
				return next(c)
			}

			newCounter := ac.RedisSession.Incr(cacheKey)

			if newCounter.Val() > config.MaxCounter {
				return ac.CustomResponse("QC-CLT-RRL-001",
					nil,
					"Too many requests",
					http.StatusTooManyRequests,
					http.StatusTooManyRequests,
					nil)
			}

			return next(c)
		}
	}
}

func getUserDataFromCtx(c echo.Context) *CheckUserV1Response {
	userDataCtx := c.Request().Context().Value(ContextUserKey)
	ctxUserData, ok := userDataCtx.(CheckUserV1Response)
	if !ok {
		return nil
	}
	return &ctxUserData
}
