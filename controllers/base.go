package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/trucking/models"
	"strings"
)

var pages = []string{"home"}

type BaseController struct {
	beego.Controller
}

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

func (c *AppController) ajaxResponseSuccess(data interface{}) {
	c.Data["json"] = models.Result{Success: true, Data: data}
	c.ServeJSON()
}

func (c *AppController) ajaxResponseFail(message string) {
	c.Data["json"] = models.Result{Success: false, Message: message}
	c.ServeJSON()
}

func (c *AppController) Prepare() {
	if c.Data["Lang"] == nil {
		lang := c.Ctx.GetCookie("lang")
		if len(lang) == 0 {
			lang = "ru"
		}
		c.Data["Lang"] = lang
	}
}
