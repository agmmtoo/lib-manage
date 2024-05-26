package setting

import (
	"context"

	"github.com/agmmtoo/lib-manage/internal/infra/http"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

type Service struct {
	store Storer
}

func New(s Storer) *Service {
	return &Service{s}
}

func (s *Service) List(ctx context.Context, input http.ListSettingsRequest) (*http.ListSettingsResponse, error) {
	result, err := s.store.ListSettings(ctx, ListRequest{
		IDs:        input.IDs,
		LibraryIDs: input.LibraryIDs,
		Limit:      input.Limit,
		Offset:     input.Skip,
		Key:        input.Key,
	})

	if err != nil {
		return nil, err
	}

	var settings []*http.Setting
	for _, st := range result.Settings {
		settings = append(settings, &http.Setting{
			ID:        st.ID,
			LibraryID: st.LibraryID,
			Key:       st.Key,
			Value:     st.Value,
			CreatedAt: st.CreatedAt,
			UpdatedAt: st.UpdatedAt,
			DeletedAt: st.DeletedAt,
		})
	}

	return &http.ListSettingsResponse{
		Data: settings,
	}, nil
}

func (s *Service) Update(ctx context.Context, input http.UpdateSettingsRequest) ([]*http.Setting, error) {
	updateReq := make([]UpdateRequest, 0, len(input))
	for _, st := range input {
		updateReq = append(updateReq, UpdateRequest{
			ID:    st.ID,
			Value: st.Value,
		})
	}
	results, err := s.store.UpdateSettingValues(ctx, updateReq)
	if err != nil {
		return nil, err
	}
	updateRes := make([]*http.Setting, 0, len(results))
	for _, st := range results {
		updateRes = append(updateRes, &http.Setting{
			ID:        st.ID,
			LibraryID: st.LibraryID,
			Key:       st.Key,
			Value:     st.Value,
			CreatedAt: st.CreatedAt,
			UpdatedAt: st.UpdatedAt,
			DeletedAt: st.DeletedAt,
		})
	}

	return updateRes, nil
}

type Storer interface {
	GetSettingValue(ctx context.Context, libraryID int, key string) (string, error)
	ListSettings(ctx context.Context, input ListRequest) (*ListResponse, error)
	UpdateSettingValues(ctx context.Context, input []UpdateRequest) ([]*libraryapp.Setting, error)
}

type ListRequest struct {
	IDs        []int
	LibraryIDs []int
	Key        string
	Limit      int
	Offset     int
}

type ListResponse struct {
	Settings []*libraryapp.Setting
	Total    int
}

type UpdateRequest struct {
	ID    int
	Value string
}
