/**
 * Pengo Project
 *
 * Copyright (c) 2015 Penlook Development Team
 *
 * --------------------------------------------------------------------
 *
 * This program is free software: you can redistribute it and/or
 * modify it under the terms of the GNU Affero General Public License
 * as published by the Free Software Foundation, either version 3
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public
 * License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 *
 * --------------------------------------------------------------------
 *
 * Author:
 *     Loi Nguyen       <loint@penlook.com>
 */
package model

import (
 	"fmt"
)

// Middleware
type Table struct {
	Name      string
	Server    string
	Port      int
	Database  string
	File 	  string
	Charset   string
	Username  string
	Password  string
	Connected bool
}

func (table Table) Connect() string {
	fmt.Println("Connecting to Sqlite ...")
	return "MySQL Connection"
}

func (table Table) ByTable(name string, schema interface {}) Table {
	return table
}

func (table Table) Create() bool {
	fmt.Println("Create new person")
	return true
}

func (table Table) First() string {
	return "abc"
}
