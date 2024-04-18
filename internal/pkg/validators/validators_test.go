package validators_test

import (
	"github.com/chrishadi/movies/internal/pkg/models"
	"github.com/chrishadi/movies/internal/pkg/validators"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validators", func() {

	Describe("validateDirector", func() {
		When("director's name is empty", func() {
			It("emits error", func() {
				director := models.Director{Name: ""}
				errs := validators.ValidateDirector(&director)
				Expect(errs).NotTo(BeEmpty())
			})
		})

		When("director's name is blank", func() {
			It("emits error", func() {
				director := models.Director{Name: "   "}
				errs := validators.ValidateDirector(&director)
				Expect(errs).NotTo(BeEmpty())
			})
		})

		When("director's name is not empty or blank", func() {
			It("does not emit error", func() {
				director := models.Director{Name: "George Lucas"}
				errs := validators.ValidateDirector(&director)
				Expect(errs).To(BeEmpty())
			})
		})
	})
})
