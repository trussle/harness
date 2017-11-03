package requests_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"testing/quick"

	"github.com/trussle/harness/requests"
)

func TestGet(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		fn := func(b []byte) bool {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			})
			server := httptest.NewServer(mux)
			res := requests.Get(t, server.URL)
			return reflect.DeepEqual(b, res)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})
}

func TestPost(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		fn := func(a, b []byte) bool {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()

				input, err := ioutil.ReadAll(r.Body)
				if err != nil {
					t.Error(err)
				}

				if !reflect.DeepEqual(a, input) {
					t.Errorf("invalid post body, expected: %v, actual: %v", a, input)
				}

				w.WriteHeader(http.StatusOK)
				w.Write(b)
			})
			server := httptest.NewServer(mux)
			res := requests.Post(t, server.URL, a)
			return reflect.DeepEqual(b, res)
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})
}
