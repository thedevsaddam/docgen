package main

import (
	"encoding/json"
	"io"
	"sort"
)

type (

	// Root describes the full data read from file
	Root struct {
		Info        Info         `json:"info"`
		Collections []Collection `json:"item"`
	}

	// Info describes the postman info section
	Info struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Schema      string `json:"schema"`
	}

	// Header describes header of  request
	Header struct {
		Key         string `json:"key"`
		Value       string `json:"value"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	// Field describes a field of request
	Field struct {
		Key         string `json:"key"`
		Value       string `json:"value"`
		Description string `json:"description"`
		Type        string `json:"type"`
	}

	// URL describes URL of the request
	URL struct {
		Raw         string   `json:"raw"`
		Host        []string `json:"host"`
		Path        []string `json:"path"`
		Description string   `json:"description"`
		Query       []Field  `json:"query"`
	}

	// Body describes a request body
	Body struct {
		Mode       string  `json:"mode"`
		FormData   []Field `json:"formdata"`
		URLEncoded []Field `json:"urlencoded"`
		Raw        string  `json:"raw"`
	}

	// Request describes a request
	Request struct {
		Method      string   `json:"method"`
		Headers     []Header `json:"header"`
		Body        Body     `json:"body"`
		URL         URL      `json:"url"`
		Description string   `json:"description"`
	}

	// Response describes a request resposne
	Response struct {
		ID      string   `json:"id"`
		Name    string   `json:"name"`
		Status  string   `json:"status"`
		Code    int      `json:"code"`
		Headers []Header `json:"header"`
		Body    string   `json:"body"`
	}

	// Item describes a request item
	Item struct {
		Name      string     `json:"name"`
		Items     []Item     `json:"item"`
		Request   Request    `json:"request"`
		Responses []Response `json:"response"` // TODO: need to build response
		Subfolder bool       `json:"_postman_isSubFolder"`
	}

	// Collection describes a request collection/item
	Collection struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Items       []Item `json:"item"`
		Item
	}

	// Assets represents template assets
	Assets struct {
		BootstrapJS  string
		BootstrapCSS string
		IndexHTML    string
		JqueryJS     string
		ScriptsJS    string
		StylesCSS    string

		IndexMarkdown string
	}
)

// Open open a new collection reader
func (r *Root) Open(rdr io.Reader) error {
	dcr := json.NewDecoder(rdr)
	if err := dcr.Decode(&r); err != nil {
		return err
	}
	r.build()
	r.sortCollections()

	r.removeEmptyCollections()

	return nil
}

// Build build UnCategorized collection
func (r *Root) build() {
	c := Collection{}
	c.Name = defaultCollection
	for i := len(r.Collections) - 1; i >= 0; i-- {
		if len(r.Collections[i].Items) <= 0 {
			if r.Collections[i].Request.Method == "" { //a collection with no sub-items and request method is empty
				continue
			}
			c.Items = append(c.Items, Item{
				Name:      r.Collections[i].Name,
				Request:   r.Collections[i].Request,
				Responses: r.Collections[i].Responses,
			})
			r.Collections = append(r.Collections[:i], r.Collections[i+1:]...)
		} else {
			clctn := Collection{}
			for j := len(r.Collections[i].Items) - 1; j >= 0; j-- {
				built := r.buildSubChildItems(r.Collections[i].Items[j], &clctn, r.Collections[i].Name)
				if built {
					r.Collections[i].Items = append(r.Collections[i].Items[:j], r.Collections[i].Items[j+1:]...) //removing the sub folder from the parent to make it a collection itself
				}
			}
		}
	}
	r.Collections = append(r.Collections, c)
}

// buildSubChildItems builds all the sub folder collections
func (r *Root) buildSubChildItems(itm Item, c *Collection, pn string) bool {
	if itm.Subfolder {
		clctn := Collection{}
		clctn.Name = pn + "/" + itm.Name
		for _, i := range itm.Items {
			r.buildSubChildItems(i, &clctn, clctn.Name)
		}
		r.Collections = append(r.Collections, clctn)
		return true
	}
	c.Items = append(c.Items, Item{
		Name:      itm.Name,
		Request:   itm.Request,
		Responses: itm.Responses,
	})

	return false
}

// sortCollections sorts the collections in the alphabetical order(except for the default)
func (r *Root) sortCollections() {
	sort.Slice(r.Collections, func(i int, j int) bool {
		if r.Collections[i].Name == defaultCollection {
			return false
		}
		return r.Collections[i].Name < r.Collections[j].Name
	})
}

//removeEmptyCollections removes any empty collection
func (r *Root) removeEmptyCollections() {
	for i, rc := range r.Collections {
		if len(rc.Items) == 0 {
			r.Collections = append(r.Collections[:i], r.Collections[i+1:]...) //popping the certain empty collection
			r.removeEmptyCollections()                                        //recurssion followed by break to ensure proper indexing after a pop
			break
		}
	}
	return
}
