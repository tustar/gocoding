package main

type Folder struct {
	ID          int64  `json:"id, omitemty"`
	App         string `json:"app,omitempty"`
	Data        string `json:"data, omitempty"`
	Description string `json:"description, omitempty"`
	Depth       int8   `json:"depth, omitempty"`
	Ttype       int8   `json:"type, omitempty"`
}
