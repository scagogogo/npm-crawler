package model

type Version struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Description     string            `json:"description"`
	Main            string            `json:"main"`
	Scripts         *Script           `json:"scripts"`
	Repository      *Repository       `json:"repository"`
	Keywords        []string          `json:"keywords"`
	Author          *User             `json:"author"`
	License         string            `json:"license"`
	Bugs            *Bugs             `json:"bugs"`
	Homepage        string            `json:"homepage"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	ID              string            `json:"_id"`
	Dist            *Dist             `json:"dist"`
	From            string            `json:"_from"`
	NpmVersion      string            `json:"_npmVersion"`
	NpmUser         *User             `json:"_npmUser"`
	Maintainers     []*User           `json:"maintainers"`

	// TODO 2023-5-26 04:22:33 这个里面不知道是啥
	Directories struct {
	} `json:"directories"`

	Deprecated string `json:"deprecated"`
}
