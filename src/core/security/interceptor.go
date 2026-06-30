package security

import (
	"reflect"
)

// PrepareForWrite takes a struct or pointer to struct, clones it,
// and encrypts all string fields marked with `encrypt:"true"`.
func PrepareForWrite[T any](entity T) (T, error) {
	val := reflect.ValueOf(entity)

	isPtr := val.Kind() == reflect.Ptr
	if isPtr {
		val = val.Elem()
	}

	// If not a struct, return as is
	if val.Kind() != reflect.Struct {
		return entity, nil
	}

	// Create a new instance of the struct type
	cloneVal := reflect.New(val.Type()).Elem()

	// Copy fields and encrypt if necessary
	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		cloneField := cloneVal.Field(i)

		// Copy the value first
		cloneField.Set(fieldVal)

		// Check the tag
		fieldType := val.Type().Field(i)
		encryptTag := fieldType.Tag.Get("encrypt")

		if encryptTag == "true" && fieldVal.Kind() == reflect.String {
			originalStr := fieldVal.String()
			encryptedStr, err := EncryptString(originalStr)
			if err != nil {
				return entity, err
			}
			if cloneField.CanSet() {
				cloneField.SetString(encryptedStr)
			}
		}
	}

	if isPtr {
		// Return a pointer to the clone
		return cloneVal.Addr().Interface().(T), nil
	}

	// Return the clone directly
	return cloneVal.Interface().(T), nil
}

// MaterializeFromRead takes a struct or pointer to struct, clones it,
// and decrypts all string fields marked with `encrypt:"true"`.
func MaterializeFromRead[T any](entity T) (T, error) {
	val := reflect.ValueOf(entity)

	isPtr := val.Kind() == reflect.Ptr
	if isPtr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return entity, nil
	}

	cloneVal := reflect.New(val.Type()).Elem()

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		cloneField := cloneVal.Field(i)

		cloneField.Set(fieldVal)

		fieldType := val.Type().Field(i)
		encryptTag := fieldType.Tag.Get("encrypt")

		if encryptTag == "true" && fieldVal.Kind() == reflect.String {
			encryptedStr := fieldVal.String()
			decryptedStr, err := DecryptString(encryptedStr)
			if err != nil {
				return entity, err
			}
			if cloneField.CanSet() {
				cloneField.SetString(decryptedStr)
			}
		}
	}

	if isPtr {
		return cloneVal.Addr().Interface().(T), nil
	}

	return cloneVal.Interface().(T), nil
}
