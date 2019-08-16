package util

import "os"

// WriteToFile takes a string and writes it to a file
func WriteToFile(filename, content string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	err = os.Truncate(filename, 0)

	if err != nil {
		return err
	}

	_, err = f.WriteString(content)

	if err != nil {
		f.Close()
		return err
	}

	err = f.Close()

	if err != nil {
		return err
	}
	return nil
}
