package jsonkit_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/umefy/godash/jsonkit"
	pb "github.com/umefy/godash/jsonkit/testdata"
)

type JsonSuite struct {
	suite.Suite
}

func (s *JsonSuite) TestMarshal_Success() {
	user := &pb.User{Name: "John", Age: 30, City: "New York"}

	data, err := jsonkit.Marshal(user)

	s.Nil(err)
	s.JSONEq(`{"name":"John","age":30,"city":"New York"}`+"\n", string(data))
}

func (s *JsonSuite) TestMarshal_Error() {
	// Create a channel which cannot be marshaled to JSON
	ch := make(chan int)

	_, err := jsonkit.Marshal(ch)

	s.NotNil(err)
}

func (s *JsonSuite) TestUnMarshal_Success() {
	jsonStr := `{"name": "John", "age": 30, "city": "New York"}`

	var user pb.User
	err := jsonkit.UnMarshal([]byte(jsonStr), &user)

	s.Nil(err)
	s.Equal("John", user.Name)
	s.Equal(int32(30), user.Age)
	s.Equal("New York", user.City)
}

func (s *JsonSuite) TestUnMarshal_ExtraFields_Error() {
	jsonStr := `{"name": "John", "age": 30, "city": "New York", "extra": "field"}`

	var user pb.User
	err := jsonkit.UnMarshal([]byte(jsonStr), &user)

	s.NotNil(err)
}

func (s *JsonSuite) TestUnMarshal_InvalidJson_Error() {
	jsonStr := `{"name": "John", "age": 30, "city": "New York" invalid`

	var user pb.User
	err := jsonkit.UnMarshal([]byte(jsonStr), &user)

	s.NotNil(err)
}

func (s *JsonSuite) TestMarshalProto_Success() {
	user := &pb.User{Name: "John", Age: 30, City: "New York"}

	data, err := jsonkit.MarshalProto(user)

	s.Nil(err)
	s.JSONEq(`{"name":"John","age":30,"city":"New York"}`, string(data))
}

func (s *JsonSuite) TestUnMarshalProto_Success() {
	jsonStr := `{"name": "John", "age": 30, "city": "New York"}`

	var user pb.User
	err := jsonkit.UnMarshalProto([]byte(jsonStr), &user)

	s.Nil(err)
	s.Equal("John", user.Name)
	s.Equal(int32(30), user.Age)
	s.Equal("New York", user.City)
}

func (s *JsonSuite) TestUnMarshalProto_InvalidJson_Error() {
	jsonStr := `{"name": "John", "age": 30, "city": "New York" invalid`

	var user pb.User
	err := jsonkit.UnMarshalProto([]byte(jsonStr), &user)

	s.NotNil(err)
}

func (s *JsonSuite) TestBindRequestBody_Success() {
	jsonStr := `{"name": "John", "age": 30, "city": "New York"}`

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Body = io.NopCloser(strings.NewReader(jsonStr))
	var user pb.User
	err := jsonkit.BindRequestBody(req, &user)
	s.Nil(err)
	s.Equal("John", user.Name)
	s.Equal(int32(30), user.Age)
	s.Equal("New York", user.City)
}

func (s *JsonSuite) TestBindRequestBody_ExtraBody_Error() {
	jsonStr := `{"name": "John", "age": 30, "city": "New York", "extra": "extra"}`
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Body = io.NopCloser(strings.NewReader(jsonStr))
	var user pb.User
	err := jsonkit.BindRequestBody(req, &user)
	s.NotNil(err)
}

func (s *JsonSuite) TestBindRequestBody_InvalidJson_Error() {
	jsonStr := `{"name": "John", "age": 30, "city": "New York"} {"name": "John", "age": 30, "city": "New York"}`
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Body = io.NopCloser(strings.NewReader(jsonStr))
	var user pb.User
	err := jsonkit.BindRequestBody(req, &user)
	s.NotNil(err)
}

func (s *JsonSuite) TestJSONResponse_Success() {
	w := httptest.NewRecorder()
	user := &pb.User{Name: "John", Age: 30, City: "New York"}

	err := jsonkit.JSONResponse(w, http.StatusOK, user)

	s.Nil(err)
	s.Equal(http.StatusOK, w.Code)
	s.Equal("application/json", w.Header().Get("Content-Type"))
	s.JSONEq(`{"name":"John","age":30,"city":"New York"}`, w.Body.String())
}

func (s *JsonSuite) TestJSONResponse_InvalidData() {
	w := httptest.NewRecorder()
	// Create a channel which cannot be marshaled to JSON
	ch := make(chan int)

	err := jsonkit.JSONResponse(w, http.StatusOK, ch)

	s.NotNil(err)
}

func (s *JsonSuite) TestBindProtoRequestBody_Success() {
	jsonStr := `{"name": "John", "age": 30}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(jsonStr))

	msg := &pb.User{}
	err := jsonkit.BindProtoRequestBody(req, msg)

	s.Nil(err)
	// Verify body can be read again
	body, err := io.ReadAll(req.Body)
	s.Nil(err)
	s.Equal(jsonStr, string(body))
}

func (s *JsonSuite) TestBindProtoRequestBody_InvalidJSON() {
	jsonStr := `{"name": "John", "age": 30` // Invalid JSON
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(jsonStr))

	msg := &pb.User{}
	err := jsonkit.BindProtoRequestBody(req, msg)

	s.NotNil(err)
}

func (s *JsonSuite) TestProtoJSONResponse_Success() {
	w := httptest.NewRecorder()
	msg := &pb.User{
		Name: "John",
		Age:  0,
	}

	err := jsonkit.ProtoJSONResponse(w, http.StatusCreated, msg)

	s.Nil(err)
	s.Equal(http.StatusCreated, w.Code)
	s.Equal("application/json", w.Header().Get("Content-Type"))
	// Note: actual JSON structure will depend on your protobuf definition
}

func TestJsonSuite(t *testing.T) {
	suite.Run(t, new(JsonSuite))
}
