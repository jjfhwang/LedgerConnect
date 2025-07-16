// cmd/ledgerconnect/main.go
package main

import (
"flag"
"log"
"os"

"ledgerconnect/internal/ledgerconnect"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := ledgerconnect.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
