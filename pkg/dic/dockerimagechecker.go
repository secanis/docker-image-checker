// Package dic fetch data from a Docker API and prepare it
package dic

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

// Image object for a Docker API call
type Image struct {
	Name        string
	FullSize    int
	ID          int
	Repository  int
	Creator     int
	LastUpdater int    `json:"last_updater"`
	LastUpdated string `json:"last_updated"`
	ImageID     string
	V2          bool
}

type Response struct {
	Image Image
}

// GetTagObject function returns a Docker image object
// example URL https://hub.docker.com/v2/repositories/library/alpine/tags/latest
func GetTagObject(url string, owner string, imageName string, tagName string) (*Image, error) {
	resp, err := http.Get("https://" + url + "/v2/repositories/" + owner + "/" + imageName + "/tags/" + tagName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	// _, err = io.Copy(os.Stdout, resp.Body)
	r := new(Image)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, errors.New(resp.Status)
	}

	return r, nil
}

// OutdatedImage function returns a bool if the "image to check" is older than base
func OutdatedImage(baseImage time.Time, imageToCheck time.Time) bool {
	return baseImage.After(imageToCheck)
}
