// Model the vnd.collection+json media format defined in
// http://amundsen.com/media-types/collection/
//
package collection

import (
	"net/url"
)

type Collection struct {
	Version string `json:"version"`
	Href url.URL `json:"href"`
	Links []Link `json:"links,omitempty"`
	Items []Item `json:"items,omitempty"`
	queries []Query `json:"queries,omitempty"`
	template Template `json:"template,omitempty"`
	error Error `json:"error,omitempty"`
	
}

type Link struct {
	Href url.URL `json:"href"`
	Rel string `json:"rel"`
	Name string `json:"name,omitempty"`
	Render string `json:"render,omitempty"`
	Prompt string `json:"prompt,omitempty"`
}

// value can be string, number, boolean or null
type Data struct {
	Name string `json:"name"`
	Value interface{}
	Prompt string `json:"prompt,omitempty"`
}

type Item struct {
	Href url.URL `json:"href"`
	Data []Data `json:"data,omitempty"`
	Links []Link `json:"links,omitempty"`
}

type Query struct {
	Href url.URL `json:"href"`
	Rel string `json:"rel"`
	Name string `json:"name,omitempty"`
	Prompt string `json:"prompt,omitempty"`
	Data []Data `json:"data,omitempty"`
}

type Template struct {
	Data []Data `json:"data,omitempty"`
}

type Error struct {
	Title string `json":"title,omitempty"`
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

