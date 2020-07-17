package sysytem_service

import "go-gin-starter/models"

func Execute(script string) error {
	session := models.NewSession()
	s := &models.Script{
		Sql: script,
	}
	err := models.NewScriptModel(session).Execute(s)
	return err
}
