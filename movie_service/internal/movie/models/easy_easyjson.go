// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels(in *jlexer.Lexer, out *SearchMovieQuery) {
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
		case "query":
			(out.MovieQuery).UnmarshalEasyJSON(in)
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels(out *jwriter.Writer, in SearchMovieQuery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"query\":"
		out.RawString(prefix[1:])
		(in.MovieQuery).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchMovieQuery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchMovieQuery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchMovieQuery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchMovieQuery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels1(in *jlexer.Lexer, out *SearchActorQuery) {
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
		case "query":
			(out.ActorQuery).UnmarshalEasyJSON(in)
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels1(out *jwriter.Writer, in SearchActorQuery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"query\":"
		out.RawString(prefix[1:])
		(in.ActorQuery).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchActorQuery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchActorQuery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchActorQuery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchActorQuery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels1(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels2(in *jlexer.Lexer, out *MovieSource) {
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
		case "title":
			out.Title = string(in.String())
		case "rating":
			out.Rating = float32(in.Float32())
		case "album_url":
			out.AlbumURL = string(in.String())
		case "card_url":
			out.CardURL = string(in.String())
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels2(out *jwriter.Writer, in MovieSource) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Float32(float32(in.Rating))
	}
	{
		const prefix string = ",\"album_url\":"
		out.RawString(prefix)
		out.String(string(in.AlbumURL))
	}
	{
		const prefix string = ",\"card_url\":"
		out.RawString(prefix)
		out.String(string(in.CardURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MovieSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MovieSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MovieSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MovieSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels2(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels3(in *jlexer.Lexer, out *MovieQuery) {
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
		case "match_phrase_prefix":
			(out.MatchMoviePhrasePrefix).UnmarshalEasyJSON(in)
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels3(out *jwriter.Writer, in MovieQuery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"match_phrase_prefix\":"
		out.RawString(prefix[1:])
		(in.MatchMoviePhrasePrefix).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MovieQuery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MovieQuery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MovieQuery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MovieQuery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels3(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels4(in *jlexer.Lexer, out *MovieHits) {
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
		case "hits":
			if in.IsNull() {
				in.Skip()
				out.MovieHits = nil
			} else {
				in.Delim('[')
				if out.MovieHits == nil {
					if !in.IsDelim(']') {
						out.MovieHits = make([]MovieHit, 0, 0)
					} else {
						out.MovieHits = []MovieHit{}
					}
				} else {
					out.MovieHits = (out.MovieHits)[:0]
				}
				for !in.IsDelim(']') {
					var v1 MovieHit
					(v1).UnmarshalEasyJSON(in)
					out.MovieHits = append(out.MovieHits, v1)
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels4(out *jwriter.Writer, in MovieHits) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"hits\":"
		out.RawString(prefix[1:])
		if in.MovieHits == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.MovieHits {
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
func (v MovieHits) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MovieHits) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MovieHits) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MovieHits) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels4(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels5(in *jlexer.Lexer, out *MovieHit) {
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
		case "_source":
			(out.MovieSource).UnmarshalEasyJSON(in)
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels5(out *jwriter.Writer, in MovieHit) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"_source\":"
		out.RawString(prefix[1:])
		(in.MovieSource).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MovieHit) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MovieHit) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MovieHit) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MovieHit) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels5(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels6(in *jlexer.Lexer, out *MovieESResponse) {
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
		case "hits":
			(out.MovieHits).UnmarshalEasyJSON(in)
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels6(out *jwriter.Writer, in MovieESResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"hits\":"
		out.RawString(prefix[1:])
		(in.MovieHits).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MovieESResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MovieESResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MovieESResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MovieESResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels6(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels7(in *jlexer.Lexer, out *MatchMoviePhrasePrefix) {
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
		case "title":
			out.Title = string(in.String())
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels7(out *jwriter.Writer, in MatchMoviePhrasePrefix) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MatchMoviePhrasePrefix) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MatchMoviePhrasePrefix) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MatchMoviePhrasePrefix) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MatchMoviePhrasePrefix) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels7(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels8(in *jlexer.Lexer, out *MatchActorPhrasePrefix) {
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
		case "full_name":
			out.FullName = string(in.String())
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels8(out *jwriter.Writer, in MatchActorPhrasePrefix) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"full_name\":"
		out.RawString(prefix[1:])
		out.String(string(in.FullName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MatchActorPhrasePrefix) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MatchActorPhrasePrefix) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MatchActorPhrasePrefix) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MatchActorPhrasePrefix) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels8(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels9(in *jlexer.Lexer, out *ActorSource) {
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
		case "full_name":
			out.Name = string(in.String())
		case "photo_big_url":
			out.PhotoBigURL = string(in.String())
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels9(out *jwriter.Writer, in ActorSource) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"full_name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"photo_big_url\":"
		out.RawString(prefix)
		out.String(string(in.PhotoBigURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ActorSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ActorSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ActorSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ActorSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels9(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels10(in *jlexer.Lexer, out *ActorQuery) {
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
		case "match_phrase_prefix":
			(out.MatchActorPhrasePrefix).UnmarshalEasyJSON(in)
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels10(out *jwriter.Writer, in ActorQuery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"match_phrase_prefix\":"
		out.RawString(prefix[1:])
		(in.MatchActorPhrasePrefix).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ActorQuery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ActorQuery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ActorQuery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ActorQuery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels10(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels11(in *jlexer.Lexer, out *ActorHits) {
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
		case "hits":
			if in.IsNull() {
				in.Skip()
				out.ActorHits = nil
			} else {
				in.Delim('[')
				if out.ActorHits == nil {
					if !in.IsDelim(']') {
						out.ActorHits = make([]ActorHit, 0, 1)
					} else {
						out.ActorHits = []ActorHit{}
					}
				} else {
					out.ActorHits = (out.ActorHits)[:0]
				}
				for !in.IsDelim(']') {
					var v4 ActorHit
					(v4).UnmarshalEasyJSON(in)
					out.ActorHits = append(out.ActorHits, v4)
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels11(out *jwriter.Writer, in ActorHits) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"hits\":"
		out.RawString(prefix[1:])
		if in.ActorHits == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.ActorHits {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ActorHits) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ActorHits) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ActorHits) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ActorHits) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels11(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels12(in *jlexer.Lexer, out *ActorHit) {
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
		case "_source":
			(out.ActorSource).UnmarshalEasyJSON(in)
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels12(out *jwriter.Writer, in ActorHit) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"_source\":"
		out.RawString(prefix[1:])
		(in.ActorSource).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ActorHit) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ActorHit) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ActorHit) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ActorHit) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels12(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels13(in *jlexer.Lexer, out *ActorESResponse) {
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
		case "hits":
			(out.ActorHits).UnmarshalEasyJSON(in)
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels13(out *jwriter.Writer, in ActorESResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"hits\":"
		out.RawString(prefix[1:])
		(in.ActorHits).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ActorESResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ActorESResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ActorESResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ActorESResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSMovieServiceInternalMovieModels13(l, v)
}
