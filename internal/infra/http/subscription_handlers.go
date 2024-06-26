package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/agmmtoo/lib-manage/pkg/libraryapp/config"
)

func (h *LibraryAppHandler) ListSubscriptions(w http.ResponseWriter, r *http.Request) error {

	qLimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(qLimit)
	if err != nil {
		limit = config.API_DEFAULT_LIMIT
	}

	qSkip := r.URL.Query().Get("skip")
	skip, err := strconv.Atoi(qSkip)
	if err != nil {
		skip = config.API_DEFAULT_SKIP
	}

	var ids []int
	if qIDs := r.URL.Query().Get("ids"); qIDs != "" {
		for _, id := range strings.Split(qIDs, ",") {
			i, err := strconv.Atoi(id)
			if err != nil {
				return InvalidRequestData(map[string]string{"ids": "invalid"})
			}
			ids = append(ids, i)
		}
	}

	var userIDs []int
	if qIDs := r.URL.Query().Get("user_ids"); qIDs != "" {
		for _, id := range strings.Split(qIDs, ",") {
			i, err := strconv.Atoi(id)
			if err != nil {
				return InvalidRequestData(map[string]string{"user_ids": "invalid"})
			}
			userIDs = append(userIDs, i)
		}
	}

	var msIDs []int
	if qIDs := r.URL.Query().Get("membership_ids"); qIDs != "" {
		for _, id := range strings.Split(qIDs, ",") {
			i, err := strconv.Atoi(id)
			if err != nil {
				return InvalidRequestData(map[string]string{"membership_ids": "invalid"})
			}
			msIDs = append(msIDs, i)
		}
	}

	var expiryDate *time.Time
	if qDate := r.URL.Query().Get("expiry_date"); qDate != "" {
		t, err := time.Parse(time.RFC3339, qDate)
		if err != nil {
			return InvalidRequestData(map[string]string{"expiry_date": "invalid"})
		}
		expiryDate = &t
	}

	var expiredBefore *time.Time
	if qDate := r.URL.Query().Get("expired_before"); qDate != "" {
		t, err := time.Parse(time.RFC3339, qDate)
		if err != nil {
			return InvalidRequestData(map[string]string{"expired_before": "invalid"})
		}
		expiredBefore = &t
	}

	var expiredAfter *time.Time
	if qDate := r.URL.Query().Get("expired_after"); qDate != "" {
		t, err := time.Parse(time.RFC3339, qDate)
		if err != nil {
			return InvalidRequestData(map[string]string{"expired_after": "invalid"})
		}
		expiredAfter = &t
	}

	var libIDs []int
	if qIDs := r.URL.Query().Get("library_ids"); qIDs != "" {
		for _, id := range strings.Split(qIDs, ",") {
			i, err := strconv.Atoi(id)
			if err != nil {
				return InvalidRequestData(map[string]string{"library_ids": "invalid"})
			}
			libIDs = append(libIDs, i)
		}
	}

	subscriptions, err := h.service.ListSubscriptions(r.Context(), ListSubscriptionsRequest{
		IDs:           ids,
		UserIDs:       userIDs,
		MembershipIDs: msIDs,
		ExpiryDate:    expiryDate,
		Skip:          skip,
		Limit:         limit,
		LibraryIDs:    libIDs,
		ExpiredBefore: expiredBefore,
		ExpiredAfter:  expiredAfter,
	})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, subscriptions)
}

func (h *LibraryAppHandler) CreateSubscription(w http.ResponseWriter, r *http.Request) error {
	var req CreateSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON(err)
	}
	defer r.Body.Close()

	subscription, err := h.service.CreateSubscription(r.Context(), req)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, subscription)
}
