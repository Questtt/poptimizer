package ports

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"poptimizer/data/domain"
	"testing"
	"time"
)

var (
	testReq     = httptest.NewRequest("GET", "http://localhost/abc/dfg", nil)
	testResBody = []byte("answer")
)

type TestJSONViewer struct {
	data []byte
	err  error
	id   domain.TableID
}

func (t *TestJSONViewer) ViewJSON(ctx context.Context, id domain.TableID) ([]byte, error) {
	t.id = id
	return t.data, t.err
}

func TestGoodRespond(t *testing.T) {
	w := httptest.NewRecorder()

	viewer := TestJSONViewer{data: testResBody, err: nil}
	mux := newTableMux(time.Second, &viewer)
	mux.ServeHTTP(w, testReq)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, domain.NewTableID("abc", "dfg"), viewer.id)

	body, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, testResBody, body)
}

func TestNoDoc(t *testing.T) {
	w := httptest.NewRecorder()

	viewer := TestJSONViewer{data: nil, err: mongo.ErrNoDocuments}
	mux := newTableMux(time.Second, &viewer)
	mux.ServeHTTP(w, testReq)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestServerErr(t *testing.T) {
	w := httptest.NewRecorder()

	viewer := TestJSONViewer{data: nil, err: errors.New("some")}
	mux := newTableMux(time.Second, &viewer)
	mux.ServeHTTP(w, testReq)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	body, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, []byte("some\n"), body)
}

func TestGoodSlashRedirected(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/abc/dfg/", nil)
	w := httptest.NewRecorder()

	mux := newTableMux(time.Second, &TestJSONViewer{})
	mux.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMovedPermanently, w.Code)
}