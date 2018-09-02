package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/oauth2"
	"regexp"
)

var languages = []string{"ru", "us"}

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
