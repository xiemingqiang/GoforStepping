package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)


func getName() (string, error) {
	var name string
	db, err := sql.Open("mysql","root:root@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		return "", errors.Wrap(err, "DB cannot open")
	}

	defer db.Close()

	err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.Wrap(err, "get empty name")
		} else {
			return "", errors.Wrap(err, "get name error")
		}
	}

	return name, nil
}


func main() {
	name, err := getName()
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("%+v\n", err)
			return
		}
	}

	fmt.Println("name:", name)
}
