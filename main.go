package main

// import log "github.com/sirupsen/logrus"
import log "github.com/sirupsen/logrus"

func main() {

	defer fileClose()

	slice := []string{"a", "b"}
	log.Println(slice, len(slice))
	slice = append(slice, "c", "d")
	log.Println(slice, len(slice))

}
