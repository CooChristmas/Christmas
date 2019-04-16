package serve

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	c "github.com/CooChristmas/Christmas/src/color"
)

func main() {
	ServeFiles()
}

// References: https://gist.github.com/paulmach/7271283
func ServeFiles() {
	fmt.Print(c.WHITE, "")
	port := flag.String("p", "8100", "Serve")
	directory := flag.String("d", ".", "Christmas Directory")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("\n\nServing Christmas directory %s on HTTP port: %s\n", *directory, *port)
	fmt.Println("\nClose this window or press `ctrl` + `c` to stop.")
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}