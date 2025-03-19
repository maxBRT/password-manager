package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

// check is a helper function to handle errors by printing and panicking.
func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

// addPassword adds a new password to the CSV file.
// It checks if the app already exists and appends the new password if not.
func addPassword(filePath string, s []string) error {
	fileExits := false

	// Check if the file exists.
	if _, err := os.Stat(filePath); err == nil {
		fileExits = true
	}

	if fileExits {
		// Open the file and read its contents.
		f, err := os.Open(filePath)
		check(err)

		reader := csv.NewReader(f)

		records, err := reader.ReadAll()
		check(err)

		// Check if the app already exists in the file.
		for _, row := range records {
			if slices.Contains(row, s[0]) {
				f.Close()
				return errors.New("Password already exist for app")
			}
		}
		f.Close()
	}

	// Open the file in append mode or create it if it doesn't exist.
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	// Write the header if the file is newly created.
	if !fileExits {
		header := []string{"App", "Password"}
		err = writer.Write(header)
		check(err)
	}

	// Write the new password row.
	row := s
	err = writer.Write(row)
	check(err)

	return nil
}

// getPassword retrieves a password for a given app from the CSV file.
// It returns the app and password if found, or an error if not.
func getPassword(filePath string, app string) ([]string, error) {
	var s []string
	f, err := os.Open(filePath)
	check(err)
	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	check(err)

	// Search for the app in the records.
	for _, row := range records {
		if slices.Contains(row, strings.ToLower(app)) {
			return row, nil
		}
	}
	return s, errors.New("Password not found.")
}

// deletePassword deletes a password for a given app from the CSV file.
// It rewrites the file without the deleted app's record.
func deletePassword(filePath string, app string) error {
	pswFound := false
	f, err := os.Open(filePath)
	defer f.Close()
	check(err)

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	check(err)

	var filteredRecord [][]string
	filteredRecord = append(filteredRecord, records[0]) // Keep the header.

	// Filter out the record for the given app.
	for _, row := range records[1:] {
		if slices.Contains(row, strings.ToLower(app)) {
			pswFound = true
			continue
		} else {
			filteredRecord = append(filteredRecord, row)
		}
	}

	if pswFound {
		// Rewrite the file with the filtered records.
		err := os.Remove(filePath)
		check(err)

		newF, err := os.Create(filePath)
		check(err)

		writer := csv.NewWriter(newF)
		defer writer.Flush()

		err = writer.WriteAll(filteredRecord)
		check(err)
		return nil
	} else {
		return errors.New("Password not found.")
	}
}

// updatePassword updates the password for a given app in the CSV file.
// It rewrites the file with the updated record.
func updatePassword(filePath string, updatedRow []string) error {
	pswFound := false
	f, err := os.Open(filePath)
	defer f.Close()
	check(err)

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	check(err)

	var filteredRecord [][]string
	filteredRecord = append(filteredRecord, records[0]) // Keep the header.

	// Replace the record for the given app with the updated record.
	for _, row := range records[1:] {
		if slices.Contains(row, strings.ToLower(updatedRow[0])) {
			pswFound = true
			continue
		} else {
			filteredRecord = append(filteredRecord, row)
		}
	}

	if pswFound {
		filteredRecord = append(filteredRecord, updatedRow)

		// Rewrite the file with the updated records.
		err := os.Remove(filePath)
		check(err)

		newF, err := os.Create(filePath)
		check(err)

		writer := csv.NewWriter(newF)
		defer writer.Flush()

		err = writer.WriteAll(filteredRecord)
		check(err)

		return nil
	} else {
		return errors.New("Password not found.")
	}
}
