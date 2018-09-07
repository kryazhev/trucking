package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/trucking/models"
)

func (c *AppController) Order() {
	email := c.GetString("email")
	phone := c.GetString("phone")
	message := c.GetString("message")

	err := models.SendEmail(email, "gadelshinaelena1985@gmail.com", "Заказ от "+phone, message)

	if err == nil {
		c.ajaxResponseSuccess(nil)
	} else {
		beego.Error("Can`t sent email ", err)
		c.ajaxResponseFail(err.Error())
	}
}
