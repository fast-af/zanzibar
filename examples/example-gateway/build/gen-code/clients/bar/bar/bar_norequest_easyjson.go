// Code generated by zanzibar
// @generated
// Checksum : H3K3mJNJXfmWeqcM0K/GeQ==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

import (
	json "encoding/json"
	fmt "fmt"
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

func easyjson830adb78DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNoRequest(in *jlexer.Lexer, out *Bar_NoRequest_Result) {
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
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(BarResponse)
				}
				easyjson830adb78DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in, out.Success)
			}
		case "barException":
			if in.IsNull() {
				in.Skip()
				out.BarException = nil
			} else {
				if out.BarException == nil {
					out.BarException = new(BarException)
				}
				easyjson830adb78DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar1(in, out.BarException)
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
func easyjson830adb78EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNoRequest(out *jwriter.Writer, in Bar_NoRequest_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson830adb78EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out, *in.Success)
	}
	if in.BarException != nil {
		const prefix string = ",\"barException\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson830adb78EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar1(out, *in.BarException)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_NoRequest_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson830adb78EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNoRequest(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_NoRequest_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson830adb78EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNoRequest(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_NoRequest_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson830adb78DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNoRequest(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_NoRequest_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson830adb78DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNoRequest(l, v)
}
func easyjson830adb78DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar1(in *jlexer.Lexer, out *BarException) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
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
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
}
func easyjson830adb78EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar1(out *jwriter.Writer, in BarException) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"stringField\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.StringField))
	}
	out.RawByte('}')
}
func easyjson830adb78DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in *jlexer.Lexer, out *BarResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
	var IntWithRangeSet bool
	var IntWithoutRangeSet bool
	var MapIntWithRangeSet bool
	var MapIntWithoutRangeSet bool
	var BinaryFieldSet bool
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
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		case "intWithRange":
			out.IntWithRange = int32(in.Int32())
			IntWithRangeSet = true
		case "intWithoutRange":
			out.IntWithoutRange = int32(in.Int32())
			IntWithoutRangeSet = true
		case "mapIntWithRange":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MapIntWithRange = make(map[UUID]int32)
				} else {
					out.MapIntWithRange = nil
				}
				for !in.IsDelim('}') {
					key := UUID(in.String())
					in.WantColon()
					var v1 int32
					v1 = int32(in.Int32())
					(out.MapIntWithRange)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
			MapIntWithRangeSet = true
		case "mapIntWithoutRange":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MapIntWithoutRange = make(map[string]int32)
				} else {
					out.MapIntWithoutRange = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 int32
					v2 = int32(in.Int32())
					(out.MapIntWithoutRange)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
			MapIntWithoutRangeSet = true
		case "binaryField":
			if in.IsNull() {
				in.Skip()
				out.BinaryField = nil
			} else {
				out.BinaryField = in.Bytes()
			}
			BinaryFieldSet = true
		case "nextResponse":
			if in.IsNull() {
				in.Skip()
				out.NextResponse = nil
			} else {
				if out.NextResponse == nil {
					out.NextResponse = new(BarResponse)
				}
				easyjson830adb78DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in, out.NextResponse)
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
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
	if !IntWithRangeSet {
		in.AddError(fmt.Errorf("key 'intWithRange' is required"))
	}
	if !IntWithoutRangeSet {
		in.AddError(fmt.Errorf("key 'intWithoutRange' is required"))
	}
	if !MapIntWithRangeSet {
		in.AddError(fmt.Errorf("key 'mapIntWithRange' is required"))
	}
	if !MapIntWithoutRangeSet {
		in.AddError(fmt.Errorf("key 'mapIntWithoutRange' is required"))
	}
	if !BinaryFieldSet {
		in.AddError(fmt.Errorf("key 'binaryField' is required"))
	}
}
func easyjson830adb78EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out *jwriter.Writer, in BarResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"stringField\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.StringField))
	}
	{
		const prefix string = ",\"intWithRange\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.IntWithRange))
	}
	{
		const prefix string = ",\"intWithoutRange\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.IntWithoutRange))
	}
	{
		const prefix string = ",\"mapIntWithRange\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.MapIntWithRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v4First := true
			for v4Name, v4Value := range in.MapIntWithRange {
				if v4First {
					v4First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v4Name))
				out.RawByte(':')
				out.Int32(int32(v4Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"mapIntWithoutRange\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.MapIntWithoutRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v5First := true
			for v5Name, v5Value := range in.MapIntWithoutRange {
				if v5First {
					v5First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v5Name))
				out.RawByte(':')
				out.Int32(int32(v5Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"binaryField\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.BinaryField)
	}
	if in.NextResponse != nil {
		const prefix string = ",\"nextResponse\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson830adb78EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out, *in.NextResponse)
	}
	out.RawByte('}')
}
func easyjson830adb78DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNoRequest1(in *jlexer.Lexer, out *Bar_NoRequest_Args) {
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
func easyjson830adb78EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNoRequest1(out *jwriter.Writer, in Bar_NoRequest_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_NoRequest_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson830adb78EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNoRequest1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_NoRequest_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson830adb78EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNoRequest1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_NoRequest_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson830adb78DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNoRequest1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_NoRequest_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson830adb78DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarNoRequest1(l, v)
}
