package detect_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	libbuildpackV3 "github.com/buildpack/libbuildpack"
	"github.com/cloudfoundry/nodejs-cnb-buildpack/build"
	"github.com/cloudfoundry/nodejs-cnb-buildpack/detect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UpdateBuildPlan", func() {
	var (
		err        error
		dir        string
		detectData libbuildpackV3.Detect
	)

	BeforeEach(func() {
		dir, err = ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())

		detectData = libbuildpackV3.Detect{
			Application: libbuildpackV3.Application{Root: dir},
			BuildPlan:   make(map[string]libbuildpackV3.BuildPlanDependency),
		}
	})

	AfterEach(func() {
		err = os.RemoveAll(dir)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("there is a package.json with a node version in engines", func() {
		const version string = "1.2.3"

		BeforeEach(func() {
			packageJSONString := fmt.Sprintf(`{
				"name": "bson-test",
				"version": "1.0.0",
				"description": "",
				"main": "index.js",
				"scripts": {
				"start": "node index.js"
			},
				"author": "",
				"license": "ISC",
				"dependencies": {
				"bson-ext": "^0.1.13"
			},
				"engines": {
				"node" : "%s"
			}
			}`, version)
			err = ioutil.WriteFile(
				filepath.Join(dir, "package.json"),
				[]byte(packageJSONString),
				0666,
			)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should create a build plan with the required node version", func() {
			err = detect.UpdateBuildPlan(&detectData)
			Expect(err).NotTo(HaveOccurred())
			Expect(detectData.BuildPlan[build.NodeDependency].Version).To(Equal(version))
		})
	})

	Context("there is no package.json", func() {
		It("returns an error", func() {
			err = detect.UpdateBuildPlan(&detectData)
			Expect(err).To(HaveOccurred())
		})
	})
})
