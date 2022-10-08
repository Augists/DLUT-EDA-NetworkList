package controllers

import (
	// "NetworkList/models"
	beego "github.com/beego/beego/v2/server/web"
)

type AddController struct {
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
func (c *AddController) Get() {
	c.Layout = "layout.tpl"
	c.TplName = "add.tpl"
}
