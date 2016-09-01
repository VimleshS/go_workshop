package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const account_template = `
Details :- 
   Account Holder {{.Name}}
   Running Balance {{.Balance}}
-----------------------
`

type Account struct {
	Name    string
	Balance float64
}

var accounts []Account

func main() {
	accounts = []Account{
		{"User-1", 20},
		{"User-2", 40},
		{"User-2", 60},
	}

	http.HandleFunc("/show", showHandler)

	http.ListenAndServe(":8000", nil)
}

func showHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	tmpl := template.Must(template.New("template").Parse(account_template))
	for _, account := range accounts {
		tmpl.Execute(w, account)
	}
}
