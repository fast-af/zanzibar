// Code generated by zanzibar
// @generated
// Checksum : YvdAjy5EIbS+irqOGwyacg==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package authtoken

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

func easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtoken(in *jlexer.Lexer, out *Product) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var IDSet bool
	var NameSet bool
	var YearSet bool
	var ColorSet bool
	var PantoneValueSet bool
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
			out.ID = int32(in.Int32())
			IDSet = true
		case "name":
			out.Name = string(in.String())
			NameSet = true
		case "year":
			out.Year = int32(in.Int32())
			YearSet = true
		case "color":
			out.Color = string(in.String())
			ColorSet = true
		case "pantone_value":
			out.PantoneValue = string(in.String())
			PantoneValueSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !IDSet {
		in.AddError(fmt.Errorf("key 'id' is required"))
	}
	if !NameSet {
		in.AddError(fmt.Errorf("key 'name' is required"))
	}
	if !YearSet {
		in.AddError(fmt.Errorf("key 'year' is required"))
	}
	if !ColorSet {
		in.AddError(fmt.Errorf("key 'color' is required"))
	}
	if !PantoneValueSet {
		in.AddError(fmt.Errorf("key 'pantone_value' is required"))
	}
}
func easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtoken(out *jwriter.Writer, in Product) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int32(int32(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"year\":"
		out.RawString(prefix)
		out.Int32(int32(in.Year))
	}
	{
		const prefix string = ",\"color\":"
		out.RawString(prefix)
		out.String(string(in.Color))
	}
	{
		const prefix string = ",\"pantone_value\":"
		out.RawString(prefix)
		out.String(string(in.PantoneValue))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Product) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtoken(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Product) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtoken(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Product) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtoken(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Product) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtoken(l, v)
}
func easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenMultiCallsGetRandomProduct(in *jlexer.Lexer, out *MultiCalls_GetRandomProduct_Result) {
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
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(Product)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Success).UnmarshalJSON(data))
				}
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
func easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenMultiCallsGetRandomProduct(out *jwriter.Writer, in MultiCalls_GetRandomProduct_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((*in.Success).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MultiCalls_GetRandomProduct_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenMultiCallsGetRandomProduct(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MultiCalls_GetRandomProduct_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenMultiCallsGetRandomProduct(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MultiCalls_GetRandomProduct_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenMultiCallsGetRandomProduct(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MultiCalls_GetRandomProduct_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenMultiCallsGetRandomProduct(l, v)
}
func easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenMultiCallsGetRandomProduct1(in *jlexer.Lexer, out *MultiCalls_GetRandomProduct_Args) {
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
func easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenMultiCallsGetRandomProduct1(out *jwriter.Writer, in MultiCalls_GetRandomProduct_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MultiCalls_GetRandomProduct_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenMultiCallsGetRandomProduct1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MultiCalls_GetRandomProduct_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenMultiCallsGetRandomProduct1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MultiCalls_GetRandomProduct_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenMultiCallsGetRandomProduct1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MultiCalls_GetRandomProduct_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenMultiCallsGetRandomProduct1(l, v)
}
func easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenAuthTokenGetAuthToken(in *jlexer.Lexer, out *AuthToken_GetAuthToken_Result) {
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
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(AuthTokenResponse)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Success).UnmarshalJSON(data))
				}
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
func easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenAuthTokenGetAuthToken(out *jwriter.Writer, in AuthToken_GetAuthToken_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((*in.Success).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AuthToken_GetAuthToken_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenAuthTokenGetAuthToken(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AuthToken_GetAuthToken_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenAuthTokenGetAuthToken(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AuthToken_GetAuthToken_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenAuthTokenGetAuthToken(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AuthToken_GetAuthToken_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenAuthTokenGetAuthToken(l, v)
}
func easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenAuthTokenGetAuthToken1(in *jlexer.Lexer, out *AuthToken_GetAuthToken_Args) {
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
func easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenAuthTokenGetAuthToken1(out *jwriter.Writer, in AuthToken_GetAuthToken_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AuthToken_GetAuthToken_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenAuthTokenGetAuthToken1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AuthToken_GetAuthToken_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenAuthTokenGetAuthToken1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AuthToken_GetAuthToken_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenAuthTokenGetAuthToken1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AuthToken_GetAuthToken_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtokenAuthTokenGetAuthToken1(l, v)
}
func easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtoken1(in *jlexer.Lexer, out *AuthTokenResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var AccessTokenSet bool
	var ExpiresInSet bool
	var TokenTypeSet bool
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
		case "access_token":
			out.AccessToken = string(in.String())
			AccessTokenSet = true
		case "expires_in":
			out.ExpiresIn = int32(in.Int32())
			ExpiresInSet = true
		case "token_type":
			out.TokenType = string(in.String())
			TokenTypeSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !AccessTokenSet {
		in.AddError(fmt.Errorf("key 'access_token' is required"))
	}
	if !ExpiresInSet {
		in.AddError(fmt.Errorf("key 'expires_in' is required"))
	}
	if !TokenTypeSet {
		in.AddError(fmt.Errorf("key 'token_type' is required"))
	}
}
func easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtoken1(out *jwriter.Writer, in AuthTokenResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"access_token\":"
		out.RawString(prefix[1:])
		out.String(string(in.AccessToken))
	}
	{
		const prefix string = ",\"expires_in\":"
		out.RawString(prefix)
		out.Int32(int32(in.ExpiresIn))
	}
	{
		const prefix string = ",\"token_type\":"
		out.RawString(prefix)
		out.String(string(in.TokenType))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AuthTokenResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtoken1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AuthTokenResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFdd054fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtoken1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AuthTokenResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtoken1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AuthTokenResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFdd054fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsAuthtokenAuthtoken1(l, v)
}
