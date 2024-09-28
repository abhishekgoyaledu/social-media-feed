package database

type MysqlMasterSlave struct {
	Master          *MysqlConfig `json:"master"`
	Slave           *MysqlConfig `json:"slave"`
	EnableSlaveRead bool         `json:"enable_slave_read"`
}

type MysqlConfig struct {
	Dsn                    string `json:"dsn"`
	Host                   string `json:"host"`
	DatabaseName           string `json:"database_name"`
	Region                 string `json:"region"`
	MaxIdleConn            int    `json:"max_idle_conn"`
	MaxOpenConn            int    `json:"max_open_conn"`
	ConnMaxLifetimeInHours int    `json:"conn_max_lifetime_in_hours"`
	ReadTimeoutInSeconds   int    `json:"read_timeout_in_seconds"`
	WriteTimeoutInSeconds  int    `json:"write_timeout_in_seconds"`
	DialTimeoutInSeconds   int    `json:"dial_timeout_in_seconds"`
	EnableLog              bool   `json:"enable_log"`
}
