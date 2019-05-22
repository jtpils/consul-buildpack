package integration_test

import (
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack/cutlass"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Simple Integration Test", func() {
	var app *cutlass.App
	AfterEach(func() {
		if app != nil {
			app.Destroy()
		}
		app = nil
	})

	It("app deploys", func() {
		app = cutlass.New(filepath.Join(bpDir, "fixtures", "static"))
		app.Manifest = filepath.Join(bpDir, "fixtures", "static", "manifest.cfdev.yml")
		V3PushAppAndConfirm(app)
		Expect(app.GetBody("/")).To(ContainSubstring("Hi, I'm an app with a consul agent sidecar!"))
	})
})
