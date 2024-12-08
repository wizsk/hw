// converts the sql db to go struct
package main

/*

import (
	"bytes"
	"fmt"
	"os"

	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

type Entry struct {
	Id          int64  `json:"id"`
	Pid         int64  `json:"pid"`
	IsRoot      bool   `json:"is_root"`
	IsHighlight bool   `json:"is_highlight"`
	Word        string `json:"word"`
	Def         string `json:"def"`
}

type Entries []Entry

func main() {
	conn, err := sqlite.OpenConn("assets/hw.db", sqlite.OpenReadOnly)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	e := Entries{}

	const q = `SELECT * FROM dict` // limit 100`

	err = sqlitex.Execute(conn, q, &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			e = append(e, getRowData(stmt))
			return nil
		},
	})

	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)

	buf.WriteString("package db\n\n")

	buf.WriteString("var dict Entries = []Entry{\n")
	ln := len(e)
	for i, e := range e {
		fmt.Fprintf(os.Stderr, "%d / %d (%d%%)\n", i+1, ln, ((i+1)*100)/ln)
		fmt.Fprintf(buf,
			// 	{Id: 8, Pid: 9, IsRoot: true, Word: "...", Def: "..."},
			"\t{Id: %d, Pid: %d, IsRoot: %v, Word: %q, Def: %q},\n",
			e.Id, e.Pid, e.IsRoot, e.Word, e.Def,
		)
	}
	buf.WriteString("\n}\n")

	fmt.Println(buf.String())

}

func getRowData(stmt *sqlite.Stmt) Entry {
	return Entry{
		Id:          stmt.GetInt64("id"),
		Pid:         stmt.GetInt64("pid"),
		IsRoot:      stmt.GetBool("is_root"),
		IsHighlight: stmt.GetBool("highlight"),
		Word:        stmt.GetText("word"),
		Def:         stmt.GetText("def"),
	}
}
*/
