package mapper

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponseAllFAQ(f []domain.FAQ) (map[int][]responses.FAQ, error) {
	categories := make(map[int][]responses.FAQ)
	for _, v := range f {
	
		categories[v.Category_id] = append(categories[v.Category_id],
			responses.FAQ{
				ID:         	v.ID,
				Question:       v.Question,
				Answer:         v.Answer,
				Category_id:    v.Category_id,
				Club_id:        v.Club_id,
			})
	}

	return categories, nil
}

func MakeResponseFAQ(f *domain.FAQ ) (*responses.GetFAQ, error) {
	return &responses.GetFAQ{
		ID:             f.ID,
		Question:       f.Question,
		Answer:         f.Answer,
		Category_id:    f.Category_id,
		Club_id:        f.Club_id,
	}, nil
}

// func MakeResponseFAQByClubid(f []domain.FAQ, faqMediaFiles map[int]domain.MediaFile) (*responses.GetFAQByClubid, error) {
// 	faq := make([]responses.FAQ, 0, len(f))
// 	for _, v := range f {
// 		media, ok := faqMediaFiles[v.MediaID]
// 		if !ok {
// 			return nil, fmt.Errorf("can't find media for faq id %v", v.MediaID)
// 		}
// 		faq = append(faq,
// 			responses.FAQ{
// 				ID:          v.ID,
// 				Title:       v.Title,
// 				Description: v.Description,
// 				Approved:    v.Approved,
// 				Media:       media,
// 				VkPostUrl:   v.VkPostUrl,
// 				UpdatedAt:   v.UpdatedAt,
// 				CreatedAt:   v.CreatedAt,
// 				CreatedBy:   v.CreatedBy,
// 				Views:       v.Views,
// 			})
// 	}

// 	return &responses.GetFAQByClubid{FAQ: faq}, nil
// }

// func MakeRequestPostFAQ(f requests.PostFAQ) *domain.FAQ {
// 	return &domain.FAQ{
// 		Title:       f.Title,
// 		Description: f.Description,
// 		Approved:    f.Approved,
// 		MediaID:     f.MediaID,
// 		VkPostUrl:   f.VkPostUrl,
// 		UpdatedAt:   f.UpdatedAt,
// 		CreatedAt:   f.CreatedAt,
// 		CreatedBy:   f.CreatedBy,
// 		Views:       f.Views,
// 	}
// }

// func MakeRequestUpdateFAQ(f requests.UpdateFAQ) *domain.FAQ {
// 	return &domain.FAQ{
// 		ID:          f.ID,
// 		Title:       f.Title,
// 		Description: f.Description,
// 		Approved:    f.Approved,
// 		MediaID:     f.MediaID,
// 		VkPostUrl:   f.VkPostUrl,
// 		UpdatedAt:   f.UpdatedAt,
// 		CreatedAt:   f.CreatedAt,
// 		CreatedBy:   f.CreatedBy,
// 		Views:       f.Views,
// 	}
// }
