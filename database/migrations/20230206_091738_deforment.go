package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Deforment_20230206_091738 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Deforment_20230206_091738{}
	m.Created = "20230206_091738"

	migration.Register("Deforment_20230206_091738", m)
}

// Run the migrations
func (m *Deforment_20230206_091738) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE deforment(`id` int(11) NOT NULL AUTO_INCREMENT,`dname` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Deforment_20230206_091738) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `deforment`")
}
