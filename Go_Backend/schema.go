package main

type Post struct {
	_id    string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
	Author string `json:"author,omitempty"`
}
