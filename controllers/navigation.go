package controllers

/* Base */
func (c *AppController) Get() {
	c.initData("home")
}

func (c *AppController) HomeNew() {
	c.initData("home-new")
}
