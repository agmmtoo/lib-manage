package subscription

import (
	"context"
	"errors"
	"time"

	"github.com/agmmtoo/lib-manage/internal/core/models"
	"github.com/agmmtoo/lib-manage/internal/infra/http"
	am "github.com/agmmtoo/lib-manage/internal/infra/http/models"
)

type Service struct {
	repo Storer
}

func New(repo Storer) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context, input http.ListSubscriptionsRequest) (*http.ListSubscriptionsResponse, error) {
	result, err := s.repo.ListSubscriptions(ctx, ListRequest{
		IDs:           input.IDs,
		UserIDs:       input.UserIDs,
		MembershipIDs: input.MembershipIDs,
		ExpiryDate:    input.ExpiryDate,
		Limit:         input.Limit,
		Offset:        input.Skip,

		ExpiredBefore: input.ExpiredBefore,
		ExpiredAfter:  input.ExpiredAfter,
		LibraryIDs:    input.LibraryIDs,
	})
	if err != nil {
		return nil, err
	}

	var Subscriptions []*am.Subscription
	for _, m := range result.Subscriptions {
		Subscriptions = append(Subscriptions, m.ToAPIModel())
	}

	return &http.ListSubscriptionsResponse{
		Data:  Subscriptions,
		Total: result.Total,
	}, nil
}

func (s *Service) Create(ctx context.Context, input http.CreateSubscriptionRequest) (*am.Subscription, error) {

	// get subscription's membership library id
	membership, err := s.repo.GetMembershipByID(ctx, input.MembershipID)
	if err != nil {
		return nil, err
	}
	if membership == nil {
		return nil, errors.New("membership not found")
	}

	now := time.Now()

	// check if user_id has an active subscription for given library
	ac, err := s.repo.ListSubscriptions(ctx, ListRequest{
		UserIDs:      []int{input.UserID},
		LibraryIDs:   []int{membership.LibraryID},
		ExpiredAfter: &now,
	})

	if err != nil {
		return nil, err
	}

	// if ac.Total > 0 {
	if ac != nil && len(ac.Subscriptions) > 0 {
		// TODO:
		// return nil, models.ErrUserAlreadyHasActiveSubscription
		return nil, errors.New("user already has an active subscription")
	}

	// NOTE:
	// since membership can be updated,
	// we set the expiry date with the current duration value
	// get membership's expiry date
	exp := now.AddDate(0, 0, membership.DurationDays)

	result, err := s.repo.CreateSubscription(ctx, CreateRequest{
		UserID:       input.UserID,
		MembershipID: input.MembershipID,
		ExpiryDate:   exp,
	})
	if err != nil {
		return nil, err
	}

	return result.ToAPIModel(), nil
}

type Storer interface {
	ListSubscriptions(ctx context.Context, input ListRequest) (*ListResponse, error)
	// GetSubscriptionByID(ctx context.Context, id int) (*models.Subscription, error)
	CreateSubscription(ctx context.Context, input CreateRequest) (*models.Subscription, error)
	// implemented by membership service
	GetMembershipByID(ctx context.Context, id int) (*models.Membership, error)
}

type ListRequest struct {
	IDs           []int
	UserIDs       []int
	MembershipIDs []int
	ExpiryDate    *time.Time
	Limit         int
	Offset        int
	OrderBy       []struct {
		Col string
		Dir string
	}

	ExpiredBefore *time.Time
	ExpiredAfter  *time.Time
	LibraryIDs    []int
}

type ListResponse struct {
	Subscriptions []*models.Subscription
	Total         int
}

type CreateRequest struct {
	UserID       int
	MembershipID int
	ExpiryDate   time.Time
}
