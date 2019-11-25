package mysql

import (
	"strings"
	"testing"
)

func TestReadConfig(t *testing.T) {
	r := strings.NewReader(`
development:
  dialect: mysql
  datasource: root:password@tcp(127.0.0.1:3306)/test?parseTime=true
  dir: migrations/mysql`)

	configs, err := NewDBConfigs(r)
	if err != nil {
		t.Fatalf("read config failed: %s", err)
	}

	c, ok := configs["development"]
	if !ok {
		t.Fatalf("cannnot read development configration.")
	}

	if ex := "root:password@tcp(127.0.0.1:3306)/test?parseTime=true"; c.Datasource != ex {
		t.Fatalf("want %s, got %s", ex, c.Datasource)
	}
	if ex := "mysql"; c.Dialect != ex {
		t.Fatalf("want %s, got %s", ex, c.Dialect)
	}
}
