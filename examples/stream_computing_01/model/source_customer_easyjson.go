// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonF58fe944DecodeGithubComMachinefiW3bstreamStreamModel(in *jlexer.Lexer, out *SourceCustomer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "firstName":
			out.FirstName = string(in.String())
		case "lastName":
			out.LastName = string(in.String())
		case "age":
			out.Age = int(in.Int())
		case "city":
			out.City = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF58fe944EncodeGithubComMachinefiW3bstreamStreamModel(out *jwriter.Writer, in SourceCustomer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"firstName\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"lastName\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"age\":"
		out.RawString(prefix)
		out.Int(int(in.Age))
	}
	{
		const prefix string = ",\"city\":"
		out.RawString(prefix)
		out.String(string(in.City))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SourceCustomer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF58fe944EncodeGithubComMachinefiW3bstreamStreamModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SourceCustomer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF58fe944EncodeGithubComMachinefiW3bstreamStreamModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SourceCustomer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF58fe944DecodeGithubComMachinefiW3bstreamStreamModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SourceCustomer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF58fe944DecodeGithubComMachinefiW3bstreamStreamModel(l, v)
}