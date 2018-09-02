package controllers

/* Base */
func (c *AppController) Get() {
	c.initData("home")
}

func (c *AppController) Example() {
	c.initData("example")
}
