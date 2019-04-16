package db_test

import (
	"testing"

	"github.com/qinhan-shu/gp-server/model/xorm"
)

func TestMysqlDriver_GetAllLanguage(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	languages, err := mysqlDriver.GetAllLanguage()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(len(languages))
	for _, d := range languages {
		t.Logf("%+v\n", d)
	}
}
func TestMysqlDriver_AddLanguage(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}

	language := &model.Language{
		Detail: "xxx",
	}

	if err := mysqlDriver.AddLanguage(language); err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", language)
}

func TestAddSomeLanguages(t *testing.T) {
	mysqlDriver, err := getMysqlDriver()
	if err != nil {
		t.Error(err)
		return
	}
	var d = []string{
		"C",
		"C_PLUS",
		"JAVA",
		"PYTHON2",
		"PYTHON3",
	}

	for _, v := range d {
		if err := mysqlDriver.AddLanguage(
			&model.Language{
				Detail: v,
			}); err != nil {
			t.Error(err)
			return
		}
	}
}
