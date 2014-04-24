package main

/*
   This simple website shows different system for creating basic webpages.

    This is getting out of hand. I need to break this out in someway to make it managable.

*/

import (
	"fmt"
	"github.com/gorilla/mux"
	//"html"
	"io/ioutil"
	"net/http"
	"os"
	//"text/template"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	r.HandleFunc("/pages/{page}", pages)
	r.HandleFunc("/simpleImportedTemplate", simpleImportedTemplate)
	r.HandleFunc("/complexImportedTemplate", complexImportedTemplate)
	r.HandleFunc("/fullyTemplatedTemplate", fullyTemplatedTemplate)

	// How to serve images, css, and js easily, but with limited control
	r.Handle("/images/{imageFile}", http.StripPrefix("/images/", http.FileServer(http.Dir("sitehelpers/simple-res/img/"))))
	r.Handle("/css/{cssFile}", http.StripPrefix("/css/", http.FileServer(http.Dir("sitehelpers/simple-res/css/"))))
	r.Handle("/js/{jsFile}", http.StripPrefix("/js/", http.FileServer(http.Dir("sitehelpers/simple-res/js/"))))

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
            <li><a href="/pages/simple-page">Simple Page Load</a> -- this is a demonstration of simply loading html pages and sending them out. No templating involved.</li>
            <li><a href="/simpleImportedTemplate">Simple Imported Template</a> -- a template imported from an HTML file.</li>
            <li><a href="/complexImportedTemplate">Complex Imported Template</a> -- a template imported from an HTML file with lots of variables.</li>
            <li><a href="/fullyTemplatedTemplate">Fully Templated Template</a> -- a template that incorporates loops, other template files, variables, and loops. Also includes css, images, and js files.</li>
        </ul>
        <ul>
            <li><a href="/image/modemConvo.jpg">Image load example</a> -- Example of loading an images.</li>
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
   Load a simple html page from the simple html page folder.

   NOTES:
    This should be considered rather sanitary as long as any page is allowed to be viewed by anyone.
    This will not allow dot-dots, so attackers cannot load anything on the system. For instance, this does not work:

    $ curl localhost:9999/pages/../../../../basics.go

    while this does:

    $ cat sitehelpers/simple-res/templates/pages/../../../../basics.go
*/
func pages(response http.ResponseWriter, request *http.Request) {
	pagesFolder := "sitehelpers/simple-res/templates/pages/"
	page := mux.Vars(request)["page"]

	filename := pagesFolder + page + ".html"

	if exists(filename) {
		html, _ := ioutil.ReadFile(filename)
		fmt.Fprintf(response, string(html))
	} else {
		do404(response)
	}
}

func simpleImportedTemplate(response http.ResponseWriter, request *http.Request) {

}

func complexImportedTemplate(response http.ResponseWriter, request *http.Request) {

}

func fullyTemplatedTemplate(response http.ResponseWriter, request *http.Request) {

}

/***

    HELPERS BELOW: MOVE TO NEW PACKAGE LATER

****/

/*
   Return simple 404
   this is broken out because I'd like to find different ways to return 404s. This simple text return is rather limited.
*/
func do404(response http.ResponseWriter) {
	http.Error(response, "404! OH NO! This page does not exist!", 404)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
