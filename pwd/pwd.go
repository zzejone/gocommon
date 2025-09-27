package pwd

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

// ------------------- 参数 -------------------
// 下面的参数是常用的安全配置，实际项目可根据硬件性能调节
const (
	argonTime    uint32 = 2         // 迭代次数
	argonMemory  uint32 = 64 * 1024 // 64 MiB
	argonThreads uint8  = 4         // 并行度
	argonKeyLen  uint32 = 32        // 输出长度（bytes）
	saltLen      int    = 16        // 128‑bit 随机盐
)

// ------------------- 生成随机盐 -------------------
func generateSalt() ([]byte, error) {
	salt := make([]byte, saltLen)
	_, err := rand.Read(salt)
	return salt, err
}

// ------------------- 哈希生成 -------------------
func HashPassword(password string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	// 计算 Argon2id 哈希
	hash := argon2.IDKey([]byte(password), salt, argonTime, argonMemory, argonThreads, argonKeyLen)

	// 使用 URL‑safe Base64（不带 padding）便于存库
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// 统一的编码格式
	encoded := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		argonMemory, argonTime, argonThreads, b64Salt, b64Hash)

	return encoded, nil
}

// ------------------- 哈希校验 -------------------
func VerifyPassword(password, encoded string) (bool, error) {
	// 1. 按 $ 分割
	parts := strings.Split(encoded, "$")
	if len(parts) != 6 {
		return false, fmt.Errorf("invalid encoded argon2 hash format")
	}
	// parts[0]为空字符串，parts[1]="argon2id", parts[2]="v=19",
	// parts[3]="m=...,t=...,p=...", parts[4]=salt, parts[5]=hash

	// 2. 解析参数 m、t、p
	paramStr := parts[3] // "m=65536,t=2,p=4"
	paramMap := map[string]uint32{}
	for _, kv := range strings.Split(paramStr, ",") {
		kvParts := strings.SplitN(kv, "=", 2)
		if len(kvParts) != 2 {
			return false, fmt.Errorf("invalid parameter segment: %s", kv)
		}
		val, err := strconv.ParseUint(kvParts[1], 10, 32)
		if err != nil {
			return false, fmt.Errorf("invalid numeric value in %s: %v", kv, err)
		}
		paramMap[kvParts[0]] = uint32(val)
	}
	memory := paramMap["m"]
	timeCost := paramMap["t"]
	threads := uint8(paramMap["p"])

	// 3. 解码盐和哈希（Base64 URL‑safe，无 padding）
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, fmt.Errorf("invalid base64 salt: %v", err)
	}
	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, fmt.Errorf("invalid base64 hash: %v", err)
	}

	// 4. 重新计算哈希
	calculated := argon2.IDKey([]byte(password), salt, timeCost, memory, threads, uint32(len(expectedHash)))

	// 5. 常量时间比较
	match := subtle.ConstantTimeCompare(calculated, expectedHash) == 1
	return match, nil
}
