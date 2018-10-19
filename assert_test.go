package assert

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type (
	AssertTestSuite struct {
		BaseTestSuite
	}
)

func (s *AssertTestSuite) TestRequestMediaTypeWithInvalidPath() {
	err := RequestMediaType("application/json", s.doc, "/pet", http.MethodPost)

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrResourceURI)
}

func (s *AssertTestSuite) TestRequestMediaTypeWithInvalidType() {
	err := RequestMediaType("text/html", s.doc, "/api/food", http.MethodGet)

	s.assert.Error(err)
	s.assert.EqualError(err, fmt.Sprintf(FailMessage, fmt.Sprintf(ErrMediaType, "text/html", "application/json")))
}

func (s *AssertTestSuite) TestRequestMediaTypeWithValidType() {
	err := RequestMediaType("application/json", s.doc, "/api/food", http.MethodGet)

	s.assert.Nil(err)
}

func (s *AssertTestSuite) TestResponseMediaTypeWithInvalidPath() {
	err := ResponseMediaType("application/json", s.doc, "/pet", http.MethodPost)

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrResourceURI)
}

func (s *AssertTestSuite) TestResponseMediaTypeWithInvalidType() {
	err := ResponseMediaType("text/html", s.doc, "/api/food", http.MethodGet)

	s.assert.Error(err)
	s.assert.EqualError(err, fmt.Sprintf(FailMessage, fmt.Sprintf(ErrMediaType, "text/html", "application/json")))
}

func (s *AssertTestSuite) TestResponseMediaTypeWithValidType() {
	err := ResponseMediaType("application/json", s.doc, "/api/food", http.MethodGet)

	s.assert.Nil(err)
}

func TestAssertTestSuite(t *testing.T) {
	suite.Run(t, new(AssertTestSuite))
}
