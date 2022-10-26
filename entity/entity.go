package entity

type Relation struct {
	InfluencedBy []string `json:"influenced-by" validate:"required"`
	Influences   []string `json:"influences" validate:"required"`
}
type Language struct {
	Language       string    `json:"language" validate:"required"`
	Appeared       int       `json:"appeared" validate:"required"`
	Created        []string  `json:"created" validate:"required"`
	Functional     bool      `json:"functional"`
	ObjectOriented bool      `json:"object-oriented"`
	Relation       *Relation `json:"relation" validate:"required"`
}

func IsPalindrom(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
