package web

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	quotes "github.com/bm4cs/mastergo/quotes"
)

//type Quote struct {
//	Author string `json:"author"`
//	Text   string `json:"text"`
//	Source string `json:"source,omitempty"`
//}

type App struct {
	//storage map[string]*quotes.Quote
	storage quotes.DB //boltdb based storage
}

// createQuote receives a JSON object representing a Quote and stores
// the quote in the app storage.
func (app *App) createQuote(b []byte) error {

	// TODO:
	// * Unmarshal b into a Quote
	// * Check if the quote exists in the storage, return an error if it does
	// * Write the quote to the storage

	var quote quotes.Quote
	err := json.Unmarshal(b, &quote)
	if err != nil {
		return fmt.Errorf("json unmarshal fail: %s", err)
	}

	//if _, ok := app.storage[quote.Author]; ok {
	//	return fmt.Errorf("quote for author '%s' already exists", quote.Author)
	//}
	//
	//app.storage[quote.Author] = &quote

	app.storage.Create(&quote)

	return nil
}

// getQuote receives an author name and returns the corresponding
// quote of the author, or an error if no quote exists for this author.
func (app *App) getQuote(author string) ([]byte, error) {

	// TODO:
	// * Fetch author's quote from the storage
	// * Return an error if the quote does not exist
	// * Marshal the quote into JSON and return the result.
	//
	// Hint: Use
	//    json.MarshalIndent(q, "", "  ")
	// for producing formatted JSON - the test file expects this!

	//if q, ok := app.storage[author]; !ok {
	//	return nil, fmt.Errorf("no quote exists for author '%s'", author)
	//} else {
	//	b, err := json.MarshalIndent(&q, "", "  ")
	//
	//	if err != nil {
	//		return nil, fmt.Errorf("could not marshall quote: %s", err)
	//	}
	//
	//	return b, nil
	//}

	if q, err := app.storage.Get(author); err != nil {
		return nil, errors.Wrap(err, "failed to query db for author '#{author}'")
	} else {
		b, err := json.MarshalIndent(&q, "", "  ")

		if err != nil {
			return nil, errors.Wrap(err, "could not JSON marshal quote for author '#{author}'")
		}

		return b, nil
	}


}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "helo wurld\n")
	io.WriteString(w, r.URL.Path)
}

func (app *App) handleQuote(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "failed to read body: "+err.Error())
			return
		}
		err = app.createQuote(b)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "failed to createQuote(): "+err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	case "GET":
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 5 {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, fmt.Sprintf("failed to parse author from [%d] URL.Path: %s", len(parts), r.URL.Path))
			return
		}
		author := parts[4]
		quote, err := app.getQuote(author)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "failed to call getQuote(): "+err.Error())
			return
		}
		w.Write(quote)
	case "PUT":
		io.WriteString(w, "Update: "+r.URL.Path+"\n")
	case "DELETE":
		io.WriteString(w, "Delete: "+r.URL.Path+"\n")
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (app *App) handleQuotesList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		io.WriteString(w, r.URL.Path)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}


// Self hosted REST API using net/http
// To test:
//
// curl --header "Content-Type: application/json" \
//  --request POST \
//  --data '{"author": "Donald Knuth", "text": "Computers are good at following instructions, but not at reading your mind.", "source": "TAOCP"}' \
//  "http://localhost:9000/api/v1/quote/Donald Knuth"
//
// curl -XGET http://localhost:9000/api/v1/quote/Donald%20Knuth
func RunWebServer() {

	db, err := quotes.Open("quotesdb")

	if err != nil {
		log.Fatal("failed to initialise db: #{err.Error()}")
		return
	}

	app := &App{
		//storage: map[string]*quotes.Quote{
		//	"Bill Gates":  {"Bill Gates", "The computer was born to solve problems that did not exist before.", "Microsoft"},
		//	"Isaac Asimov": {"Isaac Asimov", "I do not fear computers. I fear lack of them.", "Science Fiction Author"},
		//	"Donald Knuth":  {"Donald Knuth", "Computers are good at following instructions, but not at reading your mind.", "TAOCP"},
		//},

		storage: *db,
	}

	const prefix string = "/api/v1/"
	http.HandleFunc("/", hello)
	http.HandleFunc(prefix+"quote/", app.handleQuote)
	http.HandleFunc(prefix+"quotes/", app.handleQuotesList)
	err = http.ListenAndServe("localhost:9000", nil)

	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
