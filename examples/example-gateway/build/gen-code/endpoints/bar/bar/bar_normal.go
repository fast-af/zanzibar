// Code generated by thriftrw v1.18.0. DO NOT EDIT.
// @generated

package bar

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

// Bar_Normal_Args represents the arguments for the Bar.normal function.
//
// The arguments for normal are sent and received over the wire as this struct.
type Bar_Normal_Args struct {
	Request *BarRequest `json:"request,required"`
}

// ToWire translates a Bar_Normal_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Bar_Normal_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Request == nil {
		return w, errors.New("field Request of Bar_Normal_Args is required")
	}
	w, err = v.Request.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _BarRequest_Read(w wire.Value) (*BarRequest, error) {
	var v BarRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Bar_Normal_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_Normal_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_Normal_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_Normal_Args) FromWire(w wire.Value) error {
	var err error

	requestIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _BarRequest_Read(field.Value)
				if err != nil {
					return err
				}
				requestIsSet = true
			}
		}
	}

	if !requestIsSet {
		return errors.New("field Request of Bar_Normal_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Bar_Normal_Args
// struct.
func (v *Bar_Normal_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Request: %v", v.Request)
	i++

	return fmt.Sprintf("Bar_Normal_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_Normal_Args match the
// provided Bar_Normal_Args.
//
// This function performs a deep comparison.
func (v *Bar_Normal_Args) Equals(rhs *Bar_Normal_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !v.Request.Equals(rhs.Request) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_Normal_Args.
func (v *Bar_Normal_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	err = multierr.Append(err, enc.AddObject("request", v.Request))
	return err
}

// GetRequest returns the value of Request if it is set or its
// zero value if it is unset.
func (v *Bar_Normal_Args) GetRequest() (o *BarRequest) {
	if v != nil {
		o = v.Request
	}
	return
}

// IsSetRequest returns true if Request is not nil.
func (v *Bar_Normal_Args) IsSetRequest() bool {
	return v != nil && v.Request != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "normal" for this struct.
func (v *Bar_Normal_Args) MethodName() string {
	return "normal"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Bar_Normal_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Bar_Normal_Helper provides functions that aid in handling the
// parameters and return values of the Bar.normal
// function.
var Bar_Normal_Helper = struct {
	// Args accepts the parameters of normal in-order and returns
	// the arguments struct for the function.
	Args func(
		request *BarRequest,
	) *Bar_Normal_Args

	// IsException returns true if the given error can be thrown
	// by normal.
	//
	// An error can be thrown by normal only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for normal
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// normal into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by normal
	//
	//   value, err := normal(args)
	//   result, err := Bar_Normal_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from normal: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*BarResponse, error) (*Bar_Normal_Result, error)

	// UnwrapResponse takes the result struct for normal
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if normal threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Bar_Normal_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Bar_Normal_Result) (*BarResponse, error)
}{}

func init() {
	Bar_Normal_Helper.Args = func(
		request *BarRequest,
	) *Bar_Normal_Args {
		return &Bar_Normal_Args{
			Request: request,
		}
	}

	Bar_Normal_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BarException:
			return true
		default:
			return false
		}
	}

	Bar_Normal_Helper.WrapResponse = func(success *BarResponse, err error) (*Bar_Normal_Result, error) {
		if err == nil {
			return &Bar_Normal_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *BarException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Bar_Normal_Result.BarException")
			}
			return &Bar_Normal_Result{BarException: e}, nil
		}

		return nil, err
	}
	Bar_Normal_Helper.UnwrapResponse = func(result *Bar_Normal_Result) (success *BarResponse, err error) {
		if result.BarException != nil {
			err = result.BarException
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Bar_Normal_Result represents the result of a Bar.normal function call.
//
// The result of a normal execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Bar_Normal_Result struct {
	// Value returned by normal after a successful execution.
	Success      *BarResponse  `json:"success,omitempty"`
	BarException *BarException `json:"barException,omitempty"`
}

// ToWire translates a Bar_Normal_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Bar_Normal_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.BarException != nil {
		w, err = v.BarException.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Bar_Normal_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Bar_Normal_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_Normal_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_Normal_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_Normal_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _BarResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BarException, err = _BarException_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.BarException != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Bar_Normal_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Bar_Normal_Result
// struct.
func (v *Bar_Normal_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.BarException != nil {
		fields[i] = fmt.Sprintf("BarException: %v", v.BarException)
		i++
	}

	return fmt.Sprintf("Bar_Normal_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_Normal_Result match the
// provided Bar_Normal_Result.
//
// This function performs a deep comparison.
func (v *Bar_Normal_Result) Equals(rhs *Bar_Normal_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.BarException == nil && rhs.BarException == nil) || (v.BarException != nil && rhs.BarException != nil && v.BarException.Equals(rhs.BarException))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_Normal_Result.
func (v *Bar_Normal_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	if v.BarException != nil {
		err = multierr.Append(err, enc.AddObject("barException", v.BarException))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Bar_Normal_Result) GetSuccess() (o *BarResponse) {
	if v != nil && v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *Bar_Normal_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// GetBarException returns the value of BarException if it is set or its
// zero value if it is unset.
func (v *Bar_Normal_Result) GetBarException() (o *BarException) {
	if v != nil && v.BarException != nil {
		return v.BarException
	}

	return
}

// IsSetBarException returns true if BarException is not nil.
func (v *Bar_Normal_Result) IsSetBarException() bool {
	return v != nil && v.BarException != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "normal" for this struct.
func (v *Bar_Normal_Result) MethodName() string {
	return "normal"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Bar_Normal_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
