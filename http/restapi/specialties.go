package restapi

import (
	"github.com/appointment/http/models"
	"github.com/appointment/http/restapi/operations/specialties"
	rec "github.com/appointment/resources"
	middleware "github.com/go-openapi/runtime/middleware"
)

type specialtiesSearcher interface {
	Search(rec.SpecialtySearcher) ([]rec.Specialty, error)
}

func searchSpecialties(stg specialtiesSearcher) specialties.SearchSpecialtyHandlerFunc {
	return func(params specialties.SearchSpecialtyParams) middleware.Responder {
		specialtiesSearched, err := stg.Search(rec.SpecialtySearcher{
			Category:    *params.Category,
			SubCategory: *params.SubCategory,
		})

		if err != nil {
			return specialties.NewSearchSpecialtyInternalServerError().WithPayload(newRestApiError(err))
		}

		result := []*models.Specialty{}
		for _, s := range specialtiesSearched {
			thisSpecialty := s
			for _, sd := range thisSpecialty.SubCategories {
				thisSpecialtyDetail := sd
				var subCategory string = ""
				if thisSpecialty.Category != thisSpecialtyDetail.SubCategory {
					subCategory = thisSpecialtyDetail.SubCategory
				}
				result = append(result, &models.Specialty{
					ID:            thisSpecialty.ID,
					Category:      &thisSpecialty.Category,
					IDSubcategory: thisSpecialtyDetail.ID,
					SubCategory:   subCategory,
				})
			}
		}

		return specialties.NewSearchSpecialtyOK().WithPayload(result)
	}
}
