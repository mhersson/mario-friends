package app_test

import (
	"mario-friends/internal/app"
	"net/http"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = ginkgo.Describe("SimpleHandler", func() {
	var server *ghttp.Server

	ginkgo.BeforeEach(func() {
		server = ghttp.NewServer()

		server.AppendHandlers(http.HandlerFunc(app.SimpleHandler))
	})

	ginkgo.AfterEach(func() {
		server.Close()
	})

	ginkgo.Context("when making a GET request", func() {
		ginkgo.It("should return 200 OK", func() {
			resp, err := http.Get(server.URL())
			gomega.Expect(err).ToNot(gomega.HaveOccurred())
			defer resp.Body.Close()

			gomega.Expect(resp.StatusCode).To(gomega.Equal(http.StatusOK))
		})

		ginkgo.It("should return 404 Not Found", func() {
			resp, err := http.Get(server.URL() + "/not-found")
			gomega.Expect(err).ToNot(gomega.HaveOccurred())
			defer resp.Body.Close()

			gomega.Expect(resp.StatusCode).To(gomega.Equal(http.StatusNotFound))
		})
	})
})
