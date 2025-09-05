package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")
var templ = template.Must(template.New("qr").Parse(`
<!DOCTYPE html>
<html>
<head>
<title>Show QR Code</title>
</head>
<body>
{{if .}}
<h1>QR Code for "{{.}}"</h1>
<img src="https://api.qrserver.com/v1/create-qr-code/?data={{.}}">
{{end}}
<form method="GET" action="/" name="qrform">
	<label for="text">Enter text to generate QR code:</label><br>
	<input type="text" id="text" name="text" required>
	<input type="submit" value="Generate QR Code">
</form>
</body>
</html>
`))

// The showqr package is a minimal web server that displays a QR code
// for any URL path. For example, accessing http://localhost:8080/hello
// will display a QR code for the URL "hello".
// To run the server, use the command: go run showqr.go
// Then open a web browser and navigate to http://localhost:8080/yourtext
// to see the QR code for "yourtext".
// You can change the port by using the -addr flag, e.g., -addr=:9090
// to listen on port 9090.
// The server uses the free QR code generation API from goqr.me.
// Note: This package is intended for educational purposes and may not be suitable for production use.
// Be mindful of the usage limits of the QR code generation service.
// For more robust solutions, consider using a dedicated QR code generation library or service.
// This package is self-contained and does not require any external dependencies beyond the standard library.
// It is designed to be simple and easy to understand, making it a good starting point for learning about web servers in Go.
// Feel free to modify and extend the code to suit your needs.
// Enjoy!

func main() {
	flag.Parse()
	// handler function to serve the QR code page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		text := r.URL.Query().Get("text")
		templ.Execute(w, text)
	})
	log.Printf("Starting server on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
