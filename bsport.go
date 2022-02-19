package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type BookingsResponse struct {
	Links struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
	} `json:"links"`
	NextPage int `json:"next_page"`
	Count    int `json:"count"`
	Results  []struct {
		Name                  string      `json:"name"`
		RecurrenceRuleBooking interface{} `json:"recurrence_rule_booking"`
		IsDiscardable         bool        `json:"is_discardable"`
		OfferDateStart        time.Time   `json:"offer_date_start"`
		OfferDurationMinute   int         `json:"offer_duration_minute"`
		Attendance            bool        `json:"attendance"`
		BookingStatusCode     int         `json:"booking_status_code"`
		Consumer              int         `json:"consumer"`
		ConsumerPaymentPack   int         `json:"consumer_payment_pack"`
		AttendanceDateUpdated interface{} `json:"attendance_date_updated"`
		Date                  time.Time   `json:"date"`
		FirstInCompany        bool        `json:"first_in_company"`
		ID                    int         `json:"id"`
		IsDeleted             bool        `json:"is_deleted"`
		Offer                 int         `json:"offer"`
		Source                int         `json:"source"`
		WasRefunded           bool        `json:"was_refunded"`
		Member                int         `json:"member"`
		CreditConsumed        int         `json:"credit_consumed"`
		Establishment         int         `json:"establishment"`
		Coach                 int         `json:"coach"`
		Level                 int         `json:"level"`
		CoachOverride         interface{} `json:"coach_override"`
		MetaActivity          int         `json:"meta_activity"`
		DateCanceled          interface{} `json:"date_canceled"`
		SpotID                interface{} `json:"spot_id"`
	} `json:"results"`
}

func getBookingsCount() (int, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.production.bsport.io/api/v1/booking/", nil)
	if err != nil {
		return 0, err
	}

	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	q.Add("page", fmt.Sprintf("%d", 1))
	q.Add("mine", "true")
	q.Add("member", memberID)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("authorization", "Token "+token)

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var bookings BookingsResponse
	err = json.NewDecoder(resp.Body).Decode(&bookings)
	if err != nil {
		return 0, err
	}

	return bookings.Count, nil

}
