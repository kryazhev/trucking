package routers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/trucking/controllers"
)

func init() {
	beego.Router("/", &controllers.AppController{})
	beego.Router("/home.html", &controllers.AppController{}, "get:Get")
	beego.Router("/example.html", &controllers.AppController{}, "get:Example")

	/* actions */
	beego.Router("/action/oauth2-callback", &controllers.AppController{}, "get:Callback")
	beego.Router("/action/logout", &controllers.AppController{}, "get:Logout")
}
