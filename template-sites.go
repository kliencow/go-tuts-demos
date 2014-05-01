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
	"text/template"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	r.HandleFunc("/pages/{page}", pages)
	r.HandleFunc("/compositePages/{body}", compositePages)
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
        <h3>Example Pages</h3>
        <ul>
            <li><a href="/pages/simple-page">Simple Page Load</a> -- this is a demonstration of simply loading html pages and sending them out. No templating involved.</li>
            <li><a href="/compositePages/page1">Composite Page Load</a> -- This is just like the page load, but it uses a common template to wrap the pages. In other words, content only with a common wrapper.</li>
            <li><a href="/simpleImportedTemplate">Simple Imported Template</a> -- a template imported from an HTML file.</li>
            <li><a href="/complexImportedTemplate">Complex Imported Template</a> -- a template imported from an HTML file with lots of variables.</li>
            <li><a href="/fullyTemplatedTemplate">Fully Templated Template</a> -- a template that incorporates loops, other template files, variables, and loops. Also includes css, images, and js files.</li>
        </ul>
        <h3>Utility Functions</h3>
        <ul>
            <li><a href="/images/modemConvo.jpg">Image load example</a> -- Example of loading an images.</li>
            <li><a href="/js/alert.js">Javascript load example</a> -- Example of loading javascipt.</li>
            <li><a href="/css/css1.css">CSS load example</a> -- Example of loading css.</li>
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

/*
	This is an example of just printing a page with some custom templates in it.
	This is a pattern that can be used to print static HTML pages with common headers. So
	they really aren't completely static.

	The interesting part about this function is that we can name all sorts of templates in line with
	each other. So, instead of having a page body determine it's title, we can just make these things
	inline with one another.

	TODO: rename the variables to be better. They kinda suck.
*/
func compositePages(response http.ResponseWriter, request *http.Request) {
	templatesFolder := "sitehelpers/simple-res/templates/"
	bodiesFolder := templatesFolder + "bodies/"

	body := mux.Vars(request)["body"]

	bodyFilename := bodiesFolder + body + ".html"
	commonFilename := templatesFolder + "common/common.html"

	if exists(commonFilename) && exists(bodyFilename) {
		templateSet, _ := template.ParseFiles(commonFilename, bodyFilename)
		templateSet.ExecuteTemplate(response, "base", nil)
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
