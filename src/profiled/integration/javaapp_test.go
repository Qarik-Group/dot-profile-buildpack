package integration_test

import (
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack/cutlass"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Java app Integration Test", func() {
	var app *cutlass.App
	AfterEach(func() {
		if app != nil {
			// app.Destroy()
		}
		app = nil
	})

	It("app deploys", func() {
		app = cutlass.New(filepath.Join(bpDir, "fixtures", "javaapp"))
		app.Buildpacks = []string{"profiled_buildpack", "java_buildpack"}
		app.Path = filepath.Join(filepath.Join(bpDir, "fixtures", "javaapp", "build/libs/java-main-application-1.0.0.jar"))
		app.Memory = "512M"
		app.SetEnv("JAVA_OPTS", "-Xss349k")
		PushAppAndConfirm(app)
		Expect(app.GetBody("/")).To(ContainSubstring("Found created file file-created-by-profiled"))
	})
})
