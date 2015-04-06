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
package controller

import (
	. "github.com/penlook/pengo"
	. "github.com/penlook/pengo/cmd/pengo/app/generate"
)

type App struct { Controller }

// @pick "Before action"
func (app App) Before() {

}

// @pick "After action"
func (app App) After() {

}

// @route /app/home
// @method GET
// @pick "Home action"
func (app App) Home() {

	app.View(Data{"title": "Index Page",})
	app.View(Data{"sample": "Welcome to application home",})
	app.View(Data{"image": "http://img3.wikia.nocookie.net/__cb20140410201208/pokemon/images/e/ef/025Pikachu_SSB4.png",})
	app.Pick("Before call pk")
	Pk()
	app.Pick("After call pk")

	app.Pick("Assign slogan and author")
}
