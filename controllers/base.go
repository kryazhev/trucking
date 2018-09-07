package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/trucking/models"
	"net/http"
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
	/*	if c.Ctx.Request.TLS == nil {
			url := redirectURL(c.Ctx.Request)
			c.Redirect(url, 301)
		}
	*/if c.Data["Lang"] == nil {
		lang := c.Ctx.GetCookie("lang")
		if len(lang) == 0 {
			lang = "ru"
		}
		c.Data["Lang"] = lang
	}
}

func redirectURL(r *http.Request) string {
	if r.URL.IsAbs() {
		return strings.Replace(r.URL.RawQuery, "http", "https", 1)
	} else {
		return "https://" + host(r) + "/home.html"
	}
}

func host(r *http.Request) string {
	if len(r.Host) > 0 {
		return r.Host
	} else {
		return "localhost"
	}
}
