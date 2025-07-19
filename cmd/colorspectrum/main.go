// cmd/colorspectrum/main.go
package main

import (
"flag"
"log"
"os"

"colorspectrum/internal/colorspectrum"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := colorspectrum.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
