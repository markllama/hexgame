package types


type Game struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Author string `json:"author"`
	Copyright string `json:"copyright"`
	Description string `json:"description"`
}

func (g0 *Game) Copy(g1 Game) {
	g0.Name = g1.Name
	g0.Title = g1.Title
	g0.Author = g1.Author
	g0.Copyright = g1.Copyright
	g0.Description = g1.Description
}

