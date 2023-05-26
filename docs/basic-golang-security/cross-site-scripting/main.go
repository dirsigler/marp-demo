package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/comment", handleComment)

	fmt.Println("Starting insecure web server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := `<html>
	<head>
		<title>Insecure Web App</title>
	</head>
	<body>
		<h1>Comments</h1>
		{{range .}}
		<p>{{.}}</p>
		{{end}}
		<form action="/comment" method="post">
			<input type="text" name="comment" placeholder="Leave a comment">
			<input type="submit" value="Submit">
		</form>
	</body>
	</html>`

	t, err := template.New("home").Parse(tmpl)
	if err != nil {
		http.Error(w, "An error occurred", http.StatusInternalServerError)
		return
	}

	comments := []string{"Comment 1", "Comment 2", "Comment 3"}

	err = t.Execute(w, comments)
	if err != nil {
		http.Error(w, "An error occurred", http.StatusInternalServerError)
		return
	}
}

func handleComment(w http.ResponseWriter, r *http.Request) {
	comment := r.FormValue("comment")

	fmt.Fprintf(w, "New Comment: %s", comment)
}
