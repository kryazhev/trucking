package controllers

/* Base */
func (c *AppController) Get() {
	c.initData("home")
}

func (c *AppController) SiteMap() {
	c.Ctx.Output.Header("Content-Type", "application/xml")
	c.Ctx.Output.ContentType("xml")
	c.TplName = "sitemap.tpl"
}
