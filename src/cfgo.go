package cfgo

import (
	"fmt"
	"reflect"
	"strconv"
)

func LoadType(i any, cfg *EnvConfiguration) {
	iReflection := reflect.ValueOf(i)

	if iReflection.Kind() != reflect.Ptr {
		panic("cfgo: provided type is not a pointer")
	}

	if iReflection.Elem().Kind() != reflect.Struct {
		panic("cfgo: provided type is not a pointer to a struct")
	}

	valueReflection := iReflection.Elem()
	typeReflection := valueReflection.Type()

	valuesFromProviders := cfg.getAllValues()

	for index := range valueReflection.NumField() {
		field := typeReflection.Field(index)

		loadAndSetValue(valueReflection, field, index, valuesFromProviders, cfg)
	}
}

func loadAndSetValue(
	reflectValue reflect.Value,
	reflectField reflect.StructField,
	fieldIdx int,
	values map[string]string,
	cfg *EnvConfiguration,
) {
	tagKeyValue, ok := reflectField.Tag.Lookup("env")

	if !ok {
		return
	}

	value, ok := values[tagKeyValue]

	if !ok {
		panic(fmt.Sprintf("cfgo: field '%v' not found", reflectField.Name))
	}

	switch reflectField.Type.Kind() {
	case reflect.String:
		{
			reflectValue.Field(fieldIdx).SetString(value)
		}
	case reflect.Bool:
		{
			reflectValue.Field(fieldIdx).SetBool(sliceHas(cfg.BoolValidTrueValues, value))
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		{
			// TODO: get bit size based from the type directly
			// TODO: consider handling prefixes like 0b, or 0x to change the base
			parsedInt, err := strconv.ParseInt(value, 10, 64)

			if err != nil {
				panic(fmt.Sprintf("cfgo: field '%v' is not an integer", reflectField.Name))
			}

			reflectValue.Field(fieldIdx).SetInt(parsedInt)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		{
			// TODO: get bit size based from the type directly
			// TODO: consider handling prefixes like 0b, or 0x to change the base
			parsedInt, err := strconv.ParseUint(value, 10, 64)

			if err != nil {
				panic(fmt.Sprintf("cfgo: field '%v' is not an integer", reflectField.Name))
			}

			reflectValue.Field(fieldIdx).SetUint(parsedInt)
		}
	case reflect.Float32, reflect.Float64:
		{
			// TODO: get bit size based from the type directly
			parsedFloat, err := strconv.ParseFloat(value, 64)

			if err != nil {
				panic(fmt.Sprintf("cfgo: field '%v' is not a float", reflectField.Name))
			}

			reflectValue.Field(fieldIdx).SetFloat(parsedFloat)
		}
	default:
		{
			panic(fmt.Sprintf("cfgo: field '%s' is not a supported type '%s'", reflectField.Name, reflectField.Type.Kind()))
		}
	}
}
