package routers

import (
	"github.com/astaxie/beego"
	"zdflog/controllers/home"
)

func init() {
	beego.Router("/", &home.MainController{})

	//前台列表
	beego.Router("/list.html", &home.ArticleController{}, "get:List")

}
