package model

import "time"

type Package struct {
	ID             string               `json:"_id"`
	Rev            string               `json:"_rev"`
	Name           string               `json:"name"`
	Description    string               `json:"description"`
	DistTags       *DistTags            `json:"dist-tags"`
	Versions       map[string]*Version  `json:"versions"`
	Readme         string               `json:"readme"`
	Maintainers    []*User              `json:"maintainers"`
	Time           map[string]time.Time `json:"time"`
	Homepage       string               `json:"homepage"`
	Keywords       []string             `json:"keywords"`
	Repository     *Repository          `json:"repository"`
	Author         *Author              `json:"author"`
	Bugs           *Bugs                `json:"bugs"`
	License        string               `json:"license"`
	ReadmeFilename string               `json:"readmeFilename"`
	Users          map[string]bool      `json:"users"`
	Contributors   []*Contributor       `json:"contributors"`
}
