package trash

import (
	"github.com/rkoesters/xdg/basedir"
	"path/filepath"
)

var hometrash *Dir

func init() {
	var err error
	hometrash, err = New(filepath.Join(basedir.DataHome, "Trash"))
	if err != nil {
		// TODO: remove this
		panic(err)
	}
}

// Files returns a slice of the files in the default trash.
func Files() ([]string, error) { return hometrash.Files() }

// Stat returns the Info for the given file in the trash.
func Stat(s string) (*Info, error) { return hometrash.Stat(s) }

// Erase removes the given file from the trash.
func Erase(s string) error { return hometrash.Erase(s) }

// EraseAll removes the given file and all children from the trash.
func EraseAll(s string) error { return hometrash.EraseAll(s) }

// Empty erases all files in the trash.
func Empty() error { return hometrash.Empty() }

// IsEmpty returns whether or not the trash is empty.
func IsEmpty() bool { return hometrash.IsEmpty() }
