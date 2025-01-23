package jsonkit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func UnMarshal(data []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields() // Prevents extra unknown fields

	// Decode the JSON into the target struct
	if err := decoder.Decode(v); err != nil {
		return err
	}

	// 🚨 Check for leftover data
	if decoder.More() {
		return fmt.Errorf("unexpected extra JSON data found")
	}

	return nil
}

func MarshalProto(v proto.Message) ([]byte, error) {
	mashaler := protojson.MarshalOptions{
		EmitUnpopulated:   true,
		EmitDefaultValues: false,
	}
	return mashaler.Marshal(v)
}

func UnMarshalProto(data []byte, v proto.Message) error {
	return protojson.Unmarshal(data, v)
}

func BindRequestBody(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Prevents extra unknown fields

	// Decode the JSON into the target struct
	if err := decoder.Decode(v); err != nil {
		return err
	}

	// 🚨 Check for leftover data
	if decoder.More() {
		return fmt.Errorf("unexpected extra JSON data found")
	}

	return nil
}

// JSONResponse writes a Go struct as JSON to the response.
func JSONResponse(w http.ResponseWriter, statusCode int, v interface{}) error {
	bytes, err := Marshal(v)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(bytes)
	return err
}

// BindProtoRequestBody decodes JSON from the request body into a Protobuf message.
func BindProtoRequestBody(r *http.Request, v proto.Message) error {
	data, err := io.ReadAll(r.Body) // Read entire body
	if err != nil {
		return err
	}

	// Restore the request body so it can be read again later
	r.Body = io.NopCloser(bytes.NewReader(data))

	// Unmarshal Protobuf JSON
	return UnMarshalProto(data, v)
}

// ProtoJSONResponse writes a Protobuf message as JSON to the response.
func ProtoJSONResponse(w http.ResponseWriter, statusCode int, v proto.Message) error {
	data, err := MarshalProto(v)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(data)
	return err
}
