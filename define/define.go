package define

import "github.com/golang-jwt/jwt/v4"

var (
	Username     = "root"
	Host         = "mysql"
	Port         = "3306"
	Password     = "123456"
	EmqxAddr     = "http://emqx:18083/api/v5"
	DevIdsCache  = "device_ids_cache"
	ProCache     = "product_categories_cache"
	ProIdsCache  = "product_ids_cache"
	UserIdsCache = "user_ids_cache"
)

var MysqlDsn string = Username + ":" + Password + "@tcp(" + Host + ":" + Port + ")/"

type UserClaim struct {
	IsAdmin  uint   `json:"is_admin"`
	Password string `json:"password"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

var JwtKey = "iot-server"
