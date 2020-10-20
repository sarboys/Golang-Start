package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq" // here
	"github.com/sarboys/Golang-Start/hello"
)

type status struct {
	Name string
}
type info struct {
	Markets []struct {
		Name  string
		Price int
	}
}

func main() {

	connStr := "host=192.168.1.12 port=5432 user=vu_miheikin password=dhgFFx1VP2 dbname=regofficex sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	fmt.Println(db)
	fmt.Println(err)

	fmt.Println(hello.CallFromHello())
	st := status{Name: "Sergey"}
	st.Name = "Sergey Polyakov"
	fmt.Println(st)

	text := `{"markets": [{"name":"Alice", "price" : 25},{"name":"Sergey", "price" : 234},{"name":"Andrey", "price" : 2415}]}`

	infos := new(info)
	json.Unmarshal([]byte(text), &infos)
	fmt.Println(infos.Markets)

	for i := range infos.Markets {
		fmt.Println(i, infos.Markets[i].Name, infos.Markets[i].Price)
	}
}
