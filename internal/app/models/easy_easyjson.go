// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	sql "database/sql"
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

func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels(in *jlexer.Lexer, out *Season) {
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
		case "season_number":
			out.SeasonNumber = int(in.Int())
		case "episodes":
			if in.IsNull() {
				in.Skip()
				out.Episodes = nil
			} else {
				in.Delim('[')
				if out.Episodes == nil {
					if !in.IsDelim(']') {
						out.Episodes = make([]*Episode, 0, 8)
					} else {
						out.Episodes = []*Episode{}
					}
				} else {
					out.Episodes = (out.Episodes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Episode
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Episode)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Episodes = append(out.Episodes, v1)
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels(out *jwriter.Writer, in Season) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"season_number\":"
		out.RawString(prefix[1:])
		out.Int(int(in.SeasonNumber))
	}
	{
		const prefix string = ",\"episodes\":"
		out.RawString(prefix)
		if in.Episodes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Episodes {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Season) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Season) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Season) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Season) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels1(in *jlexer.Lexer, out *MovieShortInfo) {
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
			out.ID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "card_url":
			out.CardURL = string(in.String())
		case "album_url":
			out.AlbumURL = string(in.String())
		case "rating":
			out.Rating = float32(in.Float32())
		case "release_date":
			out.ReleaseDate = string(in.String())
		case "movie_type":
			out.MovieType = string(in.String())
		case "country":
			out.Country = string(in.String())
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels1(out *jwriter.Writer, in MovieShortInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"card_url\":"
		out.RawString(prefix)
		out.String(string(in.CardURL))
	}
	{
		const prefix string = ",\"album_url\":"
		out.RawString(prefix)
		out.String(string(in.AlbumURL))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Float32(float32(in.Rating))
	}
	{
		const prefix string = ",\"release_date\":"
		out.RawString(prefix)
		out.String(string(in.ReleaseDate))
	}
	{
		const prefix string = ",\"movie_type\":"
		out.RawString(prefix)
		out.String(string(in.MovieType))
	}
	{
		const prefix string = ",\"country\":"
		out.RawString(prefix)
		out.String(string(in.Country))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MovieShortInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MovieShortInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MovieShortInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MovieShortInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels1(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels2(in *jlexer.Lexer, out *MovieInfo) {
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
			out.ID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "short_description":
			out.ShortDescription = string(in.String())
		case "full_description":
			out.FullDescription = string(in.String())
		case "card_url":
			out.CardURL = string(in.String())
		case "album_url":
			out.AlbumURL = string(in.String())
		case "title_url":
			out.TitleURL = string(in.String())
		case "rating":
			out.Rating = float32(in.Float32())
		case "release_date":
			out.ReleaseDate = string(in.String())
		case "movie_type":
			out.MovieType = string(in.String())
		case "country":
			out.Country = string(in.String())
		case "video_url":
			out.VideoURL = string(in.String())
		case "actors_info":
			if in.IsNull() {
				in.Skip()
				out.Actors = nil
			} else {
				in.Delim('[')
				if out.Actors == nil {
					if !in.IsDelim(']') {
						out.Actors = make([]*ActorInfo, 0, 8)
					} else {
						out.Actors = []*ActorInfo{}
					}
				} else {
					out.Actors = (out.Actors)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *ActorInfo
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(ActorInfo)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Actors = append(out.Actors, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "director_info":
			if in.IsNull() {
				in.Skip()
				out.Director = nil
			} else {
				if out.Director == nil {
					out.Director = new(DirectorInfo)
				}
				easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels3(in, out.Director)
			}
		case "seasons":
			if in.IsNull() {
				in.Skip()
				out.Seasons = nil
			} else {
				in.Delim('[')
				if out.Seasons == nil {
					if !in.IsDelim(']') {
						out.Seasons = make([]*Season, 0, 8)
					} else {
						out.Seasons = []*Season{}
					}
				} else {
					out.Seasons = (out.Seasons)[:0]
				}
				for !in.IsDelim(']') {
					var v5 *Season
					if in.IsNull() {
						in.Skip()
						v5 = nil
					} else {
						if v5 == nil {
							v5 = new(Season)
						}
						(*v5).UnmarshalEasyJSON(in)
					}
					out.Seasons = append(out.Seasons, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "is_favorite":
			out.IsFavorite = bool(in.Bool())
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels2(out *jwriter.Writer, in MovieInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"short_description\":"
		out.RawString(prefix)
		out.String(string(in.ShortDescription))
	}
	{
		const prefix string = ",\"full_description\":"
		out.RawString(prefix)
		out.String(string(in.FullDescription))
	}
	{
		const prefix string = ",\"card_url\":"
		out.RawString(prefix)
		out.String(string(in.CardURL))
	}
	{
		const prefix string = ",\"album_url\":"
		out.RawString(prefix)
		out.String(string(in.AlbumURL))
	}
	{
		const prefix string = ",\"title_url\":"
		out.RawString(prefix)
		out.String(string(in.TitleURL))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Float32(float32(in.Rating))
	}
	{
		const prefix string = ",\"release_date\":"
		out.RawString(prefix)
		out.String(string(in.ReleaseDate))
	}
	{
		const prefix string = ",\"movie_type\":"
		out.RawString(prefix)
		out.String(string(in.MovieType))
	}
	{
		const prefix string = ",\"country\":"
		out.RawString(prefix)
		out.String(string(in.Country))
	}
	{
		const prefix string = ",\"video_url\":"
		out.RawString(prefix)
		out.String(string(in.VideoURL))
	}
	{
		const prefix string = ",\"actors_info\":"
		out.RawString(prefix)
		if in.Actors == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Actors {
				if v6 > 0 {
					out.RawByte(',')
				}
				if v7 == nil {
					out.RawString("null")
				} else {
					(*v7).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"director_info\":"
		out.RawString(prefix)
		if in.Director == nil {
			out.RawString("null")
		} else {
			easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels3(out, *in.Director)
		}
	}
	{
		const prefix string = ",\"seasons\":"
		out.RawString(prefix)
		if in.Seasons == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Seasons {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					(*v9).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"is_favorite\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsFavorite))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MovieInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MovieInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MovieInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MovieInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels2(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels3(in *jlexer.Lexer, out *DirectorInfo) {
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
		case "ID":
			out.ID = int(in.Int())
		case "Name":
			out.Name = string(in.String())
		case "Surname":
			out.Surname = string(in.String())
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels3(out *jwriter.Writer, in DirectorInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Surname\":"
		out.RawString(prefix)
		out.String(string(in.Surname))
	}
	out.RawByte('}')
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels4(in *jlexer.Lexer, out *Episode) {
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
			out.ID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "episode_number":
			out.EpisodeNumber = int(in.Int())
		case "release_date":
			out.ReleaseDate = string(in.String())
		case "rating":
			out.Rating = float32(in.Float32())
		case "preview_url":
			out.PreviewURL = string(in.String())
		case "video_url":
			out.VideoURL = string(in.String())
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels4(out *jwriter.Writer, in Episode) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"episode_number\":"
		out.RawString(prefix)
		out.Int(int(in.EpisodeNumber))
	}
	{
		const prefix string = ",\"release_date\":"
		out.RawString(prefix)
		out.String(string(in.ReleaseDate))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Float32(float32(in.Rating))
	}
	{
		const prefix string = ",\"preview_url\":"
		out.RawString(prefix)
		out.String(string(in.PreviewURL))
	}
	{
		const prefix string = ",\"video_url\":"
		out.RawString(prefix)
		out.String(string(in.VideoURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Episode) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Episode) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Episode) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Episode) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels4(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels5(in *jlexer.Lexer, out *Collection) {
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
			out.ID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "movies":
			if in.IsNull() {
				in.Skip()
				out.Movies = nil
			} else {
				in.Delim('[')
				if out.Movies == nil {
					if !in.IsDelim(']') {
						out.Movies = make([]*MovieShortInfo, 0, 8)
					} else {
						out.Movies = []*MovieShortInfo{}
					}
				} else {
					out.Movies = (out.Movies)[:0]
				}
				for !in.IsDelim(']') {
					var v10 *MovieShortInfo
					if in.IsNull() {
						in.Skip()
						v10 = nil
					} else {
						if v10 == nil {
							v10 = new(MovieShortInfo)
						}
						(*v10).UnmarshalEasyJSON(in)
					}
					out.Movies = append(out.Movies, v10)
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels5(out *jwriter.Writer, in Collection) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"movies\":"
		out.RawString(prefix)
		if in.Movies == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Movies {
				if v11 > 0 {
					out.RawByte(',')
				}
				if v12 == nil {
					out.RawString("null")
				} else {
					(*v12).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Collection) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Collection) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Collection) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Collection) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels5(l, v)
}
func easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels6(in *jlexer.Lexer, out *ActorInfo) {
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
			out.ID = int(in.Int())
		case "biography":
			out.Biography = string(in.String())
		case "post":
			out.Post = string(in.String())
		case "birthdate":
			easyjson97766e5aDecodeDatabaseSql(in, &out.Birthdate)
		case "small_photo_url":
			out.SmallPhotoURL = string(in.String())
		case "big_photo_url":
			out.BigPhotoURL = string(in.String())
		case "country":
			out.Country = string(in.String())
		case "movies":
			if in.IsNull() {
				in.Skip()
				out.Movies = nil
			} else {
				in.Delim('[')
				if out.Movies == nil {
					if !in.IsDelim(']') {
						out.Movies = make([]*MovieShortInfo, 0, 8)
					} else {
						out.Movies = []*MovieShortInfo{}
					}
				} else {
					out.Movies = (out.Movies)[:0]
				}
				for !in.IsDelim(']') {
					var v13 *MovieShortInfo
					if in.IsNull() {
						in.Skip()
						v13 = nil
					} else {
						if v13 == nil {
							v13 = new(MovieShortInfo)
						}
						(*v13).UnmarshalEasyJSON(in)
					}
					out.Movies = append(out.Movies, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Name":
			out.Name = string(in.String())
		case "Surname":
			out.Surname = string(in.String())
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
func easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels6(out *jwriter.Writer, in ActorInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"biography\":"
		out.RawString(prefix)
		out.String(string(in.Biography))
	}
	{
		const prefix string = ",\"post\":"
		out.RawString(prefix)
		out.String(string(in.Post))
	}
	{
		const prefix string = ",\"birthdate\":"
		out.RawString(prefix)
		easyjson97766e5aEncodeDatabaseSql(out, in.Birthdate)
	}
	{
		const prefix string = ",\"small_photo_url\":"
		out.RawString(prefix)
		out.String(string(in.SmallPhotoURL))
	}
	{
		const prefix string = ",\"big_photo_url\":"
		out.RawString(prefix)
		out.String(string(in.BigPhotoURL))
	}
	{
		const prefix string = ",\"country\":"
		out.RawString(prefix)
		out.String(string(in.Country))
	}
	{
		const prefix string = ",\"movies\":"
		out.RawString(prefix)
		if in.Movies == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Movies {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					(*v15).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Surname\":"
		out.RawString(prefix)
		out.String(string(in.Surname))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ActorInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ActorInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97766e5aEncodeGithubComGoParkMailRu20242GOATSInternalAppModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ActorInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ActorInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97766e5aDecodeGithubComGoParkMailRu20242GOATSInternalAppModels6(l, v)
}
func easyjson97766e5aDecodeDatabaseSql(in *jlexer.Lexer, out *sql.NullString) {
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
		case "String":
			out.String = string(in.String())
		case "Valid":
			out.Valid = bool(in.Bool())
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
func easyjson97766e5aEncodeDatabaseSql(out *jwriter.Writer, in sql.NullString) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"String\":"
		out.RawString(prefix[1:])
		out.String(string(in.String))
	}
	{
		const prefix string = ",\"Valid\":"
		out.RawString(prefix)
		out.Bool(bool(in.Valid))
	}
	out.RawByte('}')
}
