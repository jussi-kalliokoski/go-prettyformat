package prettyformat

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

var (
	// ErrInvalidType is returned when the type of the passed value cannot be determined
	ErrInvalidType = errors.New("invalid type")
	// ErrArbitraryPointerType is returned when an arbitrary pointer is passed
	ErrArbitraryPointerType = errors.New("arbitrary pointer types cannot be serialized reliably")
	// ErrFunctionType is returned when a function is passed
	ErrFunctionType = errors.New("functions cannot be serialized")
	// ErrInterfaceType is returned when an interface instance is passed
	ErrInterfaceType = errors.New("interfaces cannot be serialized")
	// ErrChanType is returned when a channel is passed
	ErrChanType = errors.New("channels cannot be serialized")
)

const (
	paddingIncrement = 2
)

// Format a value into a pretty-printed string
func Format(value interface{}) (string, error) {
	return formatPadded(value, paddingIncrement)
}

func formatPadded(value interface{}, padding int) (string, error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Invalid:
		return formatInvalid(value, padding)
	case reflect.Bool:
		return formatBool(value, padding)
	case reflect.Int:
		return formatInt(value, padding)
	case reflect.Int8:
		return formatInt8(value, padding)
	case reflect.Int16:
		return formatInt16(value, padding)
	case reflect.Int32:
		return formatInt32(value, padding)
	case reflect.Int64:
		return formatInt64(value, padding)
	case reflect.Uint:
		return formatUint(value, padding)
	case reflect.Uint8:
		return formatUint8(value, padding)
	case reflect.Uint16:
		return formatUint16(value, padding)
	case reflect.Uint32:
		return formatUint32(value, padding)
	case reflect.Uint64:
		return formatUint64(value, padding)
	case reflect.Uintptr:
		return formatUintptr(value, padding)
	case reflect.Float32:
		return formatFloat32(value, padding)
	case reflect.Float64:
		return formatFloat64(value, padding)
	case reflect.Complex64:
		return formatComplex64(value, padding)
	case reflect.Complex128:
		return formatComplex128(value, padding)
	case reflect.Array:
		return formatArray(value, padding)
	case reflect.Chan:
		return formatChan(value, padding)
	case reflect.Func:
		return formatFunc(value, padding)
	case reflect.Interface:
		return formatInterface(value, padding)
	case reflect.Map:
		return formatMap(value, padding)
	case reflect.Ptr:
		return formatPtr(value, padding)
	case reflect.Slice:
		return formatSlice(value, padding)
	case reflect.String:
		return formatString(value, padding)
	case reflect.Struct:
		return formatStruct(value, padding)
	case reflect.UnsafePointer:
		return formatUnsafePointer(value, padding)
	default:
		return formatInvalid(value, padding)
	}
}

func formatInvalid(value interface{}, padding int) (string, error) {
	return "", ErrInvalidType
}

func formatBool(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%v", value.(bool)), nil
}

func formatInt(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%v", value.(int)), nil
}

func formatInt8(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%v", value.(int8)), nil
}

func formatInt16(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%v", value.(int16)), nil
}

func formatInt32(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%v", value.(int32)), nil
}

func formatInt64(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%v", value.(int64)), nil
}

func formatUint(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%v", value.(uint)), nil
}

func formatUint8(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%v", value.(uint8)), nil
}

func formatUint16(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%v", value.(uint16)), nil
}

func formatUint32(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%v", value.(uint32)), nil
}

func formatUint64(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%v", value.(uint64)), nil
}

func formatUintptr(value interface{}, padding int) (string, error) {
	return "", ErrArbitraryPointerType
}

func formatFloat32(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%v", value.(float32)), nil
}

func formatFloat64(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%v", value.(float64)), nil
}

func formatComplex64(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%v", value.(complex64)), nil
}

func formatComplex128(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%v", value.(complex128)), nil
}

func formatArray(value interface{}, padding int) (string, error) {
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)
	elemType := t.Elem()

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("[%d]%s{", v.Len(), typeName(elemType)))

	buffer.WriteString("\n")
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		elemInterface := elem.Interface()
		formattedValue, err := formatPadded(elemInterface, padding+paddingIncrement)
		if err != nil {
			return "", err
		}

		buffer.WriteString(strings.Repeat(" ", padding))

		if elemType.Kind() == reflect.Interface && !explicitType(reflect.TypeOf(elemInterface)) {
			buffer.WriteString(fmt.Sprintf("(%s)%s,\n", typeName(reflect.TypeOf(elemInterface)), formattedValue))
		} else {
			buffer.WriteString(fmt.Sprintf("%s,\n", formattedValue))
		}
	}

	buffer.WriteString(strings.Repeat(" ", padding-paddingIncrement))

	buffer.WriteString("}")
	return buffer.String(), nil
}

func formatChan(value interface{}, padding int) (string, error) {
	return "", ErrChanType
}

func formatFunc(value interface{}, padding int) (string, error) {
	return "", ErrFunctionType
}

func formatInterface(value interface{}, padding int) (string, error) {
	return "", ErrInterfaceType
}

type mapEntry struct {
	key      string
	keyType  string
	elem     string
	elemType string
}

type mapEntries []mapEntry

func (e mapEntries) Len() int {
	return len(e)
}

func (e mapEntries) Less(i, j int) bool {
	return e[i].key < e[j].key
}

func (e mapEntries) Swap(i, j int) {
	el := e[i]
	e[i] = e[j]
	e[j] = el
}

func formatMap(value interface{}, padding int) (string, error) {
	t := reflect.TypeOf(value)
	keyType := t.Key()
	elemType := t.Elem()

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("%s{", typeName(t)))

	v := reflect.ValueOf(value)
	entries := make(mapEntries, v.Len())

	if v.Len() == 0 {
		goto finish
	}

	buffer.WriteString("\n")

	for i, key := range v.MapKeys() {
		keyInterface := key.Interface()
		formattedKey, err := formatPadded(keyInterface, padding+paddingIncrement)
		if err != nil {
			return "", err
		}

		entries[i].key = formattedKey
		if keyType.Kind() == reflect.Interface && !explicitType(reflect.TypeOf(keyInterface)) {
			entries[i].keyType = "(" + typeName(reflect.TypeOf(keyInterface)) + ")"
		}

		elem := v.MapIndex(key)
		elemInterface := elem.Interface()
		formattedElem, err := formatPadded(elemInterface, padding+paddingIncrement)
		if err != nil {
			return "", err
		}

		entries[i].elem = formattedElem
		if elemType.Kind() == reflect.Interface && !explicitType(reflect.TypeOf(elemInterface)) {
			entries[i].elemType = "(" + typeName(reflect.TypeOf(elemInterface)) + ")"
		}
	}

	sort.Sort(entries)

	for _, entry := range entries {
		buffer.WriteString(strings.Repeat(" ", padding) + entry.keyType + entry.key + ": " + entry.elemType + entry.elem + ",\n")
	}

	buffer.WriteString(strings.Repeat(" ", padding-paddingIncrement))

finish:
	buffer.WriteString("}")
	return buffer.String(), nil
}

func formatPtr(value interface{}, padding int) (string, error) {
	v := reflect.ValueOf(value)
	if v.IsNil() {
		return "nil", nil
	}

	formattedElem, err := formatPadded(reflect.Indirect(v).Interface(), padding)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("&%s", formattedElem), nil
}

func formatSlice(value interface{}, padding int) (string, error) {
	t := reflect.TypeOf(value)
	elemType := t.Elem()

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("[]%s{", typeName(elemType)))
	v := reflect.ValueOf(value)
	if v.IsNil() || v.Len() == 0 {
		goto finish
	}

	buffer.WriteString("\n")
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i).Interface()
		formattedValue, err := formatPadded(elem, padding+paddingIncrement)
		if err != nil {
			return "", err
		}

		buffer.WriteString(strings.Repeat(" ", padding))

		if elemType.Kind() == reflect.Interface && !explicitType(reflect.TypeOf(elem)) {
			elemType := reflect.TypeOf(elem)
			buffer.WriteString(fmt.Sprintf("(%s)%s,\n", typeName(elemType), formattedValue))
		} else {
			buffer.WriteString(fmt.Sprintf("%s,\n", formattedValue))
		}
	}

	buffer.WriteString(strings.Repeat(" ", padding-paddingIncrement))

finish:
	buffer.WriteString("}")
	return buffer.String(), nil
}

func formatString(value interface{}, padding int) (string, error) {
	return fmt.Sprintf("%q", value.(string)), nil
}

func formatStruct(value interface{}, padding int) (string, error) {
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("%s{", typeName(t)))

	hasExportedFields := false
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		hidden := field.PkgPath != ""
		if hidden {
			continue
		}

		if !hasExportedFields {
			hasExportedFields = true
			buffer.WriteString("\n")
		}

		buffer.WriteString(strings.Repeat(" ", padding))

		fieldValue := v.Field(i)
		fieldValueInterface := fieldValue.Interface()
		formattedValue, err := formatPadded(fieldValueInterface, padding+paddingIncrement)
		if err != nil {
			return "", err
		}

		if field.Type.Kind() == reflect.Interface && !explicitType(reflect.TypeOf(fieldValueInterface)) {
			buffer.WriteString(field.Name + ": (" + typeName(reflect.TypeOf(fieldValueInterface)) + ")" + formattedValue + ",\n")
		} else {
			buffer.WriteString(field.Name + ": " + formattedValue + ",\n")
		}
	}

	if hasExportedFields {
		buffer.WriteString(strings.Repeat(" ", padding-paddingIncrement))
	}

	buffer.WriteString("}")
	return buffer.String(), nil
}

func formatUnsafePointer(value interface{}, padding int) (string, error) {
	return "", ErrArbitraryPointerType
}

func explicitType(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Array:
		return true
	case reflect.Map:
		return true
	case reflect.Slice:
		return true
	case reflect.Struct:
		return true
	case reflect.Interface:
		return true
	default:
		return false
	}
}

func typeName(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Array:
		return fmt.Sprintf("[%d]%s", t.Len(), typeName(t.Elem()))
	case reflect.Map:
		return fmt.Sprintf("map[%s]%s", typeName(t.Key()), typeName(t.Elem()))
	case reflect.Ptr:
		return fmt.Sprintf("*%s", typeName(t.Elem()))
	case reflect.Slice:
		return fmt.Sprintf("[]%s", typeName(t.Elem()))
	case reflect.Struct:
		if t.Name() == "" {
			return "(anonymous struct)"
		}
		return t.Name()
	case reflect.Interface:
		return "interface{}"
	default:
		return fmt.Sprintf("%s", t.Kind())
	}
}
