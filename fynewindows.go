package main

import (
	b64 "encoding/base64"
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// FormData is a struct to hold form input data.
type FormData struct {
	appName      string
	password     string
	passwordConf string
}

// addPasswordWindow creates a window for adding a new password.
// It validates the input, encodes the password in base64, and writes it to a file.
func addPasswordWindow(a fyne.App) {
	data := &FormData{}
	var passwordSlice []string

	addW := a.NewWindow("Add Password")
	addW.Resize(fyne.NewSize(500, 100))

	// Create input fields for app name, password, and password confirmation.
	appEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()
	passwordConfEntry := widget.NewPasswordEntry()

	// Create a form with input fields and submit/cancel actions.
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "App", Widget: appEntry},
			{Text: "Password", Widget: passwordEntry},
			{Text: "Password Confirmation", Widget: passwordConfEntry},
		},
		OnSubmit: func() {
			// Validate input and encode the password.
			data.appName = appEntry.Text
			data.password = passwordEntry.Text
			data.passwordConf = passwordConfEntry.Text

			if data.appName == "" || data.password == "" || data.passwordConf == "" {
				dialog.ShowError(errors.New("Invalid Entry"), addW)
				return
			}

			if data.password == data.passwordConf {
				data.password = b64.StdEncoding.EncodeToString([]byte(data.password))
				passwordSlice = append(passwordSlice, data.appName, data.password)
			} else {
				dialog.ShowError(errors.New("Passwords don't match"), addW)
				return
			}

			// Write the password to a file and show success or error messages.
			err := addPassword(filePath, passwordSlice)
			if err != nil {
				dialog.ShowError(err, addW)
				return
			}

			dialog.ShowInformation("Success", "File successfully written.", addW)
		},
		OnCancel: func() {
			addW.Close()
		},
		SubmitText:  "",
		CancelText:  "",
		Orientation: 0,
	}
	addW.SetContent(form)
	addW.Show()
}

// getPasswordWindow creates a window for retrieving a password.
// It decodes the password from base64 and displays it to the user.
func getPasswordWindow(a fyne.App) {
	data := &FormData{}

	getW := a.NewWindow("Get Password")
	getW.Resize(fyne.NewSize(500, 300))

	// Create an input field for the app name.
	appEntry := widget.NewEntry()

	// Create a form with the input field and submit/cancel actions.
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "App", Widget: appEntry},
		},
		OnSubmit: func() {
			// Retrieve and decode the password.
			data.appName = appEntry.Text
			s, err := getPassword(filePath, data.appName)
			if err != nil {
				dialog.ShowError(err, getW)
				return
			}
			decodedPsw, _ := b64.StdEncoding.DecodeString(s[1])
			dialog.ShowInformation(s[0], string(decodedPsw), getW)
		},
		OnCancel: func() {
			getW.Close()
		},
		SubmitText:  "",
		CancelText:  "",
		Orientation: 0,
	}
	getW.SetContent(form)
	getW.Show()
}

// deletePasswordWindow creates a window for deleting a password.
// It confirms the action with the user before deleting the password.
func deletePasswordWindow(a fyne.App) {
	data := &FormData{}

	deleteW := a.NewWindow("Delete Password")
	deleteW.Resize(fyne.NewSize(500, 300))

	// Create an input field for the app name.
	appEntry := widget.NewEntry()

	// Create a form with the input field and submit/cancel actions.
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "App", Widget: appEntry},
		},
		OnSubmit: func() {
			// Confirm the deletion and delete the password if confirmed.
			data.appName = appEntry.Text
			dialog.ShowConfirm(
				"Confirmation",
				"Are you sure ?",
				func(response bool) {
					if response {
						err := deletePassword(filePath, data.appName)
						if err != nil {
							dialog.ShowError(err, deleteW)
						} else {
							dialog.ShowInformation("Success", "Password deleted successfully!", deleteW)
						}
					}
				},
				deleteW,
			)
		},
		OnCancel: func() {
			deleteW.Close()
		},
		SubmitText:  "",
		CancelText:  "",
		Orientation: 0,
	}
	deleteW.SetContent(form)
	deleteW.Show()
}

// updatePasswordWindow creates a window for updating an existing password.
// It validates the input, encodes the new password in base64, and updates it in the file.
func updatePasswordWindow(a fyne.App) {
	var passwordSlice []string
	data := &FormData{}

	updateW := a.NewWindow("Update Password")
	updateW.Resize(fyne.NewSize(500, 300))

	// Create input fields for app name, new password, and password confirmation.
	appEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()
	passwordConfEntry := widget.NewPasswordEntry()

	// Create a form with input fields and submit/cancel actions.
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "App", Widget: appEntry},
			{Text: "Password", Widget: passwordEntry},
			{Text: "Password Confirmation", Widget: passwordConfEntry},
		},
		OnSubmit: func() {
			// Validate input and encode the new password.
			data.appName = appEntry.Text
			data.password = passwordEntry.Text
			data.passwordConf = passwordConfEntry.Text

			if data.appName == "" || data.password == "" || data.passwordConf == "" {
				dialog.ShowError(errors.New("Invalid Entry"), updateW)
				return
			}

			if data.password == data.passwordConf {
				data.password = b64.StdEncoding.EncodeToString([]byte(data.password))
				passwordSlice = append(passwordSlice, data.appName, data.password)
			} else {
				dialog.ShowError(errors.New("Passwords don't match"), updateW)
				return
			}

			// Confirm the update and update the password if confirmed.
			dialog.ShowConfirm("Confirmation", "Are you sure?", func(response bool) {
				if response {
					err := updatePassword(filePath, passwordSlice)
					if err != nil {
						dialog.ShowError(err, updateW)
						return
					}
					dialog.ShowInformation("Success", "Password successfully updated", updateW)
				}
			}, updateW)
		},
		OnCancel: func() {
			updateW.Close()
		},
		SubmitText:  "",
		CancelText:  "",
		Orientation: 0,
	}
	updateW.SetContent(form)
	updateW.Show()
}
