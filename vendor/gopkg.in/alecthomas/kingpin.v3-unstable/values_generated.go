package kingpin

import (
	"encoding/hex"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

// This file is autogenerated by "go generate .". Do not modify.

// -- time.Duration Value
type durationValue struct{ v *time.Duration }

func newDurationValue(p *time.Duration) *durationValue {
	return &durationValue{p}
}

func (f *durationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	if err == nil {
		*f.v = (time.Duration)(v)
	}
	return err
}

func (f *durationValue) Get() interface{} { return (time.Duration)(*f.v) }

func (f *durationValue) String() string { return fmt.Sprintf("%v", *f.v) }

// Time duration.
func (p *Clause) Duration() (target *time.Duration) {
	target = new(time.Duration)
	p.DurationVar(target)
	return
}

func (p *Clause) DurationVar(target *time.Duration) {
	p.SetValue(newDurationValue(target))
}

// DurationList accumulates time.Duration values into a slice.
func (p *Clause) DurationList(options ...AccumulatorOption) (target *[]time.Duration) {
	target = new([]time.Duration)
	p.DurationListVar(target, options...)
	return
}

func (p *Clause) DurationListVar(target *[]time.Duration, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newDurationValue(v.(*time.Duration))
	}))
}

// ExistingDirs accumulates string values into a slice.
func (p *Clause) ExistingDirs(options ...AccumulatorOption) (target *[]string) {
	target = new([]string)
	p.ExistingDirsVar(target, options...)
	return
}

func (p *Clause) ExistingDirsVar(target *[]string, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newExistingDirValue(v.(*string))
	}))
}

// ExistingFiles accumulates string values into a slice.
func (p *Clause) ExistingFiles(options ...AccumulatorOption) (target *[]string) {
	target = new([]string)
	p.ExistingFilesVar(target, options...)
	return
}

func (p *Clause) ExistingFilesVar(target *[]string, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newExistingFileValue(v.(*string))
	}))
}

// ExistingFilesOrDirs accumulates string values into a slice.
func (p *Clause) ExistingFilesOrDirs(options ...AccumulatorOption) (target *[]string) {
	target = new([]string)
	p.ExistingFilesOrDirsVar(target, options...)
	return
}

func (p *Clause) ExistingFilesOrDirsVar(target *[]string, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newExistingFileOrDirValue(v.(*string))
	}))
}

// -- []byte Value
type hexBytesValue struct{ v *[]byte }

func newHexBytesValue(p *[]byte) *hexBytesValue {
	return &hexBytesValue{p}
}

func (f *hexBytesValue) Set(s string) error {
	v, err := hex.DecodeString(s)
	if err == nil {
		*f.v = ([]byte)(v)
	}
	return err
}

func (f *hexBytesValue) Get() interface{} { return ([]byte)(*f.v) }

func (f *hexBytesValue) String() string { return fmt.Sprintf("%v", *f.v) }

// Bytes as a hex string.
func (p *Clause) HexBytes() (target *[]byte) {
	target = new([]byte)
	p.HexBytesVar(target)
	return
}

func (p *Clause) HexBytesVar(target *[]byte) {
	p.SetValue(newHexBytesValue(target))
}

// HexBytesList accumulates []byte values into a slice.
func (p *Clause) HexBytesList(options ...AccumulatorOption) (target *[][]byte) {
	target = new([][]byte)
	p.HexBytesListVar(target, options...)
	return
}

func (p *Clause) HexBytesListVar(target *[][]byte, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newHexBytesValue(v.(*[]byte))
	}))
}

// IPList accumulates net.IP values into a slice.
func (p *Clause) IPList(options ...AccumulatorOption) (target *[]net.IP) {
	target = new([]net.IP)
	p.IPListVar(target, options...)
	return
}

func (p *Clause) IPListVar(target *[]net.IP, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newIPValue(v.(*net.IP))
	}))
}

// -- bool Value
type negatableBoolValue struct{ v *bool }

func newNegatableBoolValue(p *bool) *negatableBoolValue {
	return &negatableBoolValue{p}
}

func (f *negatableBoolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	if err == nil {
		*f.v = (bool)(v)
	}
	return err
}

func (f *negatableBoolValue) Get() interface{} { return (bool)(*f.v) }

func (f *negatableBoolValue) String() string { return fmt.Sprintf("%v", *f.v) }

// NegatableBool parses the next command-line value as bool.
func (p *Clause) NegatableBool() (target *bool) {
	target = new(bool)
	p.NegatableBoolVar(target)
	return
}

func (p *Clause) NegatableBoolVar(target *bool) {
	p.SetValue(newNegatableBoolValue(target))
}

// NegatableBoolList accumulates bool values into a slice.
func (p *Clause) NegatableBoolList(options ...AccumulatorOption) (target *[]bool) {
	target = new([]bool)
	p.NegatableBoolListVar(target, options...)
	return
}

func (p *Clause) NegatableBoolListVar(target *[]bool, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newNegatableBoolValue(v.(*bool))
	}))
}

// -- *regexp.Regexp Value
type regexpValue struct{ v **regexp.Regexp }

func newRegexpValue(p **regexp.Regexp) *regexpValue {
	return &regexpValue{p}
}

func (f *regexpValue) Set(s string) error {
	v, err := regexp.Compile(s)
	if err == nil {
		*f.v = (*regexp.Regexp)(v)
	}
	return err
}

func (f *regexpValue) Get() interface{} { return (*regexp.Regexp)(*f.v) }

func (f *regexpValue) String() string { return fmt.Sprintf("%v", *f.v) }

// Regexp parses the next command-line value as *regexp.Regexp.
func (p *Clause) Regexp() (target **regexp.Regexp) {
	target = new(*regexp.Regexp)
	p.RegexpVar(target)
	return
}

func (p *Clause) RegexpVar(target **regexp.Regexp) {
	p.SetValue(newRegexpValue(target))
}

// RegexpList accumulates *regexp.Regexp values into a slice.
func (p *Clause) RegexpList(options ...AccumulatorOption) (target *[]*regexp.Regexp) {
	target = new([]*regexp.Regexp)
	p.RegexpListVar(target, options...)
	return
}

func (p *Clause) RegexpListVar(target *[]*regexp.Regexp, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newRegexpValue(v.(**regexp.Regexp))
	}))
}

// -- *url.URL Value
type uRLValue struct{ v **url.URL }

func newURLValue(p **url.URL) *uRLValue {
	return &uRLValue{p}
}

func (f *uRLValue) Set(s string) error {
	v, err := url.Parse(s)
	if err == nil {
		*f.v = (*url.URL)(v)
	}
	return err
}

func (f *uRLValue) Get() interface{} { return (*url.URL)(*f.v) }

func (f *uRLValue) String() string { return fmt.Sprintf("%v", *f.v) }

// URL parses the next command-line value as *url.URL.
func (p *Clause) URL() (target **url.URL) {
	target = new(*url.URL)
	p.URLVar(target)
	return
}

func (p *Clause) URLVar(target **url.URL) {
	p.SetValue(newURLValue(target))
}

// URLList accumulates *url.URL values into a slice.
func (p *Clause) URLList(options ...AccumulatorOption) (target *[]*url.URL) {
	target = new([]*url.URL)
	p.URLListVar(target, options...)
	return
}

func (p *Clause) URLListVar(target *[]*url.URL, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newURLValue(v.(**url.URL))
	}))
}

// -- bool Value
type boolValue struct{ v *bool }

func newBoolValue(p *bool) *boolValue {
	return &boolValue{p}
}

func (f *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	if err == nil {
		*f.v = (bool)(v)
	}
	return err
}

func (f *boolValue) Get() interface{} { return (bool)(*f.v) }

func (f *boolValue) String() string { return fmt.Sprintf("%v", *f.v) }

// Bool parses the next command-line value as bool.
func (p *Clause) Bool() (target *bool) {
	target = new(bool)
	p.BoolVar(target)
	return
}

func (p *Clause) BoolVar(target *bool) {
	p.SetValue(newBoolValue(target))
}

// BoolList accumulates bool values into a slice.
func (p *Clause) BoolList(options ...AccumulatorOption) (target *[]bool) {
	target = new([]bool)
	p.BoolListVar(target, options...)
	return
}

func (p *Clause) BoolListVar(target *[]bool, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newBoolValue(v.(*bool))
	}))
}

// -- float32 Value
type float32Value struct{ v *float32 }

func newFloat32Value(p *float32) *float32Value {
	return &float32Value{p}
}

func (f *float32Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 32)
	if err == nil {
		*f.v = (float32)(v)
	}
	return err
}

func (f *float32Value) Get() interface{} { return (float32)(*f.v) }

func (f *float32Value) String() string { return fmt.Sprintf("%v", *f.v) }

// Float32 parses the next command-line value as float32.
func (p *Clause) Float32() (target *float32) {
	target = new(float32)
	p.Float32Var(target)
	return
}

func (p *Clause) Float32Var(target *float32) {
	p.SetValue(newFloat32Value(target))
}

// Float32List accumulates float32 values into a slice.
func (p *Clause) Float32List(options ...AccumulatorOption) (target *[]float32) {
	target = new([]float32)
	p.Float32ListVar(target, options...)
	return
}

func (p *Clause) Float32ListVar(target *[]float32, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newFloat32Value(v.(*float32))
	}))
}

// -- float64 Value
type float64Value struct{ v *float64 }

func newFloat64Value(p *float64) *float64Value {
	return &float64Value{p}
}

func (f *float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err == nil {
		*f.v = (float64)(v)
	}
	return err
}

func (f *float64Value) Get() interface{} { return (float64)(*f.v) }

func (f *float64Value) String() string { return fmt.Sprintf("%v", *f.v) }

// Float64 parses the next command-line value as float64.
func (p *Clause) Float64() (target *float64) {
	target = new(float64)
	p.Float64Var(target)
	return
}

func (p *Clause) Float64Var(target *float64) {
	p.SetValue(newFloat64Value(target))
}

// Float64List accumulates float64 values into a slice.
func (p *Clause) Float64List(options ...AccumulatorOption) (target *[]float64) {
	target = new([]float64)
	p.Float64ListVar(target, options...)
	return
}

func (p *Clause) Float64ListVar(target *[]float64, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newFloat64Value(v.(*float64))
	}))
}

// -- int Value
type intValue struct{ v *int }

func newIntValue(p *int) *intValue {
	return &intValue{p}
}

func (f *intValue) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err == nil {
		*f.v = (int)(v)
	}
	return err
}

func (f *intValue) Get() interface{} { return (int)(*f.v) }

func (f *intValue) String() string { return fmt.Sprintf("%v", *f.v) }

// Int parses the next command-line value as int.
func (p *Clause) Int() (target *int) {
	target = new(int)
	p.IntVar(target)
	return
}

func (p *Clause) IntVar(target *int) {
	p.SetValue(newIntValue(target))
}

// Ints accumulates int values into a slice.
func (p *Clause) Ints(options ...AccumulatorOption) (target *[]int) {
	target = new([]int)
	p.IntsVar(target, options...)
	return
}

func (p *Clause) IntsVar(target *[]int, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newIntValue(v.(*int))
	}))
}

// -- int16 Value
type int16Value struct{ v *int16 }

func newInt16Value(p *int16) *int16Value {
	return &int16Value{p}
}

func (f *int16Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 16)
	if err == nil {
		*f.v = (int16)(v)
	}
	return err
}

func (f *int16Value) Get() interface{} { return (int16)(*f.v) }

func (f *int16Value) String() string { return fmt.Sprintf("%v", *f.v) }

// Int16 parses the next command-line value as int16.
func (p *Clause) Int16() (target *int16) {
	target = new(int16)
	p.Int16Var(target)
	return
}

func (p *Clause) Int16Var(target *int16) {
	p.SetValue(newInt16Value(target))
}

// Int16List accumulates int16 values into a slice.
func (p *Clause) Int16List(options ...AccumulatorOption) (target *[]int16) {
	target = new([]int16)
	p.Int16ListVar(target, options...)
	return
}

func (p *Clause) Int16ListVar(target *[]int16, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newInt16Value(v.(*int16))
	}))
}

// -- int32 Value
type int32Value struct{ v *int32 }

func newInt32Value(p *int32) *int32Value {
	return &int32Value{p}
}

func (f *int32Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 32)
	if err == nil {
		*f.v = (int32)(v)
	}
	return err
}

func (f *int32Value) Get() interface{} { return (int32)(*f.v) }

func (f *int32Value) String() string { return fmt.Sprintf("%v", *f.v) }

// Int32 parses the next command-line value as int32.
func (p *Clause) Int32() (target *int32) {
	target = new(int32)
	p.Int32Var(target)
	return
}

func (p *Clause) Int32Var(target *int32) {
	p.SetValue(newInt32Value(target))
}

// Int32List accumulates int32 values into a slice.
func (p *Clause) Int32List(options ...AccumulatorOption) (target *[]int32) {
	target = new([]int32)
	p.Int32ListVar(target, options...)
	return
}

func (p *Clause) Int32ListVar(target *[]int32, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newInt32Value(v.(*int32))
	}))
}

// -- int64 Value
type int64Value struct{ v *int64 }

func newInt64Value(p *int64) *int64Value {
	return &int64Value{p}
}

func (f *int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	if err == nil {
		*f.v = (int64)(v)
	}
	return err
}

func (f *int64Value) Get() interface{} { return (int64)(*f.v) }

func (f *int64Value) String() string { return fmt.Sprintf("%v", *f.v) }

// Int64 parses the next command-line value as int64.
func (p *Clause) Int64() (target *int64) {
	target = new(int64)
	p.Int64Var(target)
	return
}

func (p *Clause) Int64Var(target *int64) {
	p.SetValue(newInt64Value(target))
}

// Int64List accumulates int64 values into a slice.
func (p *Clause) Int64List(options ...AccumulatorOption) (target *[]int64) {
	target = new([]int64)
	p.Int64ListVar(target, options...)
	return
}

func (p *Clause) Int64ListVar(target *[]int64, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newInt64Value(v.(*int64))
	}))
}

// -- int8 Value
type int8Value struct{ v *int8 }

func newInt8Value(p *int8) *int8Value {
	return &int8Value{p}
}

func (f *int8Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 8)
	if err == nil {
		*f.v = (int8)(v)
	}
	return err
}

func (f *int8Value) Get() interface{} { return (int8)(*f.v) }

func (f *int8Value) String() string { return fmt.Sprintf("%v", *f.v) }

// Int8 parses the next command-line value as int8.
func (p *Clause) Int8() (target *int8) {
	target = new(int8)
	p.Int8Var(target)
	return
}

func (p *Clause) Int8Var(target *int8) {
	p.SetValue(newInt8Value(target))
}

// Int8List accumulates int8 values into a slice.
func (p *Clause) Int8List(options ...AccumulatorOption) (target *[]int8) {
	target = new([]int8)
	p.Int8ListVar(target, options...)
	return
}

func (p *Clause) Int8ListVar(target *[]int8, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newInt8Value(v.(*int8))
	}))
}

// -- string Value
type stringValue struct{ v *string }

func newStringValue(p *string) *stringValue {
	return &stringValue{p}
}

func (f *stringValue) Set(s string) error {
	v, err := s, error(nil)
	if err == nil {
		*f.v = (string)(v)
	}
	return err
}

func (f *stringValue) Get() interface{} { return (string)(*f.v) }

func (f *stringValue) String() string { return string(*f.v) }

// String parses the next command-line value as string.
func (p *Clause) String() (target *string) {
	target = new(string)
	p.StringVar(target)
	return
}

func (p *Clause) StringVar(target *string) {
	p.SetValue(newStringValue(target))
}

// Strings accumulates string values into a slice.
func (p *Clause) Strings(options ...AccumulatorOption) (target *[]string) {
	target = new([]string)
	p.StringsVar(target, options...)
	return
}

func (p *Clause) StringsVar(target *[]string, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newStringValue(v.(*string))
	}))
}

// -- uint Value
type uintValue struct{ v *uint }

func newUintValue(p *uint) *uintValue {
	return &uintValue{p}
}

func (f *uintValue) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	if err == nil {
		*f.v = (uint)(v)
	}
	return err
}

func (f *uintValue) Get() interface{} { return (uint)(*f.v) }

func (f *uintValue) String() string { return fmt.Sprintf("%v", *f.v) }

// Uint parses the next command-line value as uint.
func (p *Clause) Uint() (target *uint) {
	target = new(uint)
	p.UintVar(target)
	return
}

func (p *Clause) UintVar(target *uint) {
	p.SetValue(newUintValue(target))
}

// Uints accumulates uint values into a slice.
func (p *Clause) Uints(options ...AccumulatorOption) (target *[]uint) {
	target = new([]uint)
	p.UintsVar(target, options...)
	return
}

func (p *Clause) UintsVar(target *[]uint, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newUintValue(v.(*uint))
	}))
}

// -- uint16 Value
type uint16Value struct{ v *uint16 }

func newUint16Value(p *uint16) *uint16Value {
	return &uint16Value{p}
}

func (f *uint16Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 16)
	if err == nil {
		*f.v = (uint16)(v)
	}
	return err
}

func (f *uint16Value) Get() interface{} { return (uint16)(*f.v) }

func (f *uint16Value) String() string { return fmt.Sprintf("%v", *f.v) }

// Uint16 parses the next command-line value as uint16.
func (p *Clause) Uint16() (target *uint16) {
	target = new(uint16)
	p.Uint16Var(target)
	return
}

func (p *Clause) Uint16Var(target *uint16) {
	p.SetValue(newUint16Value(target))
}

// Uint16List accumulates uint16 values into a slice.
func (p *Clause) Uint16List(options ...AccumulatorOption) (target *[]uint16) {
	target = new([]uint16)
	p.Uint16ListVar(target, options...)
	return
}

func (p *Clause) Uint16ListVar(target *[]uint16, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newUint16Value(v.(*uint16))
	}))
}

// -- uint32 Value
type uint32Value struct{ v *uint32 }

func newUint32Value(p *uint32) *uint32Value {
	return &uint32Value{p}
}

func (f *uint32Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 32)
	if err == nil {
		*f.v = (uint32)(v)
	}
	return err
}

func (f *uint32Value) Get() interface{} { return (uint32)(*f.v) }

func (f *uint32Value) String() string { return fmt.Sprintf("%v", *f.v) }

// Uint32 parses the next command-line value as uint32.
func (p *Clause) Uint32() (target *uint32) {
	target = new(uint32)
	p.Uint32Var(target)
	return
}

func (p *Clause) Uint32Var(target *uint32) {
	p.SetValue(newUint32Value(target))
}

// Uint32List accumulates uint32 values into a slice.
func (p *Clause) Uint32List(options ...AccumulatorOption) (target *[]uint32) {
	target = new([]uint32)
	p.Uint32ListVar(target, options...)
	return
}

func (p *Clause) Uint32ListVar(target *[]uint32, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newUint32Value(v.(*uint32))
	}))
}

// -- uint64 Value
type uint64Value struct{ v *uint64 }

func newUint64Value(p *uint64) *uint64Value {
	return &uint64Value{p}
}

func (f *uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	if err == nil {
		*f.v = (uint64)(v)
	}
	return err
}

func (f *uint64Value) Get() interface{} { return (uint64)(*f.v) }

func (f *uint64Value) String() string { return fmt.Sprintf("%v", *f.v) }

// Uint64 parses the next command-line value as uint64.
func (p *Clause) Uint64() (target *uint64) {
	target = new(uint64)
	p.Uint64Var(target)
	return
}

func (p *Clause) Uint64Var(target *uint64) {
	p.SetValue(newUint64Value(target))
}

// Uint64List accumulates uint64 values into a slice.
func (p *Clause) Uint64List(options ...AccumulatorOption) (target *[]uint64) {
	target = new([]uint64)
	p.Uint64ListVar(target, options...)
	return
}

func (p *Clause) Uint64ListVar(target *[]uint64, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newUint64Value(v.(*uint64))
	}))
}

// -- uint8 Value
type uint8Value struct{ v *uint8 }

func newUint8Value(p *uint8) *uint8Value {
	return &uint8Value{p}
}

func (f *uint8Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 8)
	if err == nil {
		*f.v = (uint8)(v)
	}
	return err
}

func (f *uint8Value) Get() interface{} { return (uint8)(*f.v) }

func (f *uint8Value) String() string { return fmt.Sprintf("%v", *f.v) }

// Uint8 parses the next command-line value as uint8.
func (p *Clause) Uint8() (target *uint8) {
	target = new(uint8)
	p.Uint8Var(target)
	return
}

func (p *Clause) Uint8Var(target *uint8) {
	p.SetValue(newUint8Value(target))
}

// Uint8List accumulates uint8 values into a slice.
func (p *Clause) Uint8List(options ...AccumulatorOption) (target *[]uint8) {
	target = new([]uint8)
	p.Uint8ListVar(target, options...)
	return
}

func (p *Clause) Uint8ListVar(target *[]uint8, options ...AccumulatorOption) {
	p.SetValue(newAccumulator(target, options, func(v interface{}) Value {
		return newUint8Value(v.(*uint8))
	}))
}