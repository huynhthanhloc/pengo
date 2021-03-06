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
package library

import (
  	"github.com/codegangsta/cli"
  	"log"
    "os"
    "path/filepath"
    . "github.com/penlook/pengo/cmd/pengo/library/template"
)

type Run struct {
	Context *cli.Context
	Directory string
	Parser Parser
    Data Data
}

func (run *Run) GetCurrentDirectory(appName string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]) + "/" + appName)
    if err != nil {
        log.Fatal(err)
    }
    run.Directory = dir
}

func (run *Run) ParseApplication() {
    parser := Parser {}
    run.Data = Data {}

    // Parse controller
    dir, err := filepath.Abs(run.Directory + "/generate/controller")
    if err != nil {
        panic("Folder name 'controller' does not exist !")
    }
    run.Data["controller"] = parser.Controller(dir)

    // Parse model
    dir, err = filepath.Abs(run.Directory + "/model")
    if err != nil {
        panic("Folder name 'model' does not exist !")
    }
    run.Data["model"] = parser.Model(dir)

    // Parse extend
    dir, err = filepath.Abs(run.Directory + "/extend")
    if err != nil {
        panic("Folder name 'extend' does not exist !")
    }
    run.Data["extend"] = parser.Extend(dir)
}

func (run *Run) Generate() {
    generator := Generator {}

    main_path, _ := filepath.Abs(run.Directory + "/build/main.go")
    generator.Main(TemplateMain, main_path, run.Data)

    schema_path, _ := filepath.Abs(run.Directory + "/generate/controller/schema.go")
    generator.Schema(TemplateSchema, schema_path, run.Data)
}

func (run Run) Compile() {
    path, _ := filepath.Abs(run.Directory)
    compiler := Compiler {
        Directory : path,
    }
    compiler.ParseController()
}

func (run *Run) Run() {
	run.GetCurrentDirectory(run.Context.Args().First())
    run.Compile()
    run.ParseApplication()
    run.Generate()
}
