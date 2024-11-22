package login_test

import (
	"errors"
	"unit-test/login-function/login"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AuthService", func() {
	var authService *login.AuthService

	BeforeEach(func() {
		authService = login.NewAuthService()
	})

	When("the user provides valid credentials", func() {
		It("should successfully log in", func() {
			user, err := authService.Login("admin", "password123")
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
			Expect(user.Username).To(Equal("admin"))
		})
	})

	When("the user provides an invalid password", func() {
		It("should return an error", func() {
			user, err := authService.Login("admin", "wrongpassword")
			Expect(err).To(MatchError(errors.New("invalid password")))
			Expect(user).To(BeNil())
		})
	})

	When("the user provides a non-existent username", func() {
		It("should return an error", func() {
			user, err := authService.Login("nonexistent", "password123")
			Expect(err).To(MatchError(errors.New("user not found")))
			Expect(user).To(BeNil())
		})
	})
})
