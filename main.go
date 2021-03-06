package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/kryazhev/trucking/models"
	_ "github.com/kryazhev/trucking/routers"
	"os"
	"strings"
)

var envNames = []string{"PORT", "run-mode", "enable-http", "enable-https", "smtp.address", "smtp.host", "smtp.user", "smtp.password"}

func main() {
	// Display list of environment variables
	for _, env := range envNames {
		beego.Trace(env + "=" + os.Getenv(env))
	}

	// Initialize language type list.
	langTypes := strings.Split(beego.AppConfig.String("lang::types"), "|")

	// Load locale files according to language types.
	for _, lang := range langTypes {
		beego.Trace("Loading language:", lang)
		if err := i18n.SetMessage(lang, "static/i18n/locale_"+lang+".ini"); err != nil {
			panic(err)
		}
	}

	beego.AddFuncMap("i18n", i18n.Tr)
	beego.AddFuncMap("data", models.Data)
	beego.AddFuncMap("lookupEnv", models.LookupEnv)
	beego.Run()
}
