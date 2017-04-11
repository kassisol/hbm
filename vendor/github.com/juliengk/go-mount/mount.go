package mount

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/juliengk/go-utils"
)

type Entry struct {
	Device        string
	MountPoint    string
	FSType        string
	Options       []string
	DumpFrequency int
	PassNumber    int
}

type Entries []Entry

var reEntry = regexp.MustCompile(`^(.+)\s+(.+)\s+(.+)\s+(.*)\s+([0-9]+)\s+([0-9]+)$`)

func New() (Entries, error) {
	mountFile := "/etc/mtab"
	entries := Entries{}

	file, err := os.Open(mountFile)
	if err != nil {
		return entries, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m := reEntry.FindStringSubmatch(scanner.Text())

		df, _ := strconv.Atoi(m[5])
		pn, _ := strconv.Atoi(m[6])

		e := Entry{
			Device:        m[1],
			MountPoint:    m[2],
			FSType:        m[3],
			Options:       strings.Split(m[4], ","),
			DumpFrequency: df,
			PassNumber:    pn,
		}

		entries = append(entries, e)
	}

	return entries, nil
}

func (entries Entries) Find(mountpoint string) (Entry, error) {
	for _, e := range entries {
		if e.MountPoint == mountpoint {
			return e, nil
		}
	}

	return Entry{}, fmt.Errorf("Mountpoint %s is not found", mountpoint)
}

func (entry Entry) FindOption(opt string) bool {
	if utils.StringInSlice(opt, entry.Options, false) {
		return true
	}

	return false
}
