// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package data

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

func easyjson794297d0DecodeGithubComAbhi2109TodoAPIData(in *jlexer.Lexer, out *UserArray) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(UserArray, 0, 2)
			} else {
				*out = UserArray{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 User
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson794297d0EncodeGithubComAbhi2109TodoAPIData(out *jwriter.Writer, in UserArray) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v UserArray) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson794297d0EncodeGithubComAbhi2109TodoAPIData(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserArray) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeGithubComAbhi2109TodoAPIData(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserArray) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson794297d0DecodeGithubComAbhi2109TodoAPIData(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserArray) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeGithubComAbhi2109TodoAPIData(l, v)
}
func easyjson794297d0DecodeGithubComAbhi2109TodoAPIData1(in *jlexer.Lexer, out *User) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Id":
			out.Id = int(in.Int())
		case "Name":
			out.Name = string(in.String())
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
func easyjson794297d0EncodeGithubComAbhi2109TodoAPIData1(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"Name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson794297d0EncodeGithubComAbhi2109TodoAPIData1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeGithubComAbhi2109TodoAPIData1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson794297d0DecodeGithubComAbhi2109TodoAPIData1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeGithubComAbhi2109TodoAPIData1(l, v)
}
func easyjson794297d0DecodeGithubComAbhi2109TodoAPIData2(in *jlexer.Lexer, out *TodoArray) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(TodoArray, 0, 1)
			} else {
				*out = TodoArray{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 Todo
			(v4).UnmarshalEasyJSON(in)
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson794297d0EncodeGithubComAbhi2109TodoAPIData2(out *jwriter.Writer, in TodoArray) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			(v6).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v TodoArray) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson794297d0EncodeGithubComAbhi2109TodoAPIData2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TodoArray) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeGithubComAbhi2109TodoAPIData2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TodoArray) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson794297d0DecodeGithubComAbhi2109TodoAPIData2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TodoArray) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeGithubComAbhi2109TodoAPIData2(l, v)
}
func easyjson794297d0DecodeGithubComAbhi2109TodoAPIData3(in *jlexer.Lexer, out *Todo) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Id":
			out.Id = int(in.Int())
		case "Name":
			out.Name = string(in.String())
		case "Desciption":
			out.Description = string(in.String())
		case "Completed":
			out.Completed = bool(in.Bool())
		case "UserId":
			out.UserId = int(in.Int())
		case "StartTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.StartTime).UnmarshalJSON(data))
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
func easyjson794297d0EncodeGithubComAbhi2109TodoAPIData3(out *jwriter.Writer, in Todo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"Name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Desciption\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"Completed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Completed))
	}
	{
		const prefix string = ",\"UserId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.UserId))
	}
	{
		const prefix string = ",\"StartTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.StartTime).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Todo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson794297d0EncodeGithubComAbhi2109TodoAPIData3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Todo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeGithubComAbhi2109TodoAPIData3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Todo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson794297d0DecodeGithubComAbhi2109TodoAPIData3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Todo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeGithubComAbhi2109TodoAPIData3(l, v)
}
func easyjson794297d0DecodeGithubComAbhi2109TodoAPIData4(in *jlexer.Lexer, out *Error) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Errorname":
			out.Errorname = string(in.String())
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
func easyjson794297d0EncodeGithubComAbhi2109TodoAPIData4(out *jwriter.Writer, in Error) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Errorname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Errorname))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Error) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson794297d0EncodeGithubComAbhi2109TodoAPIData4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Error) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeGithubComAbhi2109TodoAPIData4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Error) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson794297d0DecodeGithubComAbhi2109TodoAPIData4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Error) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeGithubComAbhi2109TodoAPIData4(l, v)
}