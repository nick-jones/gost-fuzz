package main

import (
	"log"
)

const lorum = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at mattis magna.
Aenean vitae est pretium, semper odio fermentum, lacinia elit. Duis viverra sem
orci, id fringilla lectus suscipit sit amet. Integer quis arcu commodo, congue
ipsum id, feugiat tellus. Cras pretium feugiat commodo. Integer dapibus lectus
interdum sapien dignissim dignissim. Fusce dictum maximus velit vitae convallis.
Proin vel mauris in eros auctor finibus. Morbi ipsum enim, bibendum in nibh ac,
cursus viverra metus.

Phasellus consectetur fringilla imperdiet. Nulla facilisi. Morbi felis urna,
accumsan eu imperdiet eget, facilisis in nisi. Maecenas urna purus, rutrum ac
posuere ut, porta ac est. Etiam ut nunc et eros cursus gravida quis sit amet
odio. Vivamus rhoncus mollis libero. Morbi aliquet pulvinar arcu, quis posuere
quam. Pellentesque pretium ullamcorper arcu. Duis bibendum quam vitae metus
imperdiet, in pretium mi commodo. Nam pellentesque consectetur varius. Sed
elementum sit amet neque eu tempus. Aliquam non eleifend ante. Quisque sodales
nibh eu porta gravida. Mauris ex augue, imperdiet et consequat et, mollis eu
nulla. Donec egestas, massa id semper dignissim, dui nisl efficitur odio, non
auctor dui mi ac risus. Quisque sed ex mi.`

func main() {
	log.Printf(lorum)
}
