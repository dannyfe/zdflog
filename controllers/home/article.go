package home

import "github.com/astaxie/beego"

type ArticleController struct {
	beego.Controller
}

func (self *ArticleController) List() {
	self.TplName = "home/list.html"
}
