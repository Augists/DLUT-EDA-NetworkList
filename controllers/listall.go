package controllers

import (
	"NetworkList/models"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// ListController operations for List
type ListAllController struct {
	beego.Controller
}

// @Title ListAll
// @Description get List of all owner and their account
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.List
// @Failure 403
// @router / [get]
func (c *ListAllController) Get() {
	owneraccount, err := models.GetAllOwnerAndAccount()
	if err != nil {
		c.Data["err"] = err
		logs.Info(err)
	} else {
		c.Data["json"] = owneraccount
		logs.Info(owneraccount)
	}
	c.Layout = "layout.tpl"
	c.TplName = "listall.tpl"
}
