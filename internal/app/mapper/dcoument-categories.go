package mapper

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponseAllDocumentCategories(d []domain.DocumentCategory) (*responses.GetAllDocumentCategories, error) {
	categories := make([]responses.Category, 0, len(d))
	for _, v := range d {
		categories = append(categories,
			responses.Category{
				ID:   v.ID,
				Name: v.Name,
			})
	}

	return &responses.GetAllDocumentCategories{Categories: categories}, nil
}

func MakeResponseDocumentCategory(v *domain.DocumentCategory) (*responses.GetDocumentCategory, error) {
	return &responses.GetDocumentCategory{
		ID:   v.ID,
		Name: v.Name,
	}, nil
}

func MakeRequestPostDocumentCategory(v *requests.PostDocumentCategory) *domain.DocumentCategory {
	return &domain.DocumentCategory{Name: v.Name}
}

func MakeRequestUpdateDocumentCategory(v *requests.UpdateDocumentCategory) *domain.DocumentCategory {
	return &domain.DocumentCategory{ID: v.ID, Name: v.Name}
}
