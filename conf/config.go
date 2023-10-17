package conf

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 全局config实例对象
// 也就是外面程序，在内存中的配置对象
// 程序内部获取配置，都通过读取该对象
// 该Config对象 什么时候被初始化呢？
//
//		配置加载时：
//	 		LoadConfigFromToml
//			LoadConfigFromEnv
//
// 为了不被程序在运行时恶意修改，设置成私有变量
var config *Config

// 全局MySQL  客户端实例
var db *sql.DB

// 要想获取配置，单独提供函数
// 全局Config对象获取函数
func C() *Config {
	return config
}

// 初始化一个有默认值的Config对象
func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		Log:   NewDefaultLog(),
		MySQL: NewDefaultMySQL(),
	}
}

// Config 应用配置
// 通过封装为一个对象，来与外部配置进行对接
type Config struct {
	App   *App   `toml:"app"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}

func NewDefaultApp() *App {
	return &App{
		Name: "demo",
		Host: "127.0.0.1",
		Port: "8050",
	}
}

type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
	Key  string `toml:"key" env:"APP_KEY"`
	// EnableSSL bool   `toml:"enable_ssl" env:"APP_ENABLE_SSL"`
	// CertFile  string `toml:"cert_file" env:"APP_CERT_FILE"`
	// KeyFile   string `toml:"key_file" env:"APP_KEY_FILE"`
}

func NewDefaultMySQL() *MySQL {
	return &MySQL{
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "root",
		Password:    "root",
		Database:    "demo",
		MaxOpenConn: 200,
		MaxIdleConn: 100,
	}
}

// MySQL TODO
type MySQL struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`

	// 因为使用的MySQL连接池，需要池做一些规划配置
	//	控制当前程序的MySQL打开的连接数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`

	// 控制MySQL复用， 比如5，最多运行5个来复用
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`

	// 一个连接的生命周期，这个和MySQL Server配置有关系，必须小于Server配置
	// 一个连接用12h 换一个conn，保证一定的可用性
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`

	// Idle 连接最多允许存活多久
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_idle_TIME"`

	// 作为私有变量，控制GetDB
	lock sync.Mutex
}

// 1.第一种方法，使用LoadGlobal，在加载时，初始化全局db实例
// 2.第二种方法，惰性加载  获取DB时，动态判断再初始化
func (m *MySQL) GetDB() *sql.DB {
	// 直接加锁，锁住临界区
	m.lock.Lock()
	defer m.lock.Unlock()

	// 如果实例不存在，会报错
	if db == nil {
		// 为了防止报错，db为空我们就  初始化一个全局db
		conn, err := m.getDBConn()
		if err != nil {
			panic(err)
		}
		db = conn
	}

	// 到这里 db一定存在
	return db
}

// 连接池，driverConn具体的连接对象，他维护着一个Socket
// pool []* driverConn，维护pool里面的连接都是可用的，定期检查我们的conn健康情况
// 某一个driverConn 已经失效，driverConn.Reset()，清空该结构体的数据，Reconn获取连接，让该conn借壳托活
// 避免driverConn结构体的内存申请和释放的一个成本 （为了防止并发安全还要加锁）
func (m *MySQL) getDBConn() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true", m.UserName, m.Password, m.Host, m.Port, m.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error , %s", dsn, err.Error())
	}
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil
}

func NewDefaultLog() *Log {
	return &Log{
		// debug, info,error,warn
		Level:  "info",
		Format: TextFormat,
		To:     ToStdout,
	}
}

// Log todo
// 用于配置全局Logger对象
type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}
