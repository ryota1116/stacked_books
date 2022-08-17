package model

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	jatranslations "github.com/go-playground/validator/v10/translations/ja"
	"github.com/ryota1116/stacked_books/domain/model/user"
)

// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

type Error struct {
	Messages []string `json:"error_message"`
}

////エラーレスポンス用の構造体
type ErrResponse struct {
	Code   int `json:"code"`
	Errors Error
}

func RespondErrJson(code int, errmap map[string]string) ErrResponse {
	var errResponse ErrResponse
	errResponse.Code = code

	for _, v := range errmap {
		errResponse.Errors.Messages = append(errResponse.Errors.Messages, v)
		fmt.Println(errResponse.Errors.Messages)
	}

	return errResponse
}

// TODO: 最終的に全構造体のバリデーションを１つのメソッドに集約させる
// Userストラクト用のバリデーター
func UserValidate(user user.User) (int, map[string]string) {
	translator := ja.New()
	uni := ut.New(translator, translator)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("ja")

	validate = validator.New()
	jatranslations.RegisterDefaultTranslations(validate, trans)

	code, errmap := translateAll(trans, user)

	return code, errmap
}

func translateAll(trans ut.Translator, user user.User) (int, map[string]string) {
	var code int
	errmap := map[string]string{}

	err := validate.Struct(user)

	if err != nil {
		code = 400

		errs := err.(validator.ValidationErrors)
		// 全てのエラーを一度に翻訳
		errmap = errs.Translate(trans)
	}
	// バリデーションエラーが無い場合はnullを返す
	return code, errmap
}

func translateIndividual(trans ut.Translator, user user.User) {
	err := validate.Struct(user)
	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			fmt.Println(e.Translate(trans))
		}
	}
}

func translateOverride(trans ut.Translator, user user.User) {
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

// TODO: map[string]stringをmap[string]interfaceに変える
func MapToStruct(m map[string]string, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}
