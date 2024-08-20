package data

import (
	"context"
	"time"
)

const userTable = `create table if not exists users (
	id serial primary key,
	email varchar(255) not null unique,
	first_name varchar(100) not null,
	last_name varchar(100) not null,
	password varchar(100) not null,
	user_active int not null,
	created_at timestamp not null,
	updated_at timestamp not null
)`

const passwordTable = `create table if not exists passwords (
	id serial primary key,
	email varchar(255) not null,
	password varchar(100) null,
	code varchar(6) not null unique,
	used boolean not null,
	created_at timestamp not null,
	updated_at timestamp not null
)`

func CreateTable(tables ...string) {
	// query maps table names to their respective queries
	var tablesQuery = map[string]string{
		"users":     userTable,
		"passwords": passwordTable,
	}

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for _, name := range tables {
		query := tablesQuery[name]
		db.ExecContext(ctx, query)
	}
}
