package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cheekybits/is"
	"github.com/gorilla/context"
	app "fairlance.io/application"
)

func TestFreelancerHandler(t *testing.T) {
	is := is.New(t)
	requestBody := `
	{
	  "password": "123",
	  "email": "pera@gmail.com",
	  "firstName":"Pera",
	  "lastName":"Peric",
	  "title":"dev",
	  "hourlyRateFrom": 12,
	  "hourlyRateTo": 22,
	  "timeZone": "CET"
	}`

	w := httptest.NewRecorder()
	r := getRequest("GET", requestBody)
	emptyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	app.FreelancerHandler(emptyHandler).ServeHTTP(w, r)
	freelancer := context.Get(r, "freelancer").(*app.Freelancer)
	is.Equal(freelancer.FirstName, "Pera")
	is.Equal(freelancer.LastName, "Peric")
	is.Equal(freelancer.Email, "pera@gmail.com")
	is.Equal(freelancer.Title, "dev")
	is.Equal(freelancer.HourlyRateFrom, 12)
	is.Equal(freelancer.HourlyRateTo, 22)
	is.Equal(freelancer.TimeZone, "CET")
}

func TestFreelancerHandlerWithInvalidBody(t *testing.T) {
	is := is.New(t)
	requestBody := `
	{
		"empty": "invalid body"
	}`

	w := httptest.NewRecorder()
	r := getRequest("GET", requestBody)
	app.FreelancerHandler(emptyHandler).ServeHTTP(w, r)
	is.Equal(w.Code, http.StatusBadRequest)
	var errorBody map[string]string
	is.NoErr(json.Unmarshal(w.Body.Bytes(), &errorBody))
	is.OK(errorBody["Email"])
	is.OK(errorBody["FirstName"])
	is.OK(errorBody["LastName"])
	is.OK(errorBody["HourlyRateFrom"])
	is.OK(errorBody["HourlyRateTo"])
	is.OK(errorBody["Password"])
	is.OK(errorBody["TimeZone"])
	is.OK(errorBody["Title"])
}

func TestFreelancerHandlerWithInvalidEmail(t *testing.T) {
	is := is.New(t)
	requestBody := `
	{
		"email": "invalid email",
	  "password": "123",
	  "firstName":"Pera",
	  "lastName":"Peric",
	  "title":"dev",
	  "hourlyRateFrom": 12,
	  "hourlyRateTo": 22,
	  "timeZone": "CET"
	}`

	w := httptest.NewRecorder()
	r := getRequest("GET", requestBody)
	app.FreelancerHandler(emptyHandler).ServeHTTP(w, r)
	is.Equal(w.Code, http.StatusBadRequest)
	var body map[string]string
	is.NoErr(json.Unmarshal(w.Body.Bytes(), &body))
	is.OK(body["Email"])
}

func TestFreelancerReviewHandler(t *testing.T) {
	is := is.New(t)
	requestBody := `
	{
		"title":        "tetetetetet",
		"content":      "content",
		"rating":       2.4,
		"clientId":     2,
		"freelancerId": 12
	}`

	w := httptest.NewRecorder()
	r := getRequest("GET", requestBody)
	app.FreelancerReviewHandler(emptyHandler).ServeHTTP(w, r)
	review := context.Get(r, "review").(*app.Review)
	is.Equal(review.ClientId, 2)
	is.Equal(review.Content, "content")
	is.Equal(review.Rating, 2.4)
	is.Equal(review.Title, "tetetetetet")
	is.Equal(review.FreelancerId, 12)
}

func TestFreelancerReferenceHandler(t *testing.T) {
	is := is.New(t)
	requestBody := `
	{
	  "title": "ttttt",
		"content": "ccccc",
		"media": {
			"image": "i",
			"video": "v"
		}
	}`

	w := httptest.NewRecorder()
	r := getRequest("GET", requestBody)
	app.FreelancerReferenceHandler(emptyHandler).ServeHTTP(w, r)
	reference := context.Get(r, "reference").(*app.Reference)
	is.Equal(reference.Title, "ttttt")
	is.Equal(reference.Content, "ccccc")
	is.Equal(reference.Media.Image, "i")
	is.Equal(reference.Media.Video, "v")
}