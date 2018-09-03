package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/oauth2"
	"github.com/kryazhev/trucking/models"
	"strings"
)

var pages = []string{"home"}

type AppController struct {
	beego.Controller
}

func (c *AppController) initData(page string) {
	if !models.HasElem(pages, page) {
		page = "home"
	}

	c.Data["Page"] = page
	c.TplName = strings.Replace(page, ".", "/", 1) + ".html"
}

func (c *AppController) initDataWithUser(page string, user *oauth2.User) {
	if user != nil {
		c.Data["User"] = user
	} else {
		delete(c.Data, "User")
	}

	c.initData(page)
}

func (c *AppController) ajaxResponseSuccess(data interface{}) {
	c.Data["json"] = models.Result{Success: true, Data: data}
	c.ServeJSON()
}

func (c *AppController) ajaxResponseFail(message string) {
	c.Data["json"] = models.Result{Success: false, Message: message}
	c.ServeJSON()
}

func (c *AppController) Prepare() {
	c.Data["OAuthConfigs"] = oauth2.AuthConfigs

	if beego.GlobalSessions != nil {
		user := c.GetSession("user")
		if user != nil {
			c.Data["User"] = user
		}
	}

	if c.Data["Lang"] == nil {
		lang := c.Ctx.GetCookie("lang")
		if len(lang) == 0 {
			lang = "ru"
		}
		c.Data["Lang"] = lang
	}
}
