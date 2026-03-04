package common

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	orderIDCounter int
	mutex          sync.Mutex
)

//func ReadConfig(cfg *types.Config, path string) error {
//
//	configPathFromEnv := os.Getenv("pre")
//
//	if configPathFromEnv != "" { // allow the location of the config file to be passed via env args
//		path = configPathFromEnv
//	}
//	var err error
//	switch cfg.Network {
//	case "pre":
//		err = yaml.Unmarshal([]byte(config.PreConfigYml), &cfg)
//	case "prd":
//		err = yaml.Unmarshal([]byte(config.PrdConfigYml), &cfg)
//	default:
//		err = yaml.Unmarshal([]byte(config.TestConfigYml), &cfg)
//	}
//	if err != nil {
//		return err
//	}
//	return nil
//}

func CreateToken(user string, lastLoginTime int64) (string, error) {
	claim := jwt.MapClaims{
		"user":            user,
		"sign":            fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%s%s", user, viper.GetString("accountKey"))))),
		"nbf":             time.Now().Unix(),
		"iat":             time.Now().Unix(),
		"last_login_time": lastLoginTime,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	token, err := tokenClaims.SignedString([]byte(viper.GetString("jwtKey")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(tokens string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt_key")), nil
	})
	if err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot convert claim to mapclaim")
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		return nil, err
	}
	return claim, nil
}

func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func RandString(length int) string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomNum := rand.Intn(10000)
	rand.Seed(time.Now().UnixNano() + int64(randomNum))
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomString[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(randomString)
}

func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(email)
}

func RandOrderId() string {
	mutex.Lock()
	defer mutex.Unlock()
	orderIDCounter++
	if orderIDCounter > 9999 {
		orderIDCounter = 0
	}
	counterStr := fmt.Sprintf("%04d", orderIDCounter)

	timestamp := time.Now().UnixNano()
	randomNum := rand.Intn(10000)
	randomStr := fmt.Sprintf("%04d", randomNum)

	orderID := strconv.FormatInt(timestamp, 10) + counterStr + randomStr
	return orderID
}

func ReplaceMiddleWithAsterisks(str string) string {
	length := len(str)
	if length <= 3 {
		return str
	}
	replacement := strings.Repeat("*", length-3)
	result := str[:2] + replacement + str[length-1:]
	return result
}
func CountPrecision(s string) int {
	if s == "" || !strings.Contains(s, ".") {
		return 0
	}
	parts := strings.SplitN(s, ".", 2)
	if len(parts) < 2 {
		return 0
	}
	return len(parts[1])
}
