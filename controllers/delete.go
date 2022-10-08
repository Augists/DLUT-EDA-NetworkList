package controllers

import (
	"NetworkList/models"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type DeleteController struct {
	beego.Controller
}

// Add ...
// @Title Add
// @Description add a new account
// @Param	owner		query 	string	true		"owner"
// @Param	account		query 	string	true		"account"
// @Param	password		query 	string	true		"password"
// @Success 200 {string} add success!
// @Failure 403 body is empty
// @router /add [get]
func (c *DeleteController) Get() {
	c.Data["get"] = "get"
	c.Layout = "layout.tpl"
	c.TplName = "delete.tpl"
}

func (c *DeleteController) Post() {
	account := c.GetString("account")
	if err := models.DeleteAccount(account); err != nil {
		c.Data["err"] = err
		logs.Info("Error:", err)
	} else {
		c.Data["delete"] = "Delete Account " + account + " Success"
		logs.Info("Delete Account", account, "Success")
	}
	c.Layout = "layout.tpl"
	c.TplName = "delete.tpl"
}
