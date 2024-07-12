package mapper

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponseAllDocuments(d []domain.Document) (*responses.GetAllDocuments, error) {
	documents := make([]responses.Document, 0, len(d))
	for _, v := range d {
		documents = append(documents,
			responses.Document{
				ID:         v.ID,
				Name:       v.Name,
				Key:        v.Key,
				ClubID:     v.ClubID,
				CategoryID: v.CategoryID,
			})
	}

	return &responses.GetAllDocuments{Documents: documents}, nil
}

func MakeResponseDocument(v *domain.Document) (*responses.GetDocument, error) {
	return &responses.GetDocument{
		ID:         v.ID,
		Name:       v.Name,
		Key:        v.Key,
		ClubID:     v.ClubID,
		CategoryID: v.CategoryID,
	}, nil
}

func MakeResponseDocumentsByClubID(d []domain.Document) (*responses.GetDocumentsByClubID, error) {
	documents := make([]responses.Document, 0, len(d))
	for _, v := range d {
		documents = append(documents,
			responses.Document{
				ID:         v.ID,
				Name:       v.Name,
				Key:        v.Key,
				ClubID:     v.ClubID,
				CategoryID: v.CategoryID,
			})
	}

	return &responses.GetDocumentsByClubID{Documents: documents}, nil
}

func MakeResponseDocumentsByCategory(d []domain.Document) (*responses.GetDocumentsByCategory, error) {
	documents := make([]responses.Document, 0, len(d))
	for _, v := range d {
		documents = append(documents,
			responses.Document{
				ID:         v.ID,
				Name:       v.Name,
				Key:        v.Key,
				ClubID:     v.ClubID,
				CategoryID: v.CategoryID,
			})
	}

	return &responses.GetDocumentsByCategory{Documents: documents}, nil
}

func MakeResponsePostDocument(key string) (*responses.PostDocument, error) {
	return &responses.PostDocument{
		NewKey: key,
	}, nil
}

func MakeResponseUpdateDocument(key string) (*responses.UpdateDocument, error) {
	return &responses.UpdateDocument{
		NewKey: key,
	}, nil
}
