package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/algo", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("Got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/algo", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/algo", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Its showing does not have required field when it does have them")
	}

}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/algo", nil)
	form := New(r.PostForm)

	valid := form.Has("someField")
	if valid {
		t.Error("There is no such field")
	}

	r, _ = http.NewRequest("POST", "/algo", nil)

	postedData := url.Values{}
	postedData.Add("a", "a")

	r.PostForm = postedData
	form = New(r.PostForm)

	valid = form.Has("a")
	if !valid {
		t.Error("The field was present")
	}
}

func TestForm_isEmail(t *testing.T) {

	r := httptest.NewRequest("POST", "/algo", nil)
	form := New(r.PostForm)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("The string entered was not an email so there should exist an error")
	}

	postedData := url.Values{}
	postedData.Add("email", "test@123.com")
	form = New(postedData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("The string entered was an email so there should not exist an error")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/algo", nil)
	form := New(r.PostForm)

	form.MinLenght("field1", 3)
	if form.Valid() {
		t.Error("The field didnt had more than 3 required characters, it should not be accepted")
	}

	isError := form.Errors.Get("field1")
	if isError == "" {
		t.Error("Should have an error but did not get one")
	}

	postedData := url.Values{}
	postedData.Add("field1", "some_value")
	form = New(postedData)

	form.MinLenght("field1", 100)
	if form.Valid() {
		t.Error("The word didn't had 100 characters, it should not be accepted")
	}

	postedData = url.Values{}
	postedData.Add("field2", "ezequiel")
	form = New(postedData)

	form.MinLenght("field2", 3)
	if !form.Valid() {
		t.Error("The field had more than 3 required characters, it should not be rejected")
	}

	isError = form.Errors.Get("field2")
	if isError != "" {
		t.Error("Shouldn't have an error but got one")
	}

}
