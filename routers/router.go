package routers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/trucking/controllers"
)

func init() {
	beego.Router("/sitemap.xml", &controllers.BaseController{}, "get:SiteMap")

	beego.Router("/", &controllers.AppController{})
	beego.Router("/home.html", &controllers.AppController{}, "get:Get")

	/* actions */
	beego.Router("/action/order", &controllers.AppController{}, "post:Order")

}
