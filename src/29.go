package main

import (
	"encoding/json"
	"log"
)

const j = `{
  "fruit": ["banana"],
  "other": [
    "potato",
    "onion"
  ],
  "comment": "fruity"
}`

func main() {
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Panicln(err)
	}
	log.Printf("banana: %s", m)
}
