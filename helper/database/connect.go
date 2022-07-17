package database

import "fmt"

type MySQLService interface {
	ConnectionString() string
}

type Connection struct {
	User     string
	Password string
	Host     string
	Port     string
	DB       string
}

func Connect() MySQLService {
	return &Connection{
		User:     "app_user",
		Password: "#XRvxRXT2v9aTg7LC7yZ",
		Host:     "127.0.0.1",
		Port:     "3306",
		DB:       "pien2022db",
	}

}

func (conn *Connection) ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", conn.User, conn.Password, conn.Host, conn.Port, conn.DB)
}
