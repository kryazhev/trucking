package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/oauth2"
	"github.com/kryazhev/trucking/models"
	"regexp"
)

var languages = []string{"ru"}

/* OAuth 2.0 */
func (c *AppController) Callback() {
	state := c.GetString("state")

	var user *oauth2.User
	var err error
	m := regexp.MustCompile(`^([a-z]+)`).FindStringSubmatch(state)
	if len(m) > 1 {
		endpointName := m[1]
		code := c.GetString("code")

		config := oauth2.AuthConfigs[endpointName]

		user, err = config.GetUser(endpointName, code)

		if err != nil {
			beego.Error(err)
		} else {
			session := c.StartSession()
			session.Set("user", user)
		}
	}

	page := c.GetString("page")
	c.initDataWithUser(page, user)
}

func (c *AppController) Logout() {
	beego.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)

	c.initDataWithUser(c.GetString("page"), nil)
}

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
