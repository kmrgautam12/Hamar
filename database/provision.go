package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type dbCred struct {
	db *sql.DB
}

func CheckProvisioning() *dbCred {
	connStr := "host=localhost port=5432 user=test password=AdminSex dbname=mydatabase sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot ping database:", err)
	}

	fmt.Println("Connected to DB!")
	d := &dbCred{
		db: db,
	}
	return d
}

func (d *dbCred) DBProvisioningPipeline() error {
	err := d.checkRolesTable()
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = d.checkPermissionTable()
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = d.assignPermissionToRoles()
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = d.CreateUserTable()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (d *dbCred) checkRolesTable() error {
	// check role table
	schema := `CREATE TABLE IF NOT EXISTS roles (
        id SERIAL PRIMARY KEY,
        name VARCHAR(50) UNIQUE NOT NULL
    )`
	_, err := d.db.Exec(schema)
	if err != nil {
		log.Fatal(err)
		return err
	}
	roles := []string{"user", "admin", "owner"}
	for _, role := range roles {
		_, err = d.db.Exec(`INSERT INTO roles (name) VALUES ($1) ON CONFLICT (name) DO NOTHING`, role)
		if err != nil {
			return fmt.Errorf("insertRoles: %w", err)
		}
	}
	return nil

}

func (d *dbCred) checkPermissionTable() error {

	// Create Permission Table If Not Exist

	schema := `CREATE TABLE IF NOT EXISTS permissions (
        id SERIAL PRIMARY KEY,
        name VARCHAR(50) UNIQUE NOT NULL
    )`
	_, err := d.db.Exec(schema)
	if err != nil {
		log.Fatal(err)
		return err
	}
	permissions := []string{"read", "write", "delete"}
	for _, perm := range permissions {
		_, err := d.db.Exec(`INSERT INTO permissions (name) VALUES ($1) ON CONFLICT (name) DO NOTHING`, perm)
		if err != nil {
			return fmt.Errorf("insertPermissions: %w", err)
		}
	}
	return nil
}

func (d *dbCred) assignPermissionToRoles() error {
	// Create role permission table

	schema := `CREATE TABLE IF NOT EXISTS role_permissions (
        role_id INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
        permission_id INT NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
        PRIMARY KEY (role_id, permission_id))`

	_, err := d.db.Exec(schema)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func (d *dbCred) CreateUserTable() error {
	schema := `CREATE TABLE users (
    user_id         UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- or SERIAL if UUID not supported
    username        VARCHAR(150) UNIQUE NOT NULL,
    password_hash   TEXT NOT NULL,
    access_level    VARCHAR(50) NOT NULL CHECK (access_level IN ('admin', 'user', 'owner')),
    email           VARCHAR(255) UNIQUE,
    is_active       BOOLEAN NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
	)`
	_, err := d.db.Exec(schema)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}
