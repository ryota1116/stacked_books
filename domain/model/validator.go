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
func UserValidate(user User) map[string]string {
	translator := ja.New()
	uni := ut.New(translator, translator)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("ja")

	validate = validator.New()
	jatranslations.RegisterDefaultTranslations(validate, trans)

	fmt.Println("--------------")
	var errmap map[string]string
	errmap = translateAll(trans, user)
	fmt.Println("--------------")
	//translateIndividual(trans, user)
	fmt.Println("--------------")
	//translateOverride(trans, user)

	return errmap
}


func translateAll(trans ut.Translator, user User) map[string]string {
	err := validate.Struct(user)
	fmt.Println(err)

	errmap := map[string]string{}
	if err != nil {
		// 全てのエラーを一度に翻訳
		errs := err.(validator.ValidationErrors)

		errmap = errs.Translate(trans)
	}
	// バリデーションエラーが無い場合はnullを返す
	return errmap
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

////エラーレスポンス用の構造体
//type ErrResponse struct {
//	Code string `json:"error"`
//	Message string `json:"error_description"`
//	Status int `json:"status"`
//}
//
//func respondJson()  {
//	errRes := ErrResponse{
//		Code:    "",
//		Message: "",
//		Status:  0,
//	}
//	json.NewEncoder().Encode(errRes)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusBadRequest)
//}
//
//// Mapを構造体に変換する
//func MapToStruct(m map[string]interface{}, val interface{}) error {
//	// mapをJSONに変換
//	tmp, err := json.Marshal(m)
//	if err != nil {
//		return err
//	}
//	// JSONをStructに変換
//	err = json.Unmarshal(tmp, val)
//	if err != nil {
//		return err
//	}
//	return nil
//}