package model

import (
	"fmt"
	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	jatranslations "github.com/go-playground/validator/v10/translations/ja"
)

// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

// TODO: 最終的に全構造体のバリデーションを１つのメソッドに集約させる
// Userストラクト用のバリデーター
func UserValidate(user User) {
	translator := ja.New()
	uni := ut.New(translator, translator)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("ja")

	validate = validator.New()
	jatranslations.RegisterDefaultTranslations(validate, trans)

	translateAll(trans, user)
	translateIndividual(trans, user)
	translateOverride(trans, user)
}



func translateAll(trans ut.Translator, user User) {
	err := validate.Struct(user)
	fmt.Println(err)
	if err != nil {
		// 全てのエラーを一度に翻訳
		errs := err.(validator.ValidationErrors)

		fmt.Println(errs.Translate(trans))
	}
}

func translateIndividual(trans ut.Translator, user User) {
	err := validate.Struct(user)
	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			fmt.Println(e.Translate(trans))
		}
	}
}

func translateOverride(trans ut.Translator, user User) {
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	err := validate.Struct(user)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			fmt.Println(e.Translate(trans))
		}
	}
}
