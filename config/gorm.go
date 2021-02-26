package config

// Mysql connection configuration
type Mysql struct {
	Path         string `json:"path" yaml:"path"`
	Config       string `json:"config" yaml:"config"`
	Dbname       string `json:"dbname" yaml:"db-name"`
	Username     string `json:"username" yaml:"username"`
	Password     string `json:"password" yaml:"password"`
	MaxIdleConns int    `json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `json:"maxOpenConns" yaml:"max-open-conns"`
}

// Dsn for establishing mysql connection
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}
