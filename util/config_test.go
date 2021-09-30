package util

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	fmt.Println("Starting Test")
	os.Setenv("API_BASE_URL", "testing")
	config, err := LoadConfig(".")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if strings.Compare(config.APIUrl, "testing") != 0 {
		fmt.Println("Config doesnt match env; expecting testing, got:", config.APIUrl)
		t.Fail()
	}
}
