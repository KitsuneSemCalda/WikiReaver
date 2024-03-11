package urltest

import (
	"errors"
	"net/http"
)

func CheckUrl(url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	defer resp.Body.Close()

	return errors.New("occurs an unknown error")
}
