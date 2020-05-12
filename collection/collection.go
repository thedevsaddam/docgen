package collection

import (
	"encoding/json"
	"io"
	"sort"
)

var (
	defaultCollection = "Default"
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
		Disabled    bool   `json:"disabled"`
	}

	// URL describes URL of the request
	URL struct {
		Raw         string     `json:"raw"`
		Host        []string   `json:"host"`
		Path        []string   `json:"path"`
		Description string     `json:"description"`
		Query       []Field    `json:"query"`
		Variables   []Variable `json:"variable"`
	}

	Variable struct {
		Key         string `json:"key"`
		Value       string `json:"value"`
		Description string `json:"description"`
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
		ID              string   `json:"id"`
		Name            string   `json:"name"`
		OriginalRequest Request  `json:"originalRequest"`
		Status          string   `json:"status"`
		Code            int      `json:"code"`
		Headers         []Header `json:"header"`
		Body            string   `json:"body"`
		PreviewLanguage string   `json:"_postman_previewlanguage"`
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
func (r *Documentation) Open(rdr io.Reader) error {
	dcr := json.NewDecoder(rdr)
	if err := dcr.Decode(&r); err != nil {
		return err
	}
	r.build()
	if r.sortEnabled {
		r.sortCollections()
	}

	r.removeEmptyCollections()

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
}

//removeEmptyCollections removes any empty collection
func (d *Documentation) removeEmptyCollections() {
	for i, c := range d.Collections {
		if len(c.Items) == 0 {
			d.Collections = append(d.Collections[:i], d.Collections[i+1:]...) //popping the certain empty collection
			d.removeEmptyCollections()                                        //recursion followed by break to ensure proper indexing after a pop
			break
		}
	}
	return
}
