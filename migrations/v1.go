package migrations

var migartion_v1 = []string{
	`CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(40) PRIMARY KEY,
		first_name VARCHAR(40),
		last_name VARCHAR(40),
		is_admin BOOL NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(100) NOT NULL
	)`,
	`CREATE TABLE IF NOT EXISTS doors (
		id VARCHAR(40) PRIMARY KEY,
		acme_device_id VARCHAR(20),
		name VARCHAR(40)
	)`,
	`CREATE TABLE IF NOT EXISTS permissions (
		id VARCHAR(40) PRIMARY KEY,
		user_id VARCHAR(40) REFERENCES users(id) ON DELETE CASCADE,
		door_id VARCHAR(40) REFERENCES doors(id) ON DELETE CASCADE
	)`,
}
