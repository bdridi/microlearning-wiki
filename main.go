package main

import (
    "fmt"
    "log"
	"net/http"
	"encoding/json"
	"strconv"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to Go WikiServer !")
    fmt.Println("Welcome to Go WikiServer !")
}

func wiki(w http.ResponseWriter, r *http.Request){
	
	// read category param
	query := r.URL.Query()
	categoryArr, present := query["category"]
	if !present {
		fmt.Println("category not present")
	}	
	category := categoryArr[0]
	// search wiki pages
	var result WikiResult
	response, err := http.Get("https://en.wikipedia.org/w/api.php?action=query&list=search&format=json&srsearch="+category)
	if err != nil {
        log.Fatal(err)
    }
	err1 := json.NewDecoder(response.Body).Decode(&result)
    if err1 != nil {
         http.Error(w, err1.Error(), http.StatusBadRequest)
         return
     }
	var wikis []Wiki
	for _, wiki := range result.Query.Search {
		pageidstr := strconv.Itoa(wiki.Pageid)
		url := "https://en.wikipedia.org/?curid="+pageidstr
		wikis = append(wikis, Wiki{Title: wiki.Title, Category: category, Url: url})
	} 	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wikis)

}
 
	
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api/v1/wiki", wiki)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	fmt.Println("Wiki Go server started ...")
	handleRequests()
}

// Structures 

type Wiki struct {
    Title string `json:"title"`
    Category string `json:"category"`
    Url string `json:"url"`
}

type WikiResult struct {
	Batchcomplete string `json:"batchcomplete"`
	Continue      struct {
		Sroffset int    `json:"sroffset"`
		Continue string `json:"continue"`
	} `json:"continue"`
	Query struct {
		Searchinfo struct {
			Totalhits         int    `json:"totalhits"`
			Suggestion        string `json:"suggestion"`
			Suggestionsnippet string `json:"suggestionsnippet"`
		} `json:"searchinfo"`
		Search []struct {
			Ns        int       `json:"ns"`
			Title     string    `json:"title"`
			Pageid    int       `json:"pageid"`
			Size      int       `json:"size"`
			Wordcount int       `json:"wordcount"`
			Snippet   string    `json:"snippet"`
			Timestamp time.Time `json:"timestamp"`
		} `json:"search"`
	} `json:"query"`
}