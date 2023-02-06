package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Users_20230206_084348 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20230206_084348{}
	m.Created = "20230206_084348"

	migration.Register("Users_20230206_084348", m)
}

// Run the migrations
func (m *Users_20230206_084348) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE users(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`tel` varchar(128) NOT NULL,`d_id` int(11) DEFAULT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Users_20230206_084348) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `users`")
}
