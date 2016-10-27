package hashmap

import "github.com/blevesearch/bleve"

type HashIndex struct {
	Path 	string
	Index 	bleve.Index
}

func NewHashIndex(path string) *HashIndex {
	return &HashIndex{
		Path: path,
		Index: NewIndex(path),
	}
}

func (im *HashIndex)AddIndex(indexValue string, value interface{})  {
	im.Index.Index(indexValue, value)
}

func (im *HashIndex)ExecQuery(queryString string)  *bleve.SearchResult{
	//Declare a search request
	stringQuery := bleve.NewQueryStringQuery(queryString)
	searchRequest := bleve.NewSearchRequest(stringQuery)

	//Execute search
	searchResult, _ := im.Index.Search(searchRequest)
	return searchResult
}


func NewIndex(path string)  bleve.Index{
	opindex, err := bleve.Open(path)
	if err != nil{
		mapping := bleve.NewIndexMapping()
		opindex, err = bleve.New(path, mapping)
	}
	return opindex
}