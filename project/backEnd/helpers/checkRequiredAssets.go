package helpers

import (
	"fmt"
	"os"
)

func CheckRequiredAssets() []error {
	assets := []string{
		"project/frontEnd/template/home.html",
		"project/frontEnd/template/artist.html",
		"project/frontEnd/template/aboutus.html",
		"project/frontEnd/template/400.html",
		"project/frontEnd/template/404.html",
		"project/frontEnd/template/405.html",
		"project/frontEnd/template/500.html",
		"project/frontEnd/static/style.css",
	}

	var errors []error
	for _, asset := range assets {
		if _, err := os.Stat(asset); err != nil {
			errors = append(errors, fmt.Errorf("missing or unreadable asset: %s", asset))
		}
	}
	return errors
}
