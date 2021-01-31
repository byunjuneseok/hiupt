package petitions

type Petition struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	AuthorId string `json:"author_id"`
	Contents string `json:"contents"`
}