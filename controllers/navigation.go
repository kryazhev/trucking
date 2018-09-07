package controllers

/* Base */
func (c *AppController) SiteMap() {
	c.Ctx.Output.Header("Content-Type", "application/xml")
	c.TplName = "sitemap.tpl"
}

func (c *AppController) Get() {
	c.initData("home")
}
