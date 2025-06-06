package conf

import "fmt"

type DB struct {
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	DB          string `yaml:"db"`
	Debug       bool   `yaml:"debug"`  //输出全部日志
	Source      string `yaml:"source"` //mysql或pgsql
	MaxConn     int    `yaml:"max_conn"`
	MaxIdle     int    `yaml:"max_idle"`
	MaxLifeTime int    `yaml:"max_lifetime"`
}

func (db *DB) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.User, db.Password, db.Host, db.Port, db.DB)
}

func (db *DB) Empty() bool {
	return db.User == "" && db.Password == "" && db.Host == "" && db.Port == 0
}

func (db DB) Addr() string {
	return fmt.Sprintf("%s:%d", db.Host, db.Port)
}
