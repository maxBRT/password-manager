package main

import (
	"fyne.io/fyne/v2"           // Fyne GUI framework
	"fyne.io/fyne/v2/app"       // Fyne app creation
	"fyne.io/fyne/v2/container" // Fyne container utilities
	"fyne.io/fyne/v2/layout"    // Fyne layout utilities
	"fyne.io/fyne/v2/widget"    // Fyne widget utilities
)

const filePath = "pswd.csv" // File path for storing passwords

// main initializes the password manager application and its UI.
func main() {
	a := app.New()                       // Create a new Fyne application
	w := a.NewWindow("Password Manager") // Create a new window with the title "Password Manager"
	w.Resize(fyne.NewSize(320, 50))      // Set the initial size of the window

	// Create a button for adding a password, which opens a new window for the operation
	addButton := widget.NewButton("Add Password", func() {
		addPasswordWindow(a)
	})

	// Create a button for retrieving a password, which opens a new window for the operation
	getButton := widget.NewButton("Get Password", func() {
		getPasswordWindow(a)
	})

	// Create a button for updating a password, which opens a new window for the operation
	updateButton := widget.NewButton("Update Password", func() {
		updatePasswordWindow(a)
	})

	// Create a button for deleting a password, which opens a new window for the operation
	deleteButton := widget.NewButton("Delete Password", func() {
		deletePasswordWindow(a)
	})

	// Arrange the buttons in a horizontal layout
	content := container.New(
		layout.NewHBoxLayout(), // Horizontal box layout
		addButton,
		getButton,
		updateButton,
		deleteButton,
	)

	// Set the content of the main window and run the application
	w.SetContent(
		container.New(
			layout.NewVBoxLayout(), // Vertical box layout
			content,
		),
	)
	w.ShowAndRun() // Display the window and start the application event loop
}
