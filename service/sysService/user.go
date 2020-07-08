package sysService

import (
	"gokins/comm"
	"gokins/models"
)

func FindUser(xid string) *models.SysUser {
	if xid == "" {
		return nil
	}
	e := new(models.SysUser)
	ok, err := comm.Db.Where("xid=?", xid).Get(e)
	if err != nil {
		return nil
	}
	if ok {
		return e
	}
	return nil
}