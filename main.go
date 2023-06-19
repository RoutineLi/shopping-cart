package main

import (
	"graduate_design/define"
	"graduate_design/models"
)

func main() {
	models.NewDB(define.ProDB)
	models.NewDB(define.DevDB)
	models.NewDB(define.UserDB)
	models.NewDB(define.PayDB)
	models.NewDB(define.OrderDB)
}
