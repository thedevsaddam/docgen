package collection

import (
	"encoding/json"
	"io"
	"sort"
)

var (
	defaultCollection = "Ungrouped"
)

type (
	// Documentation describes the full API documentation
	Documentation struct {
		Info        Info         `json:"info"`
		Collections []Collection `json:"item"`
		Variables   []Field      `json:"variable"`
		sortEnabled bool         `json:"-"`
	}

	// Info describes the postman info section
	Info struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Schema      string `json:"schema"`
	}

	// Field describes query, header, form-data and urlencoded field of request
	Field struct {
		Key         string `json:"key"`
		Value       string `json:"value"`
		Description string `json:"description"`
		Type        string `json:"type"`
		Disabled    bool   `json:"disabled"`
	}

	// URL describes URL of the request
	URL struct {
		Raw         string   `json:"raw"`
		Host        []string `json:"host"`
		Path        []string `json:"path"`
		Description string   `json:"description"`
		Query       []Field  `json:"query"`
		Variables   []Field  `json:"variable"`
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
		Method      string  `json:"method"`
		Headers     []Field `json:"header"`
		Body        Body    `json:"body"`
		URL         URL     `json:"url"`
		Description string  `json:"description"`
	}

	// Response describes a request resposne
	Response struct {
		ID              string  `json:"id"`
		Name            string  `json:"name"`
		OriginalRequest Request `json:"originalRequest"`
		Status          string  `json:"status"`
		Code            int     `json:"code"`
		Headers         []Field `json:"header"`
		Body            string  `json:"body"`
		PreviewLanguage string  `json:"_postman_previewlanguage"`
	}

	// Item describes a request item
	Item struct {
		Name        string     `json:"name"`
		Items       []Item     `json:"item"`
		Request     Request    `json:"request"`
		Responses   []Response `json:"response"`
		IsSubFolder bool       `json:"_postman_isSubFolder"`
	}

	// Collection describes a request collection/item
	Collection struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Items       []Item `json:"item"`
		Item
	}
)

// Open open a new collection reader
func (d *Documentation) Open(rdr io.Reader) error {
	dcr := json.NewDecoder(rdr)
	if err := dcr.Decode(&d); err != nil {
		return err
	}
	d.build()
	if d.sortEnabled {
		d.sortCollections()
	}

	// remove empty collections
	d.removeEmptyCollections()
	// after building all the collections, remove disabled fields
	d.removeItemRequestDisabledField()
	d.removeItemResponseRequestDisabledField()
	return nil
}

func (d *Documentation) WithSortEnabled() *Documentation {
	d.sortEnabled = true
	return d
}

// Build build UnCategorized collection
func (d *Documentation) build() {
	c := Collection{}
	c.Name = defaultCollection
	for i := len(d.Collections) - 1; i >= 0; i-- {
		if len(d.Collections[i].Items) <= 0 {
			if d.Collections[i].Request.Method == "" { //a collection with no sub-items and request method is empty
				continue
			}
			c.Items = append(c.Items, Item{
				Name:      d.Collections[i].Name,
				Request:   d.Collections[i].Request,
				Responses: d.Collections[i].Responses,
			})
			d.Collections = append(d.Collections[:i], d.Collections[i+1:]...)
		} else {
			collection := Collection{}
			for j := len(d.Collections[i].Items) - 1; j >= 0; j-- {
				built := d.buildSubChildItems(d.Collections[i].Items[j], &collection, d.Collections[i].Name)
				if built {
					d.Collections[i].Items = append(d.Collections[i].Items[:j], d.Collections[i].Items[j+1:]...) //removing the sub folder from the parent to make it a collection itself
				}
			}
		}
	}
	d.Collections = append(d.Collections, c)
}

// buildSubChildItems builds all the sub folder collections
func (d *Documentation) buildSubChildItems(itm Item, c *Collection, pn string) bool {
	if itm.IsSubFolder {
		collection := Collection{}
		collection.Name = pn + "/" + itm.Name
		collection.IsSubFolder = true
		for _, i := range itm.Items {
			d.buildSubChildItems(i, &collection, collection.Name)
		}
		d.Collections = append(d.Collections, collection)
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
func (d *Documentation) sortCollections() {
	sort.Slice(d.Collections, func(i int, j int) bool {
		if d.Collections[i].Name == defaultCollection {
			return false
		}
		return d.Collections[i].Name < d.Collections[j].Name
	})

	for index, _ := range d.Collections {
		sort.Slice(d.Collections[index].Items, func(i, j int) bool {
			return d.Collections[index].Items[i].Name < d.Collections[index].Items[j].Name
		})
	}
}

//removeEmptyCollections removes any empty collection
func (d *Documentation) removeEmptyCollections() {
	for i := 0; i < len(d.Collections); i++ {
		if len(d.Collections[i].Items) == 0 {
			d.Collections = append(d.Collections[:i], d.Collections[i+1:]...) //popping the certain empty collection
			d.removeEmptyCollections()                                        //recursion followed by break to ensure proper indexing after a pop
			break
		}
	}
	return
}

//removeEmptyCollections removes disabled field from request
func (d *Documentation) removeItemRequestDisabledField() {
	for i, c := range d.Collections {
		for j, _ := range c.Items {
			// remove disabled headers
			for k := len(d.Collections[i].Items[j].Request.Headers) - 1; k >= 0; k-- {
				if d.Collections[i].Items[j].Request.Headers[k].Disabled {
					d.Collections[i].Items[j].Request.Headers = append(d.Collections[i].Items[j].Request.Headers[:k], d.Collections[i].Items[j].Request.Headers[k+1:]...)
				}
			}
			// remove disabled queries
			for l := len(d.Collections[i].Items[j].Request.URL.Query) - 1; l >= 0; l-- {
				if d.Collections[i].Items[j].Request.URL.Query[l].Disabled {
					d.Collections[i].Items[j].Request.URL.Query = append(d.Collections[i].Items[j].Request.URL.Query[:l], d.Collections[i].Items[j].Request.URL.Query[l+1:]...)
				}
			}
			// remove disabled form-data
			for m := len(d.Collections[i].Items[j].Request.Body.FormData) - 1; m >= 0; m-- {
				if d.Collections[i].Items[j].Request.Body.FormData[m].Disabled {
					d.Collections[i].Items[j].Request.Body.FormData = append(d.Collections[i].Items[j].Request.Body.FormData[:m], d.Collections[i].Items[j].Request.Body.FormData[m+1:]...)
				}
			}
			// remove disabled urlencoded
			for n := len(d.Collections[i].Items[j].Request.Body.URLEncoded) - 1; n >= 0; n-- {
				if d.Collections[i].Items[j].Request.Body.URLEncoded[n].Disabled {
					d.Collections[i].Items[j].Request.Body.URLEncoded = append(d.Collections[i].Items[j].Request.Body.URLEncoded[:n], d.Collections[i].Items[j].Request.Body.URLEncoded[n+1:]...)
				}
			}
		}
	}
	return
}

//removeItemResponseRequestDisabledField removes disabled field from response originalRequest
func (d *Documentation) removeItemResponseRequestDisabledField() {
	for i, c := range d.Collections {
		for j, item := range c.Items {
			for o, _ := range item.Responses {
				// remove disabled headers
				for k := len(d.Collections[i].Items[j].Responses[o].OriginalRequest.Headers) - 1; k >= 0; k-- {
					if d.Collections[i].Items[j].Responses[o].OriginalRequest.Headers[k].Disabled {
						d.Collections[i].Items[j].Responses[o].OriginalRequest.Headers = append(d.Collections[i].Items[j].Responses[o].OriginalRequest.Headers[:k], d.Collections[i].Items[j].Responses[o].OriginalRequest.Headers[k+1:]...)
					}
				}
				// remove disabled queries
				for l := len(d.Collections[i].Items[j].Responses[o].OriginalRequest.URL.Query) - 1; l >= 0; l-- {
					if d.Collections[i].Items[j].Responses[o].OriginalRequest.URL.Query[l].Disabled {
						d.Collections[i].Items[j].Responses[o].OriginalRequest.URL.Query = append(d.Collections[i].Items[j].Responses[o].OriginalRequest.URL.Query[:l], d.Collections[i].Items[j].Responses[o].OriginalRequest.URL.Query[l+1:]...)
					}
				}
				// remove disabled form-data
				for m := len(d.Collections[i].Items[j].Responses[o].OriginalRequest.Body.FormData) - 1; m >= 0; m-- {
					if d.Collections[i].Items[j].Responses[o].OriginalRequest.Body.FormData[m].Disabled {
						d.Collections[i].Items[j].Responses[o].OriginalRequest.Body.FormData = append(d.Collections[i].Items[j].Responses[o].OriginalRequest.Body.FormData[:m], d.Collections[i].Items[j].Responses[o].OriginalRequest.Body.FormData[m+1:]...)
					}
				}
				// remove disabled urlencoded
				for n := len(d.Collections[i].Items[j].Responses[o].OriginalRequest.Body.URLEncoded) - 1; n >= 0; n-- {
					if d.Collections[i].Items[j].Responses[o].OriginalRequest.Body.URLEncoded[n].Disabled {
						d.Collections[i].Items[j].Responses[o].OriginalRequest.Body.URLEncoded = append(d.Collections[i].Items[j].Responses[o].OriginalRequest.Body.URLEncoded[:n], d.Collections[i].Items[j].Responses[o].OriginalRequest.Body.URLEncoded[n+1:]...)
					}
				}
			}
		}
	}
	return
}
