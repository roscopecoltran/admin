package admin

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/analyzer/keyword"
	"github.com/blevesearch/bleve/analysis/analyzer/simple"
	"github.com/blevesearch/bleve/analysis/lang/en"
)

// InitIndex initializes the search index at the specified path
func InitIndex(filepath string) (bleve.Index, error) {
	index, err := bleve.Open(filepath)

	// Doesn't yet exist (or error opening) so create a new one
	if err != nil {
		index, err = bleve.New(filepath, buildIndexMapping())
		if err != nil {
			return nil, err
		}
	}
	return index, nil
}

func buildIndexMapping() *bleve.IndexMapping {
	simpleTextFieldMapping := bleve.NewTextFieldMapping()
	simpleTextFieldMapping.Analyzer = simple_analyzer.Name

	englishTextFieldMapping := bleve.NewTextFieldMapping()
	englishTextFieldMapping.Analyzer = en.AnalyzerName

	keywordFieldMapping := bleve.NewTextFieldMapping()
	keywordFieldMapping.Analyzer = keyword_analyzer.Name

	mapping := bleve.NewDocumentMapping()
	mapping.AddFieldMappingsAt("Name", simpleTextFieldMapping)
	mapping.AddFieldMappingsAt("FullName", simpleTextFieldMapping)
	mapping.AddFieldMappingsAt("Description", englishTextFieldMapping)
	mapping.AddFieldMappingsAt("Language", keywordFieldMapping)
	mapping.AddFieldMappingsAt("Tags.Name", keywordFieldMapping)

	indexMapping := bleve.NewIndexMapping()
	indexMapping.AddDocumentMapping("Entity", mapping)

	return indexMapping
}
