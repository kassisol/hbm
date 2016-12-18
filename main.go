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
	"log"
	"os"
	"os/user"

	"github.com/kassisol/hbm/cmd"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)

		os.Exit(-1)
	}

	if user.Uid != "0" {
		fmt.Println("You must be root to run that command")

		os.Exit(-1)
	}

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)

		os.Exit(-1)
	}
}
