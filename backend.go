package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"
	//"html"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var messages []string
var senders []string

func clear(w http.ResponseWriter, r *http.Request) {
	messages = nil
	senders = nil
}

func submit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var message = r.FormValue("message");
	var sender = r.FormValue("sender");
	fmt.Println(message)

	if sender == "" {
		sender = "unkown";
	}

	if message != "" {
		messages = append(messages, message);
		senders = append(senders, sender);
	}

	fmt.Fprintf(w, "<html>")
	fmt.Fprintf(w, "<header>")
	fmt.Fprintf(w, "   <style> body { font-family: sans-serif; } </style>")
	fmt.Fprintf(w, "</header>")
	fmt.Fprintf(w, "<body>")
	fmt.Fprintf(w, "   <div>")
	fmt.Fprintf(w, "      <h2>Powercoders bootcamp message board</h2>")
	fmt.Fprintf(w, "      <p><em>You can post to this board by sending POST request to this page's url with a 'message' and a 'sender' field.</em></h2>")
	fmt.Fprintf(w, "      <p><em>To see new messages you can <a href=''>reload</a> the page.</em></h2>")
	for i := 0; i < len(messages); i++ {
		fmt.Fprintf(w, "      <p><b>%s</b>: %s</p>", senders[i], messages[i])
	}
	fmt.Fprintf(w, "   </div>")
	fmt.Fprintf(w, "   <form action=\"\" method=\"post\">")
	fmt.Fprintf(w, "      <p><label>your name:</label> <input name=\"sender\"></p>")
	fmt.Fprintf(w, "      <p><label>message:</label> <input name=\"message\"></p>")
	fmt.Fprintf(w, "      <p><input type=\"submit\"></p>")
	fmt.Fprintf(w, "   </form>")
	fmt.Fprintf(w, "</body>")
	fmt.Fprintf(w, "</html>")

	// hpage, err := ioutil.ReadFile("index.html")
	// check(err)
	// fmt.Fprintf(w, "%s", hpage)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8");
}

// for AJAX practice
func messages_json(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	fmt.Fprintln(w, "{")
	fmt.Fprintln(w, "  \"messages\": [")
	for i := 0; i < len(messages); i++ {
		fmt.Fprintf(w, "    {\"sender\": \"%s\", \"text\": \"%s\" }", senders[i], messages[i]);
		if (i < len(messages)-1) {
			fmt.Fprintf(w, ",");
		}
		fmt.Fprintf(w, "\n");
	}
	fmt.Fprintln(w, "  ]");
	fmt.Fprintln(w, "}");
}

func main() {
	http.HandleFunc("/submit", submit)
	http.HandleFunc("/", submit)
	http.HandleFunc("/clear", clear)
	http.HandleFunc("/messages-json", messages_json)
	port := ":8080";
	fmt.Println("Serving on port", port );
	j := http.ListenAndServe(port, nil)
	check(j)
}
