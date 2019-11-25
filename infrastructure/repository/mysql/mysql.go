package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/lestrrat-go/test-mysqld"
	"github.com/rubenv/sql-migrate"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
)

type DBConfigs map[string]*DBConfig

type DBConfig struct {
	// RDBMS name (ex: mysql)
	Dialect    string `yaml:"dialect"`
	Datasource string `yaml:"datasource"`
	Dir        string `yaml:"dir"`
}

func (c *DBConfig) Open() (*gorm.DB, error) {
	conn, err := gorm.Open(c.Dialect, c.Datasource)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewDBConfigsFromFile(path string) (DBConfigs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewDBConfigs(f)
}

func NewDBConfigs(r io.Reader) (DBConfigs, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var configs DBConfigs
	if err = yaml.Unmarshal(b, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}

func (c *DBConfig) OpenMock(migrationFilePath string) (*sqlx.DB, error) {
	mysqld, err := mysqltest.NewMysqld(nil)
	if err != nil {
		return nil, err
	}
	ds := mysqld.Datasource("test", "", "", 0, mysqltest.WithParseTime(true))
	conn, err := sqlx.Open(c.Dialect, ds)
	if err != nil {
		return nil, err
	}

	migrations := &migrate.FileMigrationSource{
		Dir: migrationFilePath,
	}

	_, err = migrate.Exec(conn.DB, c.Dialect, migrations, migrate.Up)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
