package handle

import (
	"fmt"
	"log"
)

// Error handle method to declutter other packages
func Error(e error) {
	if e != nil {
		fmt.Printf("An error occured: %s", e.Error())
		log.Fatal(e)
	}
}
