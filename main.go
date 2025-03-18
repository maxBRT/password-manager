package main

import (
	b64 "encoding/base64"
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const filePath = "pswd.csv"

type FormData struct {
	appName      string
	password     string
	passwordConf string
}

func main() {
	a := app.New()
	w := a.NewWindow("Password Manager")
	w.Resize(fyne.NewSize(320, 320))

	addButton := widget.NewButton("Add Password", func() {
		addPasswordWindow(a)
	})
	getButton := widget.NewButton("Get Password", func() {
		getPasswordWindow(a)
	})

	updateButton := widget.NewButton("Update Password", func() {
		updatePasswordWindow(a)
	})
	deleteButton := widget.NewButton("Delete Password", func() {
		deletePasswordWindow(a)
	})

	content := container.New(
		layout.NewHBoxLayout(),
		addButton,
		getButton,
		updateButton,
		deleteButton,
	)

	w.SetContent(
		container.New(
			layout.NewVBoxLayout(),
			widget.NewLabel("Password Manager"),
			content,
		),
	)
	w.ShowAndRun()
}

func addPasswordWindow(a fyne.App) {
	data := &FormData{}
	var passwordSlice []string

	addW := a.NewWindow("Add Password")
	addW.Resize(fyne.NewSize(500, 100))

	appEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()
	passwordConfEntry := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "App", Widget: appEntry},
			{Text: "Password", Widget: passwordEntry},
			{Text: "Password Confirmation", Widget: passwordConfEntry},
		},
		OnSubmit: func() {
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

func getPasswordWindow(a fyne.App) {
	data := &FormData{}

	getW := a.NewWindow("Get Password")
	getW.Resize(fyne.NewSize(500, 300))

	appEntry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "App", Widget: appEntry},
		},
		OnSubmit: func() {
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

func deletePasswordWindow(a fyne.App) {
	data := &FormData{}

	deleteW := a.NewWindow("Delete Password")
	deleteW.Resize(fyne.NewSize(500, 300))

	appEntry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "App", Widget: appEntry},
		},
		OnSubmit: func() {
			data.appName = appEntry.Text
			dialog.ShowConfirm(
				"Confirmation",
				"Are you sure ?",
				func(response bool) {
					// Only delete the password if the user confirms
					if response {
						// Call the deletePassword function and capture the error
						err := deletePassword(filePath, data.appName)

						// If there's an error, show it, otherwise show success
						if err != nil {
							dialog.ShowError(err, deleteW)
						} else {
							// Optionally, show a success message
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

func updatePasswordWindow(a fyne.App) {
	var passwordSlice []string
	data := &FormData{}

	updateW := a.NewWindow("Update Password")
	updateW.Resize(fyne.NewSize(500, 300))

	appEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()
	passwordConfEntry := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "App", Widget: appEntry},
			{Text: "Password", Widget: passwordEntry},
			{Text: "Password Confirmation", Widget: passwordConfEntry},
		},
		OnSubmit: func() {
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
