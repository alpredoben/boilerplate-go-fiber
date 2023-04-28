package database

import (
	packagesConfig "boilerplate-go/src/packages/config"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/** Struct Database with Gorm **/
type GormDB struct {
	*gorm.DB
}

/** Instance Database with Gorm to variable */
var defaultDB = &GormDB{}

func GetGormDB() *GormDB {
	return defaultDB
}

/** Connect to database with gorm */
func (db *GormDB) connect(cfg *packagesConfig.DBEnv) (err error) {
	// postgreSQL URI environment
	postgresURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Jakarta",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.Database,
		cfg.Port,
		cfg.SslMode,
	)

	// open new database connection
	conn, err := sql.Open("postgres", postgresURI)

	if err != nil {
		log.Println("Connect to database failed...")
		return err
	}

	// set config connection
	conn.SetMaxOpenConns(cfg.MaxOpenConnection)
	conn.SetMaxIdleConns(cfg.MaxIdleConnection)
	conn.SetConnMaxLifetime(cfg.MaxLifeTimeConnection)

	sqlConn, err := conn.Conn(context.Background())

	if err != nil {
		defer conn.Close()
		return fmt.Errorf("can't get connection from pool, %w", err)
	}

	// test ping to database
	if err := sqlConn.PingContext(context.Background()); err != nil {
		defer sqlConn.Close()
		return fmt.Errorf("can't send ping to database, %w", err)
	}

	// create a new gorm.DB instance using the database connection
	db.DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlConn,
	}), &gorm.Config{})

	if err != nil {
		log.Println("Connect to database failed...")
		return err
	}

	log.Printf("Successfully gorm connected to database %s ...", cfg.Database)

	// Set logging mode
	db.DB.Logger = logger.Default.LogMode(logger.Info)

	return nil
}

func GormDatabaseConnection() error {
	return defaultDB.connect(packagesConfig.GetDatabaseConfig())
}
