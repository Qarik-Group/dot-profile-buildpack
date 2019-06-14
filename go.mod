module simple-pancake-buildpack

go 1.12

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/cloudfoundry/libbuildpack v0.0.0-20190606141245-8119571fa48b
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.4.3
	golang.org/x/sys v0.0.0-20190613124609-5ed2794edfdc // indirect
)

// replace github.com/cloudfoundry/libbuildpack => github.com/drnic/libbuildpack cutlass-skip-tls-verify
replace github.com/cloudfoundry/libbuildpack => github.com/drnic/libbuildpack v0.0.0-20190606002122-7960f8cf829c
