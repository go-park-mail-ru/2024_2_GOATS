// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package errors

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

func easyjsonD31a5a85DecodeGithubComGoParkMailRu20242GOATSInternalAppErrors(in *jlexer.Lexer, out *ErrorItem) {
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
		case "code":
			out.Code = string(in.String())
		case "error":
			out.Error = string(in.String())
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
func easyjsonD31a5a85EncodeGithubComGoParkMailRu20242GOATSInternalAppErrors(out *jwriter.Writer, in ErrorItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix[1:])
		out.String(string(in.Code))
	}
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		out.String(string(in.Error))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD31a5a85EncodeGithubComGoParkMailRu20242GOATSInternalAppErrors(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD31a5a85EncodeGithubComGoParkMailRu20242GOATSInternalAppErrors(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD31a5a85DecodeGithubComGoParkMailRu20242GOATSInternalAppErrors(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD31a5a85DecodeGithubComGoParkMailRu20242GOATSInternalAppErrors(l, v)
}
func easyjsonD31a5a85DecodeGithubComGoParkMailRu20242GOATSInternalAppErrors1(in *jlexer.Lexer, out *DeliveryError) {
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
		case "errors":
			if in.IsNull() {
				in.Skip()
				out.Errors = nil
			} else {
				in.Delim('[')
				if out.Errors == nil {
					if !in.IsDelim(']') {
						out.Errors = make([]ErrorItem, 0, 2)
					} else {
						out.Errors = []ErrorItem{}
					}
				} else {
					out.Errors = (out.Errors)[:0]
				}
				for !in.IsDelim(']') {
					var v1 ErrorItem
					(v1).UnmarshalEasyJSON(in)
					out.Errors = append(out.Errors, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonD31a5a85EncodeGithubComGoParkMailRu20242GOATSInternalAppErrors1(out *jwriter.Writer, in DeliveryError) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"errors\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Errors == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Errors {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeliveryError) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD31a5a85EncodeGithubComGoParkMailRu20242GOATSInternalAppErrors1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeliveryError) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD31a5a85EncodeGithubComGoParkMailRu20242GOATSInternalAppErrors1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeliveryError) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD31a5a85DecodeGithubComGoParkMailRu20242GOATSInternalAppErrors1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeliveryError) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD31a5a85DecodeGithubComGoParkMailRu20242GOATSInternalAppErrors1(l, v)
}