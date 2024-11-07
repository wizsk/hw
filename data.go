package main

import (
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

func searchByTxt(conn *sqlite.Conn, str string) (Entries, error) {
	e := Entries{}
	const q = `SELECT word, is_root, REPLACE(def, ?, '<span style="background: gray;">' || ? || '</span>') AS def
 	FROM dict WHERE INSTR(def, ?) > 0 LIMIT 50`

	return e, sqlitex.Execute(conn, q, &sqlitex.ExecOptions{
		Args: []any{str, str, str},
		ResultFunc: func(stmt *sqlite.Stmt) error {
			e = append(e, getRowData(stmt))
			return nil
		},
	})
}

func searchByRoot(conn *sqlite.Conn, root string) (Entries, error) {
	e := Entries{}
	const q = `SELECT id, word, CASE word when ? then 1 else 0 end as highlight, def, is_root
		FROM dict WHERE pid IN (SELECT pid FROM dict WHERE word = ?) ORDER BY id;`

	return e, sqlitex.Execute(conn, q, &sqlitex.ExecOptions{
		Args: []any{root, root},

		ResultFunc: func(stmt *sqlite.Stmt) error {
			e = append(e, getRowData(stmt))
			return nil
		},
	})
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
