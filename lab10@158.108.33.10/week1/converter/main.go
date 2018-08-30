package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// JSONData is data format for rendering graph
type JSONData struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

// Node is graph node format
type Node struct {
	Color      string  `json:"color"`
	Label      string  `json:"label"`
	Attributes Attr    `json:"attributes"`
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	ID         string  `json:"id"`
	Size       float64 `json:"size"`
}

// Edge is graph edge format
type Edge struct {
	SourceID   string  `json:"sourceID"`
	Attributes Attr    `json:"attributes"`
	TargetID   string  `json:"targetID"`
	Size       float64 `json:"size"`
}

// Attr is ...
type Attr struct {
}

var colors = []string{
	"#4f19c7",
	"#c71969",
	"#c71919",
	"#1984c7",
	"#c76919",
	"#8419c7",
	"#c79f19",
	"#c78419",
	"#c719b9",
	"#199fc7",
	"#9f19c7",
	"#69c719",
	"#19c719",
	"#1919c7",
	"#c74f19",
	"#19c7b9",
	"#9fc719",
	"#c7b919",
	"#b9c719",
	"#3419c7",
	"#19b9c7",
	"#34c719",
	"#19c784",
	"#c7199f",
	"#1969c7",
	"#c71984",
	"#1934c7",
	"#84c719",
	"#194fc7",
	"#c7194f",
	"#19c74f",
	"#b919c7",
	"#19c769",
	"#19c79f",
	"#4fc719",
	"#c73419",
	"#19c734",
	"#6919c7",
	"#c71934",
}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getRandomColor() string {
	return colors[rand.Intn(len(colors))]
}

func main() {
	data, err := ioutil.ReadFile("./internal_data.csv")
	failOnError(err)
	// fmt.Println(string(data))
	r := csv.NewReader(strings.NewReader(string(data)))

	jsonData := JSONData{
		Nodes: []Node{},
		Edges: []Edge{},
	}

	nodeMap := make(map[string]bool)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		failOnError(err)

		sourceAsn := record[0]
		asn := record[1]
		name := record[2]
		size, err := strconv.ParseFloat(record[4], 64)
		failOnError(err)

		if !nodeMap[asn] {
			node := Node{
				Label:      name,
				Color:      getRandomColor(),
				Attributes: Attr{},
				X:          math.Round(rand.NormFloat64()*100) / 100,
				Y:          math.Round(rand.NormFloat64()*100) / 100,
				ID:         asn,
				Size:       size,
			}
			jsonData.Nodes = append(jsonData.Nodes, node)
			nodeMap[asn] = true
		} else {
			for _, node := range jsonData.Nodes {
				if node.ID == asn {
					node.Size += size
				}
			}
		}

		edge := Edge{
			SourceID:   sourceAsn,
			Attributes: Attr{},
			TargetID:   asn,
			Size:       1,
		}
		jsonData.Edges = append(jsonData.Edges, edge)
	}

	jsonDataBytes, err := json.Marshal(jsonData)
	failOnError(err)
	fmt.Println(string(jsonDataBytes))
}
