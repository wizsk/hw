package db

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Entry struct {
	Id          int64 `json:"id"`
	Pid         int64 `json:"pid"`
	IsRoot      bool  `json:"is_root"`
	IsHighlight bool  `json:"is_highlight"`
	// Word is also the root when IsRoot is true
	Word string `json:"word"`
	Def  string `json:"def"`
}

type Entries []Entry

// by id
func (e Entries) sort() {
	sort.Slice(e, func(i, j int) bool {
		return i < j
	})
}

var (
	ErrorNotFound = errors.New("Entry not found")
)

func RootSuggestion(root string, lim int) []string {
	root = strings.TrimSpace(root)
	if root == "" {
		return nil
	}

	// root lenght
	r := []rune(root)
	found := 0
	res := []string{}

loop:
	for i := range len(dict) {
		e := &dict[i]
		if dict[i].IsRoot && isSub([]rune(e.Word), r) {
			// only add uniqe words
			for _, v := range res {
				if v == e.Word {
					continue loop
				}
			}
			res = append(res, e.Word)
			found++
		}
		if found >= lim {
			break
		}
	}
	return res
}

// s = abc, sub = ab -> true
func isSub(s, sub []rune) bool {
	if len(s) < len(sub) {
		return false
	}
	for i, r := range sub {
		if s[i] != r {
			return false
		}
	}
	return true
}

// input is cleaned while calling func
func SearchByRoot(root string, lim int) (Entries, error) {
	root = strings.TrimSpace(root)

	found := 0
	res := Entries{}

	// parent id
	pid := int64(-1)
	for i := range len(dict) {
		if dict[i].IsRoot && root == dict[i].Word {
			pid = dict[i].Pid
		}
	}

	if pid < 0 {
		return nil, ErrorNotFound
	}

	for i := range len(dict) {
		e := &dict[i]
		if e.Pid == pid {
			res = append(res, *e)
			found++
		}

		if found >= lim {
			break
		}
	}

	res.sort()
	return res, nil
}

// fmt is the replaced text
// defaut: `<span style="background: yellow;">%s</span>`
// provide "" to use the default
//
// input is cleaned while calling func
func SearchByTxt(str string, lim int, format string) (Entries, error) {
	str = strings.TrimSpace(str)
	if format == "" {
		format = `<span style="background: yellow;">%s</span>`
	}
	found := 0
	res := Entries{}

	for i := range len(dict) {
		if strings.Contains(dict[i].Def, str) {
			e := dict[i] // copying
			e.Def = strings.ReplaceAll(e.Def, str, fmt.Sprintf(format, str))
			res = append(res, e)
			found++
		}
		if found >= lim {
			break
		}
	}

	if len(res) == 0 {
		return nil, ErrorNotFound
	}

	res.sort()
	return res, nil
}
