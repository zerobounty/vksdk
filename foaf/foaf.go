/*
Package foaf wrapper for FOAF.

Specification https://web.archive.org/web/20140909053226/http://api.yandex.ru/blogs/doc/indexation/concepts/what-is-foaf.xml
*/
package foaf // import "github.com/SevereCloud/vksdk/foaf"

import (
	"context"
	"encoding/xml"
	"net/http"

	"github.com/SevereCloud/vksdk/internal"
	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/net/html/charset"
)

// FOAFURL url foaf
const FOAFURL = "https://vk.com/foaf.php"

// HTTPClient is the context key to use with golang.org/x/net/context's
// WithValue function to associate an *http.Client value with a context.
var HTTPClient internal.ContextKey // nolint:gochecknoglobals

// rdf is a standard model for data interchange on the Web.
//
// See https://www.w3.org/RDF/
type rdf struct {
	XMLName xml.Name `xml:"RDF"`
	Lang    string   `xml:"lang,attr"`
	Rdf     string   `xml:"rdf,attr"`
	Rdfs    string   `xml:"rdfs,attr"`
	Foaf    string   `xml:"foaf,attr"`
	Ya      string   `xml:"ya,attr"`
	Img     string   `xml:"img,attr"`
	Dc      string   `xml:"dc,attr"`
	Person  Person   `xml:"Person"`
	Group   Group    `xml:"Group"`
}

// URI struct
type URI struct {
	Primary  string `xml:"primary,attr"`
	Resource string `xml:"resource,attr"`
}

// Weblog struct
type Weblog struct {
	Title    string `xml:"title,attr"`
	Resource string `xml:"resource,attr"`
}

// Date may be used to express temporal information at any level of granularity.
//
// Use time.Parse(time.RFC3339, v.Date)
type Date struct {
	Date string `xml:"date,attr"`
}

// getFoaf return RDF
func getFoaf(ctx context.Context, req *http.Request) (r rdf, err error) {
	resp, err := ctxhttp.Do(ctx, internal.ContextClient(ctx), req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel // For support WINDOWS-1251
	err = decoder.Decode(&r)

	return
}
