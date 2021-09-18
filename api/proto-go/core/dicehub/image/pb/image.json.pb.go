// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: image.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*ImageGetRequest)(nil)
var _ json.Unmarshaler = (*ImageGetRequest)(nil)
var _ json.Marshaler = (*ImageGetResponse)(nil)
var _ json.Unmarshaler = (*ImageGetResponse)(nil)
var _ json.Marshaler = (*ImageListRequest)(nil)
var _ json.Unmarshaler = (*ImageListRequest)(nil)
var _ json.Marshaler = (*ImageListResponse)(nil)
var _ json.Unmarshaler = (*ImageListResponse)(nil)
var _ json.Marshaler = (*ImageListResponseData)(nil)
var _ json.Unmarshaler = (*ImageListResponseData)(nil)
var _ json.Marshaler = (*ImageGetResponseData)(nil)
var _ json.Unmarshaler = (*ImageGetResponseData)(nil)

// ImageGetRequest implement json.Marshaler.
func (m *ImageGetRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ImageGetRequest implement json.Marshaler.
func (m *ImageGetRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ImageGetResponse implement json.Marshaler.
func (m *ImageGetResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ImageGetResponse implement json.Marshaler.
func (m *ImageGetResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ImageListRequest implement json.Marshaler.
func (m *ImageListRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ImageListRequest implement json.Marshaler.
func (m *ImageListRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ImageListResponse implement json.Marshaler.
func (m *ImageListResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ImageListResponse implement json.Marshaler.
func (m *ImageListResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ImageListResponseData implement json.Marshaler.
func (m *ImageListResponseData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ImageListResponseData implement json.Marshaler.
func (m *ImageListResponseData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ImageGetResponseData implement json.Marshaler.
func (m *ImageGetResponseData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ImageGetResponseData implement json.Marshaler.
func (m *ImageGetResponseData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
