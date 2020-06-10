package main

import (
	"encoding/xml"
	"log"
)

const x = `<x>
<y><-- ok! --></y>
<z>banana</z>
</x>`

func main() {
	m := make(map[string]interface{})
	if err := xml.Unmarshal([]byte(x), &m); err != nil {
		log.Panicln(err)
	}
	log.Printf("banana: %s", m)
}
