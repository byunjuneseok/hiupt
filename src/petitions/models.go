package petitions

type Petition struct {
	HashKey  string `json:"hash_key"`
	Title    string `json:"title"`
	AuthorId string `json:"author_id"`
	Contents string `json:"contents"`
}