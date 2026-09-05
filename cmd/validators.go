package main

import (
	"pizza-tracker/internal/models"
	"slices"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)


func RegisterCustomValidators(){

	if v,ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.registerValidation("valid_pizza_type",createSliceValidator(models.PizzaTypes))
		v.registerValidation("valid_pizza_size",createSliceValidator(models.PizzaSizes))

	}

}

func createSliceValidator(allowedValues []string) validator.Func {
	return func (fl validator.FieldLevel) bool  {

		return slices.Contains(allowedValues,fl.Field().String())
		
	}
}