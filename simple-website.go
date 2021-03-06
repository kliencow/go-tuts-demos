package main

/*
   This simple website shows different system for creating basic webpages.

    This is getting out of hand. I need to break this out in someway to make it managable.

*/

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"text/template"
)

/*
   The router is pretty easy to use. You don't have to name the functions the same as the path. I only do so to
   make creating these examples easire to pump out. aka Walt is Lazy.
*/
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	r.HandleFunc("/simplestTemplate", simplestTemplate)
	r.HandleFunc("/regexRouter/{foo:a.*}", regexRouter) // note this evaluates from beginning to end, like using ^[]$.
	r.HandleFunc("/inlineWithFunctionCall", inlineWithFunctionCall)
	r.HandleFunc("/inlineWithALoop", inlineWithALoop)

	http.Handle("/", r)
	http.ListenAndServe(":9999", nil)
}

/*
   This simple home handler show how to imput a parameter as a single option which is interpreted as %s for string.
*/
func Home(response http.ResponseWriter, request *http.Request) {
	page := `
<html>
    <head></head>
    <body>
        <h1>Index of examples</h1>
        <h2>Notes</h2>
        <div>
            %s
        </div>
        <ul>
            <li><a href="/simplestTemplate">Simplest Template</a> -- a template with variables subbed in via Fprintf.</li>
            <li><a href="/regexRouter/apple">Regex Var Success!</a> -- regex router with a named variable defined by the regex</li>
            <li><a href="/regexRouter/zounds">Regex Var Fail!</a> -- regex router with a named variable defined by the regex</li>
            <li><a href="/inlineWithFunctionCall">Inline Template With Function Call</a> -- a template with a function call.</li> 
            <li><a href="/inlineWithALoop">Inline Template with a loop</a> -- a template with a loop.</li> 
        </ul>
    </body>
</html>`

	f, err := os.Open(".")

	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	fileInfo, _ := f.Stat()

	testString := fileInfo.IsDir()

	fmt.Fprintf(response, page, testString)
}

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

func regexRouter(response http.ResponseWriter, request *http.Request) {
	foo := mux.Vars(request)["foo"]

	page := `
<html>
    <head></head>
    <body>
        <h1>Value of foo: </h1>
        <p>%s</p>
    </body>
</html>`
	fmt.Fprintf(response, page, foo)
}

type PageFuncDeets struct {
	Title   string "TEST PAGE"
	seed    int
	FuncYou func() string
}

func (pfd PageFuncDeets) Rando() int {
	return pfd.seed + 6
}

func (pfd PageFuncDeets) ThisThat() string {
	return pfd.FuncYou()
}

/*

*/
func inlineWithFunctionCall(response http.ResponseWriter, request *http.Request) {
	details := PageFuncDeets{
		Title: "Function Call Test Page",
		seed:  10,
		FuncYou: func() string {
			return "Whoa whoa whoa whoa whoa there"
		},
	}

	page := `
<html>
    <head></head>
    <body>
            <h1>{{.Title}}</h1>        
            <h3>The number: {{.Rando}}</h3>
            <h3>Another Func: {{.ThisThat}}</h3>
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

type PageLoopDeets struct {
	Title      string "TEST PAGE"
	Ranger     map[string]string
	StartIndex int
	Factor     int
	MaxVal     int
}

func inlineWithALoop(response http.ResponseWriter, request *http.Request) {
	details := PageLoopDeets{
		Title: "Loop Page",
		Ranger: map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
			"key4": "value4",
		},
		StartIndex: 37,
		Factor:     3,
		MaxVal:     1200,
	}

	page := `
<html>
    <head></head>
    <body>
        <h1>{{.Title}}</h1>
        <h2>THE SIMPLE RANGE</h2>
        <div>
            <ol>
                {{range .Ranger}}
                    <li>{{.}}</li>
                {{end}}
            </ol>
        </div>

        <h2>THE KEY-VALUE RANGE</h2>
        <div>
            <ol>
                {{range $key, $value := .Ranger}}
                    <li>{{$key}}::{{$value}}</li>
                {{end}}
            </ol>
        </div>

        <h2>THE LOOP</h2>
        <div>
            <ol>

            </ol>
        </div>        
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
