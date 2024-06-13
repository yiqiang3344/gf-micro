package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"yijunqiang/gf-micro/user/internal/utility/mcache"
	"yijunqiang/gf-micro/user/internal/utility/mstring"
)

type tokenCache struct {
	Uid    string `json:"uid"`
	Nbf    int64  `json:"nbf"`
	Iat    int64  `json:"iat"`
	Exp    int64  `json:"exp"`
	Token  string `json:"token"`
	Secret string `json:"secret"`
}

const CachePrefix = "user:jwt:"

func getCacheKey(no string) string {
	return CachePrefix + no
}

func GetCacheToken(ctx context.Context, no string) (token string, err error) {
	var (
		cache tokenCache
	)
	res, err := mcache.RedisCache().Get(ctx, getCacheKey(no))
	if err != nil {
		return
	}
	if res.String() == "" {
		return
	}
	err = json.Unmarshal(res.Bytes(), &cache)
	if err != nil {
		return
	}
	token = cache.Token
	return
}

func Token(ctx context.Context, no string) (tokenString string, err error) {
	var (
		duration  time.Duration
		key       = getCacheKey(no)
		cache     tokenCache
		tokenData jwt.MapClaims
	)
	duration, err = time.ParseDuration("24h")
	if err != nil {
		return
	}
	cache = tokenCache{
		Uid: no,
		Nbf: time.Now().Unix(),
		Iat: time.Now().Unix(),
		Exp: time.Now().Add(duration).Unix(),
	}
	tokenData = jwt.MapClaims{
		"uid": cache.Uid,
		"nbf": cache.Nbf,
		"iat": cache.Iat,
		"exp": cache.Exp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenData)
	if err != nil {
		return
	}
	cache.Secret = mstring.RandomString(32)
	tokenString, err = token.SignedString([]byte(cache.Secret))
	if err != nil {
		return
	}
	cache.Token = tokenString
	tokenCacheJson, err := json.Marshal(cache)
	if err != nil {
		return
	}
	err = mcache.RedisCache().Set(ctx, key, string(tokenCacheJson), duration)
	if err != nil {
		return
	}
	return
}

func Parse(ctx context.Context, tokenString string) (no string, err error) {
	var (
		cache tokenCache
	)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (ret interface{}, err error) {
		if c, ok := token.Claims.(jwt.MapClaims); ok {
			no = fmt.Sprintf("%v", c["uid"])
		} else {
			err = jwt.ErrInvalidType
			return
		}
		res, err := mcache.RedisCache().Get(ctx, getCacheKey(no))
		if err != nil {
			return
		}
		if res.String() == "" {
			err = jwt.ErrTokenExpired
			return
		}
		err = json.Unmarshal(res.Bytes(), &cache)
		if err != nil {
			return
		}
		ret = []byte(cache.Secret)
		return
	})

	if errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, jwt.ErrInvalidType) {
		err = errors.New("token不合法")
		return
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		err = errors.New("token已失效")
		return
	} else if token.Valid {
		return
	}
	err = errors.New("token异常")
	return
}

func Clear(ctx context.Context, no string) (err error) {
	_, err = mcache.RedisCache().Remove(ctx, getCacheKey(no))
	return
}
