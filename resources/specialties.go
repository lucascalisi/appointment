package resources

type SpecialtySearcher struct {
	Category    string
	SubCategory string
}

type SpecialtyStorage interface {
	Search(SpecialtySearcher) ([]Specialty, error)
}

type Specialty struct {
	ID            int64
	Category      string
	SubCategories []SpecialtyDetails
}

type SpecialtyDetails struct {
	ID          int64
	SubCategory string
}
