package params

import (
	"net/http"
	"net/url"
	"testing"

	"reflect"
)

func newRequest(rawurl string) (*http.Request, error) {
	var req http.Request
	url, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	req.URL = url
	return &req, nil
}

func TestUnpack(t *testing.T) {
	type constraintTest struct {
		Email  string `email`
		Credit string `credit`
		Zip    string `zip`
	}
	tests := []struct {
		name    string
		input   string
		isError bool
		want    constraintTest
	}{
		{name: "empty", input: "http://localhost:12345/search", isError: false, want: constraintTest{}},
		{name: "only mail", input: "http://localhost:12345/search?email=i@tensai.com", isError: false, want: constraintTest{Email: "i@tensai.com"}},
		{name: "wrong mail no at mark", input: "http://localhost:12345/search?email=i", isError: true, want: constraintTest{}},
		{name: "wrong mail no domain", input: "http://localhost:12345/search?email=i@", isError: true, want: constraintTest{}},
		{name: "wrong mail no account", input: "http://localhost:12345/search?email=@i", isError: true, want: constraintTest{}},
		{name: "only credit", input: "http://localhost:12345/search?credit=0123456789012345", isError: false, want: constraintTest{Credit: "0123456789012345"}},
		{name: "wrong credit -1", input: "http://localhost:12345/search?credit=012345678901234", isError: true, want: constraintTest{}},
		{name: "wrong credit +1", input: "http://localhost:12345/search?credit=01234567890123456", isError: true, want: constraintTest{}},
		{name: "wrong credit alphabet", input: "http://localhost:12345/search?credit=012345678901234i", isError: true, want: constraintTest{}},
		{name: "only zip", input: "http://localhost:12345/search?zip=1234567", isError: false, want: constraintTest{Zip: "1234567"}},
		{name: "wrong zip -1", input: "http://localhost:12345/search?zip=123456", isError: true, want: constraintTest{}},
		{name: "wrong zip +1", input: "http://localhost:12345/search?zip=12345678", isError: true, want: constraintTest{}},
		{name: "wrong zip alphabet", input: "http://localhost:12345/search?zip=123456a", isError: true, want: constraintTest{}},
		{name: "mix", input: "http://localhost:12345/search?email=i@tensai.com&credit=0123456789012345&zip=1234567", isError: false, want: constraintTest{Email: "i@tensai.com", Credit: "0123456789012345", Zip: "1234567"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got constraintTest
			req, err := newRequest(test.input)
			if err != nil {
				t.Errorf("Error input %s\n", test.input)
			}
			err = Unpack(req, &got)
			if (err != nil) != test.isError {
				t.Errorf("Error input %s, error want %t, got %t", test.input, test.isError, err != nil)
			}
			if (err != nil) && !reflect.DeepEqual(test.want, got) {
				t.Errorf("Error input %s, error want %#v, got %#v", test.input, test.want, got)
			}
		})
	}
}
