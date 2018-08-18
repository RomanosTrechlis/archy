package graph

import (
	"database/sql"
	"log"

	_ "gopkg.in/cq.v1"
	"gopkg.in/cq.v1/types"
	"fmt"
	"strconv"
)

type E struct {
	Name string
	Type string
	From string
	To string
}
type V struct {
	Id string
	Name string
	Type string
	Out []E
	In []E
}


var db *sql.DB

func connect() *sql.DB {
	if db == nil {
		dbI, err := sql.Open("neo4j-cypher", "http://neo4j:password@192.168.99.100:7474")
		if err != nil {
			log.Fatal(err)
		}
		db = dbI
	}
	return db
}

func Close() {
	db.Close()
}

func Get() []*LightElement{
	db = connect()
	// MATCH (n1)-[r]->(n2) RETURN r, n1, n2 LIMIT 25
	stmt, err := db.Prepare(`
		match (n1)-[r]->(n2) 
		return r, n1, n2
		limit 25
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	type test struct {
		r, n1, n2 map[string]types.CypherValue
	}
	res := make([]*LightElement, 0)
	for rows.Next() {
		t := new(test)
		err := rows.Scan(&t.r, &t.n1, &t.n2)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, createLightElements(t.r, t.n1, t.n2)...)
	}

	return res
}

func createLightElements(r, n1, n2 map[string]types.CypherValue) []*LightElement {
	elements := make([]*LightElement, 0)
	rel := NewElement(EDGE)
	rel.Data.Id = fmt.Sprintf("%v", r["self"].Val)
	rel.Data.Name = fmt.Sprintf("%v", r["self"].Val)
	rel.Data.Source = fmt.Sprintf("%v", r["start"].Val)
	rel.Data.Target = fmt.Sprintf("%v", r["end"].Val)

	node1 := NewElement(NODE)
	node1.Data.Id = fmt.Sprintf("%v", n1["self"].Val)
	id, _ := strconv.Atoi(fmt.Sprintf("%v", n1["self"].Val))
	node1.Data.IdInt = id
	node1.Data.Name = fmt.Sprintf("%v", n1["data"].Val.(map[string]string)["name"])

	node2 := NewElement(NODE)
	node2.Data.Id = fmt.Sprintf("%v", n2["self"].Val)
	id, _ = strconv.Atoi(fmt.Sprintf("%v", n2["self"].Val))
	node2.Data.IdInt = id
	node2.Data.Name = fmt.Sprintf("%v", n2["data"].Val.(map[string]string)["name"])

	elements = append(elements, rel, node1, node2)
	return elements
}

func test() {
	db, err := sql.Open("neo4j-cypher", "http://neo4j:password@192.168.99.100:7474")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// MATCH (n1)-[r]->(n2) RETURN r, n1, n2 LIMIT 25
	stmt, err := db.Prepare(`
		match (n1)-[r]->(n2) 
		return r, n1, n2
		limit 25
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	type test struct {
		r, n1, n2 map[string]types.CypherValue
	}
	res := make([]*test, 0)
	for rows.Next() {
		t := new(test)
		err := rows.Scan(&t.r, &t.n1, &t.n2)
		if err != nil {
			log.Fatal(err)
		}
		//log.Printf("%v\n\n\n", t)
		res = append(res, t)
	}

	for _, t := range res {
		fmt.Printf("r\n=\n%v\n\n", t.r)
		fmt.Printf("n1\n==\n%v\n\n", t.n1)
		fmt.Printf("n2\n==\n%v\n\n", t.n2)
		fmt.Print("==========================\n\n")
	}

	for _, t := range res {
		rel := t.r["metadata"].Val
		r := rel.(map[string]types.CypherValue)["type"]
		data1 := t.n1["data"]
		data2 := t.n2["data"]

		fmt.Printf("Rel --> Type: %v, Value: %v\n", r.Type, r.Val)
		fmt.Printf("Data1 --> Type: %v, Value: %v\n", data1.Type, data1.Val)
		fmt.Printf("Data2 --> Type: %v, Value: %v\n", data2.Type, data2.Val)
		fmt.Println()
	}


	nodes := make(map[string]V, 0)
	rels := make(map[string]E, 0)
	for _, t := range res {
		r := E{
			Name: fmt.Sprintf("%v", t.r["self"].Val),
			Type: fmt.Sprintf("%v", t.r["metadata"].Val.(map[string]types.CypherValue)["type"].Val),
			From: fmt.Sprintf("%v", t.r["start"].Val),
			To: fmt.Sprintf("%v", t.r["end"].Val),
		}
		n1 := V{
			Id: fmt.Sprintf("%v", t.n1["self"].Val),
			Name: fmt.Sprintf("%v", t.n1["data"].Val.(map[string]string)["name"]),
			Type: fmt.Sprintf("%v", t.n1["metadata"].Val),
			In: make([]E, 0),
			Out: make([]E, 0),
		}
		n2 := V{
			Id: fmt.Sprintf("%v", t.n2["self"].Val),
			Name: fmt.Sprintf("%v", t.n2["data"].Val.(map[string]string)["name"]),
			Type: fmt.Sprintf("%v", t.n1["labels"].Val),
			In: make([]E, 0),
			Out: make([]E, 0),
		}

		if r.From == n1.Id {
			n1.Out = append(n1.Out, r)
		}
		if r.From == n2.Id {
			n2.Out = append(n2.Out, r)
		}
		if r.To == n1.Id {
			n1.In = append(n1.In, r)
		}
		if r.To == n2.Id {
			n2.In = append(n2.In, r)
		}

		nodes[n1.Id] =  n1
		nodes[n2.Id] =  n2
		rels[r.Name] =  r
	}

	for _, r := range rels {
		fmt.Printf("%s (%s) ----> %s ----> %s (%s)\n\n", nodes[r.From].Name, nodes[r.From].Type,
			r.Type, nodes[r.To].Name, nodes[r.From].Type)
	}
}
