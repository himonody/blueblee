package config

import "strconv"

type Mysql struct {
	Host     string `yaml:"host"`     //服务器地址:端口
	Port     int    `yaml:"port"`     //端口
	Config   string `yaml:"config"`   //高级配置
	DB       string `yaml:"DB"`       //数据库名
	User     string `yaml:"user"`     //数据库用户名
	Password string `yaml:"password"` //数据库密码
	//MaxIdleConns int    `yaml:"max-idle-conns"` //空闲最大连接数
	//MaxOpenConns int    `yaml:"max-open-conns"` //打开数据库的最大连接数
	LogLevel string `yaml:"logLevel"`
}

func (m *Mysql) Dsn() string {
	return m.User + ":" +
		m.Password + "@tcp(" +
		m.Host + ":" +
		strconv.Itoa(m.Port) + ")/" +
		m.DB + "?" + m.Config

}
