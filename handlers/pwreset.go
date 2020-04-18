package handlers

import (
	"app/conf"
	"app/models"
	"app/utils"
	"net/http"

	"github.com/go-playground/validator"
)

type pwResetForm struct {
	Email  string `validate:"required,email"`
	Errors map[string]string
}

func (f *pwResetForm) Validate() bool {
	f.Errors = make(map[string]string)
	msgs := map[string]string{
		"required": "This field is required.",
		"email":    "Please enter a valid email address.",
	}
	if err := conf.Validate.Struct(f); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			f.Errors[err.Field()] = msgs[err.Tag()]
		}
	}
	return len(f.Errors) == 0
}

func PWResetGet(appContext *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	context := struct {
		Form *pwResetForm
	}{
		Form: &pwResetForm{},
	}
	return appContext.TemplateResponse(w, "pwresetHTML", context)
}

func PWResetPost(appContext *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	form := &pwResetForm{
		Email: r.PostFormValue("email"),
	}
	if form.Validate() == false {
		context := struct {
			Form *pwResetForm
		}{
			Form: form,
		}
		return appContext.TemplateResponse(w, "pwresetHTML", context)
	}
	// get the user - for security reasonns we dont want to alert the user
	// if the email doesnt exist
	user := &models.User{}
	if err := appContext.Db.Where("email = ?", form.Email).First(&user).Error; err == nil {
		go func() { user.SendPasswordReset(r) }()
	}
	// finally redirect them even if anything failed
	http.Redirect(w, r, "/auth/password/reset/done", http.StatusFound)
	return http.StatusFound, nil
}

func PWResetDone(appContext *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	return appContext.TemplateResponse(w, "pwresetdoneHTML", nil)
}

type pwResetConfirmForm struct {
	NewPassword    string `validate:"required"`
	RepeatPassword string `validate:"required,eqfield=NewPassword"`
	Errors         map[string]string
}

func (f *pwResetConfirmForm) Validate() bool {
	f.Errors = make(map[string]string)
	msgs := map[string]string{
		"required": "This field is required.",
		"eqfield":  "The passwords do not match.",
	}
	if err := conf.Validate.Struct(f); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			f.Errors[err.Field()] = msgs[err.Tag()]
		}
	}
	return len(f.Errors) == 0
}

func getUserFromToken(appContext *conf.AppContext, r *http.Request) (*models.User, error) {
	token := r.FormValue("token")
	// check the token is valid
	pwhash, err := utils.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	// check the users password is still as it was when they requested a reset
	user := &models.User{}
	if err := appContext.Db.Where("password = ?", pwhash).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func PWResetConfirmGet(appContext *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	if _, err := getUserFromToken(appContext, r); err != nil {
		return http.StatusNotFound, err
	}
	context := struct {
		Form *pwResetConfirmForm
	}{
		Form: &pwResetConfirmForm{},
	}
	return appContext.TemplateResponse(w, "pwresetconfirmHTML", context)
}

func PWResetConfirmPost(appContext *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	user, err := getUserFromToken(appContext, r)
	if err != nil {
		return http.StatusNotFound, err
	}
	// validate the form
	form := &pwResetConfirmForm{
		NewPassword:    r.PostFormValue("new_password"),
		RepeatPassword: r.PostFormValue("repeat_password"),
	}
	if form.Validate() == false {
		context := struct {
			Form *pwResetConfirmForm
		}{
			Form: form,
		}
		return appContext.TemplateResponse(w, "pwresetconfirmHTML", context)
	}
	user.SetPassword(form.NewPassword)
	if err := models.Db.Save(&user).Error; err != nil {
		return http.StatusInternalServerError, err
	}
	http.Redirect(w, r, "/auth/password/reset/complete", http.StatusFound)
	return http.StatusFound, nil
}

func PWResetComplete(appContext *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	return appContext.TemplateResponse(w, "pwresetcompleteHTML", nil)
}
