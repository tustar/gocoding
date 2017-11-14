package omdb

//"http://www.omdbapi.com/?t=Hacker&apikey=fae7b94f"
const OmdbApiURL = "http://www.omdbapi.com/?apikey=fae7b94f&t="

type SearchResult struct {
	Title      string
	Year       string
	Released   string
	Runtime    string
	Genre      string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []*Rating
	Metascore  string
	imdbRating float32
	imdbVotes  int64
	imdbID     string
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}

type Rating struct {
	Source string
	Value  string
}
