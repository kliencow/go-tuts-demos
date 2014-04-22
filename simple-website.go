package main

/*
   This simple website shows different system for creating basic webpages.

*/

import (
	"fmt"
	"github.com/gorilla/mux"
	//"html/template"
	"net/http"
	"text/template"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/simplestTemplate", simplestTemplate)
	r.HandleFunc("/simpleInlineTemplate", simpleInlineTemplate)
	r.HandleFunc("/simpleImportedTemplate", simpleImportedTemplate)
	r.HandleFunc("/complexImportedTemplate", complexImportedTemplate)
	r.HandleFunc("/fullyTemplatedTemplate", fullyTemplatedTemplate)
	http.Handle("/", r)
	http.ListenAndServe(":9999", nil)
}

/*
   This simple home handler show how to imput a parameter as a single option which is interpreted as %s for string.
*/
func HomeHandler(response http.ResponseWriter, request *http.Request) {
	page := `
<html>
    <head></head>
    <body>
        <h1>Index of examples</h1>
        <ul>
            <li><a href="/simplestTemplate">Simplest Template</a> -- a template with variables subbed in via Fprintf.</li>
            <li><a href="/simpleInlineTemplate">Simple Inline Template</a> -- a template with multiple variables and a loop./li>
            <li><a href="/simpleImportedTemplate">Simple Imported Template</a> -- a template imported from an HTML file.</li>
            <li><a href="/complexImportedTemplate">Complex Imported Template</a> -- a template imported from an HTML file with lots of variables.</li>
            <li><a href="/fullyTemplatedTemplate">Fully Templated Template</a> -- a template that incorporates loops, other template files, variables, and loops. Also includes css, images, and js files.</li>            
        </ul>
    </body>
</html>`
	fmt.Fprintf(response, page)
}

/**
 *
 *  EXAMPLE 1: SIMPLEST TEMPLATE
 *
 */

/*
   This example shows how to output using fmt.Fprintf. The trick going on here is we can pass as many variables to Fprintf as we'd like. The
   page then just formats each variable accounding to our rules. In this case, %s for string and %d for decimal.

   for more formatting options, check out: http://golang.org/pkg/fmt
*/
func simplestTemplate(response http.ResponseWriter, request *http.Request) {
	title := "TEST PAGE"
	title2 := "Suprise another variable"
	oneHundred := 100
	page := `
<html>
    <head></head>
    <body>
        <h1>%s</h1>
        <h2>%s</h2>
        <h3>The number: %d</h3>
    </body>
</html>`

	fmt.Fprintf(response, page, title, title2, oneHundred)
}

/*
 *
 *  EXAMPLE 2: SIMPLE INLINE TEMPLATE
 *
 */

/*
   Example struct for simpleInlineTemplate
*/
type PageDeets struct {
	Title        string "TEST PAGE"
	AnotherTitle string
	OneHundred   int
}

/*
   Simple example of using a struct to pass data to a genuine template.
   Note the odd fact that the template must be named. I'm not sure why this is.

   TODO: Find out why teh template must have a name.
*/
func simpleInlineTemplate(response http.ResponseWriter, request *http.Request) {
	details := PageDeets{
		Title:        "Test Page",
		AnotherTitle: "Another TEST page",
		OneHundred:   100,
	}
	page := `
<html>
    <head></head>
    <body>
        <h1>{{.Title}}</h1>
        <h2>{{.AnotherTitle}}</h3>
        <h3>The number: {{.OneHundred}}</h3>
    </body>
</html>`
	t, err := template.New("Test").Parse(page)
	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(response, details)
	if err != nil {
		fmt.Println(err)
	}

}

func simpleImportedTemplate(response http.ResponseWriter, request *http.Request) {

}

func complexImportedTemplate(response http.ResponseWriter, request *http.Request) {

}

func fullyTemplatedTemplate(response http.ResponseWriter, request *http.Request) {

}
