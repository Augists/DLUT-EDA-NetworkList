package controllers

import (
	"NetworkList/models"
	"NetworkList/utils"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// ListController operations for List
type ListController struct {
	beego.Controller
}

// @Title Get
// @Description get a random account and login
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.List
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ListController) Get() {
	account, err := models.GetRandomAccount()
	utils.Login(account)
	// ac := models.OwnerAndAccount{Owner: account.Owner, Account: account.Account}
	if err != nil {
		c.Data["err"] = err.Error()
		logs.Info(err.Error())
	} else {
		c.Data["json"] = account
		logs.Info("Log in using", account.Owner, account.Account)
	}
	c.Layout = "layout.tpl"
	c.TplName = "list.tpl"
}
