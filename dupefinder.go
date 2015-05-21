package dupefinder

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Generate(catalog string, folders ...string) error {
	err := validateFolders(folders...)
	if err != nil {
		return err
	}

	errs := make(chan error)

	go walkAllFolders(errs, folders...)

	err = <-errs
	if err != nil {
		return err
	}

	return nil
}

func Detect(catalog string, echo, rm bool, folders ...string) error {
	err := validateFolders(folders...)
	if err != nil {
		return err
	}

	return nil
}

func validateFolders(folders ...string) error {
	for _, f := range folders {
		isfolder, err := isFolder(f)
		if err != nil {
			return err
		}
		if !isfolder {
			return fmt.Errorf("%s is not a folder", f)
		}
	}

	return nil
}

func isFolder(filename string) (bool, error) {
	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}

func walkAllFolders(errs chan error, folders ...string) {
	defer close(errs)

	for _, f := range folders {
		err := walkFolder(f)
		if err != nil {
			errs <- err
			return
		}
	}
}

func walkFolder(filename string) error {
	fi, err := ioutil.ReadDir(filename)
	if err != nil {
		return err
	}

	fmt.Println(filename)
	for _, child := range fi {
		fmt.Println(child)
	}

	return nil
}
