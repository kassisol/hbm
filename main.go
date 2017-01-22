//    HBM is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.
//
//    You should have received a copy of the GNU General Public License
//    along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os/user"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cmd"
)

func main() {
	user, err := user.Current()
	if err != nil {
		utils.Exit(err)
	}

	if user.Uid != "0" {
		utils.Exit(fmt.Errorf("You must be root to run that command"))
	}

	if err := cmd.RootCmd.Execute(); err != nil {
		utils.Exit(err)
	}
}
