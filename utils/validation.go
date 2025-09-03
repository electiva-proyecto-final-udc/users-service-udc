package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Map con mensajes personalizados
var customMessages = map[string]string{
	"required": "is required",
	"numeric":  "is not a numeric value",
	"email":    "is not a valid email",
	// "len":      "no tiene la longitud correcta",
	// "min":      "tiene un valor menor al permitido",
	// "max":      "tiene un valor mayor al permitido",
	// "oneof":    "no está entre los valores permitidos",
}

func ValidateEntity(entity interface{}) map[string]string {
	err := validate.Struct(entity)
	if err == nil {
		return nil
	}

	errs := err.(validator.ValidationErrors)
	errorsMap := make(map[string]string)

	for _, e := range errs {
		// Nombre del campo en JSON (no el nombre en Go)
		fieldName := e.Field()
		// Nombre de la regla que falló
		tag := e.Tag()

		// Busca mensaje en customMessages o usa uno por defecto
		msg, ok := customMessages[tag]
		if !ok {
			msg = tag
		}

		errorsMap[fieldName] = fmt.Sprintf("field %s %s", fieldName, msg)
	}

	return errorsMap
}
