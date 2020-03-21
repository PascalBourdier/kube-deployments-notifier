package config

import (
	"fmt"
	"testing"
)

const nonExistentPath = "\\/hopefully/non/existent/path"

func TestConfig(t *testing.T) {
	conf := FakeConfig()
	if conf == nil {
		t.Error("Failed to initialize a fake config object")
	}

	// test config's provided FakeClientSet we'll use throughout this file
	cs := FakeClientSet()
	if fmt.Sprintf("%T", cs) != "*fake.Clientset" {
		t.Errorf("FakeClientSet() failed")
	}

	// test with the fake clientset (should panic on error)
	err := conf.Init("", "")
	if err != nil {
		t.Errorf("Failed to initialize conf: %+v", err)
	}
	if fmt.Sprintf("%T", conf.ClientSet) != "*fake.Clientset" {
		t.Errorf("conf.Init() shouldn't overwrite an existing ClientSet")
	}
}
