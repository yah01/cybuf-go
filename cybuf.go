package cybuf

import (
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

var (
	//debugLog *log.Logger
	// errorLog *log.Logger
)

func init() {
	//debugLog = log.New(os.Stdout, "Debug ", log.LstdFlags|log.Lshortfile)
	// errorLog = log.New(ioutil.Discard, "Error ", log.LstdFlags|log.Lshortfile)
}




func IsValidKeyName(name []byte) bool {

	name = []byte(strings.TrimSpace(string(name)))
	if len(name) == 0 {
		return false
	}

	if len(name) > 0 && !unicode.IsLetter(rune(name[0])) && !(name[0] != '_') {
		return false
	}
	for _, c := range name {
		if !unicode.IsLetter(rune(c)) && unicode.IsDigit(rune(c)) && c != '_' {
			return false
		}
	}
	return true
}
