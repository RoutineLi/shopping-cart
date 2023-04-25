package define

import "github.com/golang-jwt/jwt/v4"

var (
	Username  = "root"
	Host      = "192.168.0.101"
	Port      = "3306"
	Password  = "123456"
	EmqxAddr  = "http://192.168.0.101:18083/api/v5"
	DevCache  = "device_cache"
	ProCache  = "product_cache"
	UserCache = "user_cache"
)

var MysqlDsn string = Username + ":" + Password + "@tcp(" + Host + ":" + Port + ")/"

type UserClaim struct {
	IsAdmin  uint   `json:"is_admin"`
	Password string `json:"password"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

var JwtKey = "iot-server"
