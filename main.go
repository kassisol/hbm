// HBM is an application to authorize and manage authorized docker command using Docker AuthZ plugin.
// Copyright (C) 2016-2017 Kassisol inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"github.com/juliengk/go-log"
	"github.com/juliengk/go-utils/user"
	"github.com/kassisol/hbm/cli/command"
)

func main() {
	u := user.New()
	l, _ := log.NewDriver("standard", nil)

	if !u.IsRoot() {
		l.Error("You must be root to run that command")
	}

	cmd := command.NewHBMCommand()
	cmd.Execute()
}
