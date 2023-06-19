package define

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	"sync"
)

var (
	Username     = "root"
	Host         = "mysql"
	Port         = "3306"
	Password     = "123456"
	ProDB        = "product_db"
	DevDB        = "device_db"
	UserDB       = "user_db"
	OrderDB      = "order_db"
	PayDB        = "pay_db"
	EmqxAddr     = "http://emqx:18083/api/v5"
	EmqxKey      = "f0bf2fffb190abf3"
	EmqxSec      = "da0ATny9ARy7kWjySxpiKF29CLgtJsPcIovayJNwH9CgCM"
	DevIdsCache  = "device_ids_cache"
	ProCache     = "product_categories_cache"
	ProIdsCache  = "product_ids_cache"
	UserIdsCache = "user_ids_cache"
)

var MysqlDsn string = Username + ":" + Password + "@tcp(" + Host + ":" + Port + ")/"

var ClientMap sync.Map

type KData struct {
	DeviceKey string `json:"device_key"`
	Payload   string `json:"payload"`
}

type KBData struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}

type UserClaim struct {
	IsAdmin  uint   `json:"is_admin"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Id       uint   `json:"userid"`
	jwt.RegisteredClaims
}

type Client struct {
	Conn *websocket.Conn
	Uid  string
}

var JwtKey = "iot-server"
