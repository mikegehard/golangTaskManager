package main_test

import (
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"os/exec"
)

var Start = gexec.Start
var Say = gbytes.Say
var Build = gexec.Build
var Exit = gexec.Exit

var _ = Describe("The greeting", func() {
	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	It("should be show when cli is run", func() {
		path, err := Build("github.com/msgehard/golangTaskManager/main")
		Expect(err).NotTo(HaveOccurred())

		cli, err := Start(exec.Command(path), GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(cli).Should(Exit(0))
		Eventually(cli).Should(Say("Welcome to the Golang Task Manager!"))
	})
})
