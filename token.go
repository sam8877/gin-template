package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/google/uuid"
	"strings"
	"time"
)

type TokenStruct struct {
	UserId    string `json:"user_id"`    // 用户标志
	IssuedAt  int64  `json:"issued_at"`  // 签发时间
	ExpiresAt int64  `json:"expires_at"` // 失效时间
	TokenId   string `json:"token_id"`   // token唯一ID
}

// Expired token 是否失效
func (receiver *TokenStruct) Expired() bool {
	return time.Now().Unix() > receiver.ExpiresAt
}

func CreateToken(userId string, dur time.Duration, jwt_key string) (string, error) {
	// 生成token结构体
	tokenStruct := TokenStruct{
		TokenId: uuid.New().String(),
		UserId:  userId,
	}
	return CreateTokenFromTokenStruct(&tokenStruct, dur, jwt_key)
}

func CreateTokenFromTokenStruct(tokenStruct *TokenStruct, dur time.Duration, jwt_key string) (string, error) {
	tokenStruct.IssuedAt = time.Now().Unix()
	tokenStruct.ExpiresAt = time.Now().Add(dur).Unix()
	bytes, err := json.Marshal(tokenStruct)
	if err != nil {
		return "", err
	}
	// 获取结构体序列化之后的base64
	base64Str := base64.StdEncoding.EncodeToString(bytes)

	// 获取base64加密串
	sign, err := getTokenSign(base64Str, jwt_key)
	if err != nil {
		return "", err
	}

	// 返回token
	return base64Str + "." + sign, nil
}

// 生成token加密串 tokenBase64 -> md5 -> 加密串
func getTokenSign(tokenBase64 string, signKey string) (string, error) {
	if signKey == "" {
		return "", errors.New("signKey is empty")
	}
	md5Str := cryptor.Md5String(tokenBase64)
	encrypted := cryptor.AesEcbEncrypt([]byte(md5Str), []byte(signKey))
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func ParseToken(tokenString, jwt_key string) (*TokenStruct, error) {
	strs := strings.Split(tokenString, ".")
	if len(strs) != 2 {
		return nil, errors.New("token format error")
	}
	sign, err := getTokenSign(strs[0], jwt_key)
	if err != nil {
		return nil, err
	}
	if sign != strs[1] {
		return nil, errors.New("token sign error")
	}
	var tokenStruct TokenStruct
	decoded, err := base64.StdEncoding.DecodeString(strs[0])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(decoded, &tokenStruct)
	if err != nil {
		return nil, err
	}
	return &tokenStruct, nil
}
