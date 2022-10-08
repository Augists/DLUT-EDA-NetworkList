package controllers

import (
	"NetworkList/models"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Layout = "layout.tpl"
	c.TplName = "index.tpl"
}

// Post ...
// @Title Create
// @Description create List
// @Param	body		body 	models.List	true		"body for List content"
// @Success 201 {object} models.List
// @Failure 403 body is empty
// @router / [post]
func (c *MainController) Post() {
	c.Layout = "layout.tpl"
	c.TplName = "index.tpl"
	account := new(models.Account)
	account.Owner = c.GetString("owner")
	account.Account = c.GetString("account")
	account.Password = c.GetString("password")
	logs.Info("Add Owner:", account.Owner)
	id, err := models.AddAccount(account)
	if err != nil {
		logs.Info("Add Account Error:", err)
		c.Data["err"] = err
	} else {
		c.Data["id"] = id
	}
}
