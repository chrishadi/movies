package validators

import (
	"errors"
	"strings"

	"github.com/chrishadi/movies/internal/pkg/models"
)

func ValidateDirector(director *models.Director) []error {
	var errs []error

	name := strings.TrimSpace(director.Name)
	if name == "" {
		errs = append(errs, errors.New("director's name cannot be blank"))
	}

	return errs
}
