package controllers

/* Base */
func (c *BaseController) SiteMap() {
	c.Ctx.Output.Header("Content-Type", "application/xml")
	c.Ctx.Output.ContentType("xml")
	c.TplName = "sitemap.tpl"
}

func (c *AppController) Get() {
	c.initData("home")
}
