package integration_test

import (
	"log"
	"os"
	"testing"
)

var (
	testStatus = os.Getenv("TEST_INTEGRATION")
)

func TestMain(m *testing.M) {
	if testStatus != "true" {
		log.Println("[Integration Test] TEST_INTEGRATION os environment is not defined or not true, integration test skipped, define the env like 'TEST_INTEGRATION=true'")
		return
	}

	m.Run()
}
