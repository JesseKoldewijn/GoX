package handlers

import (
	"io"
	"net/http"
	"time"
)

// GET /
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	htmlIndex := `<html>
					<head>
						<title>GoX</title>
					</head>
					<body>
						<div id="dynamic_value">` + now.String() + `</div>
						<button hx-get="/dynamic" hx-trigger="click" hx-swap="outerHTML" hx-target="#dynamic_value">Get dynamic value</button>
						<script src="https://unpkg.com/htmx.org@latest" defer></script>
					</body>
				</html>`

	_, err := io.WriteString(w, htmlIndex)
	if err != nil {
		return
	}
	request := r
	err = request.Body.Close()
	if err != nil {
		return
	}
}

// GET /dynamic
func HandleIndexDynamic(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	htmlIndexDynamic := "<div id=\"dynamic_value\">" + now.String() + "</div>"

	_, err := io.WriteString(w, htmlIndexDynamic)
	if err != nil {
		return
	}
	request := r
	err = request.Body.Close()
	if err != nil {
		return
	}
}
