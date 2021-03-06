package entity

type PhotoKeywordMap map[string]PhotoKeyword

var PhotoKeywordFixtures = PhotoKeywordMap{
	"1": {
		PhotoID:   1000004,
		KeywordID: 1000000,
	},
	"2": {
		PhotoID:   1000001,
		KeywordID: 1000001,
	},
}

// CreatePhotoKeywordFixtures inserts known entities into the database for testing.
func CreatePhotoKeywordFixtures() {
	for _, entity := range PhotoKeywordFixtures {
		Db().Create(&entity)
	}
}
