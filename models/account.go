package models

import (
	"errors"
	"reflect"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
	"time"
)

type Account struct {
	Id       int64  `orm:"auto"`
	Owner    string `orm:"size(128)"`
	Account  string `orm:"size(128)"`
	Password string `orm:"size(128)"`
}

type OwnerAndAccount struct {
	Owner   string
	Account string
}

var (
	DBfile = "./data.db"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "file:"+DBfile+"?cache=shared&mode=rwc&_busy_timeout=10000")
	orm.RegisterModel(new(Account))
}

// AddAccount insert a new Account into database and returns
// last inserted Id on success.
func AddAccount(m *Account) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAccountById retrieves Account by Id. Returns error if
// Id doesn't exist
func GetAccountById(id int64) (v *Account, err error) {
	o := orm.NewOrm()
	v = &Account{Id: id}
	if err = o.QueryTable(new(Account)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOwnerAndAccount
func GetAllOwnerAndAccount() (oaa []OwnerAndAccount, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Account))
	var l []Account
	_, err = qs.All(&l)
	if err == nil {
		for _, v := range l {
			oaa = append(oaa, OwnerAndAccount{Owner: v.Owner, Account: v.Account})
		}
		return oaa, nil
	}
	return nil, err
}

// GetAllAccount retrieves all Account matches certain condition. Returns empty list if
// no records exist
func GetAllAccount(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Account))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Account
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateAccount updates Account by Id and returns error if
// the record to be updated doesn't exist
func UpdateAccountById(m *Account) (err error) {
	o := orm.NewOrm()
	v := Account{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			logs.Info("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAccount deletes Account by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAccount(account string) (err error) {
	o := orm.NewOrm()
	v := Account{Account: account}
	// ascertain account exists in the database
	if err = o.Read(&v, "Account"); err == nil {
		var num int64
		if num, err = o.Delete(&v); err == nil {
			logs.Info("Number of records deleted in database:", num)
		}
	}
	return
}

// Get the Count of Account
func GetAccountCount() (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Account))
	count, err = qs.Count()
	return
}

// Get Random Account for log in
func GetRandomAccount() (v *Account, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Account))
	count, err := qs.Count()
	if err != nil {
		return nil, err
	}
	rand.Seed(time.Now().Unix())
	index := rand.Intn(int(count))
	v = &Account{}
	err = qs.Limit(1, int64(index)).One(v)
	return
}
