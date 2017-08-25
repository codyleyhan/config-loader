package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	if err := os.Setenv("ENV", "test"); err != nil {
		t.Log(err)
		t.Fail()
	}

	anotherConfig := LoadConfig("./test_folder", "another")
	if anotherConfig.FileName() != "another.test.json" {
		t.Error("Expected another.test.json, got", anotherConfig.FileName())
	}

	if err := os.Setenv("ENV", "DEV"); err != nil {
		t.Log(err)
		t.Fail()
	}

	devConfig := LoadConfig("./test_folder", "config")
	if devConfig.FileName() != "config.dev.json" {
		t.Error("Expected config.dev.json, got", devConfig.FileName())
	}
}

func TestLoaderGet(t *testing.T) {
	if err := os.Setenv("ENV", "test"); err != nil {
		t.Log(err)
		t.Fail()
	}

	config := LoadConfig("./test_folder", "config")

	if config.Get("hello") != "world" {
		t.Error("Expected world, got", config.Get("hello"))
	}

	if config.Get("nested.value") != "good" {
		t.Error("Expected good, got", config.Get("nested.value"))
	}

	if config.Get("nested.more.nest") != "" {
		t.Error("Expected \"\", got", config.Get("nested.more.nest"))
	}
}

func TestLoaderGetBool(t *testing.T) {
	if err := os.Setenv("ENV", "test"); err != nil {
		t.Log(err)
		t.Fail()
	}

	config := LoadConfig("./test_folder", "config")

	if !config.GetBool("nested.more.nest") {
		t.Error("Expected nested.more.nest to be true")
	}

	if config.GetBool("nested.nothing") {
		t.Error("Expected nested.nothing to be false")
	}
}

func TestLoaderGetFloat(t *testing.T) {
	if err := os.Setenv("ENV", "test"); err != nil {
		t.Log(err)
		t.Fail()
	}

	config := LoadConfig("./test_folder", "config")

	if config.GetFloat("pie") != 3.14 {
		t.Error("Expected pie to be 3.14, got", config.GetFloat("pie"))
	}
}

func TestLoaderGetInt(t *testing.T) {
	if err := os.Setenv("ENV", "test"); err != nil {
		t.Log(err)
		t.Fail()
	}

	config := LoadConfig("./test_folder", "config")

	if config.GetInt("universe") != 42 {
		t.Error("Expected pie to be 42, got", config.GetInt("universe"))
	}

	if config.GetInt("nested.nothing") != 0 {
		t.Error("Expected pie to be 0, got", config.GetInt("nested.nothing"))
	}
}

func TestLoaderGetUInt(t *testing.T) {
	if err := os.Setenv("ENV", "test"); err != nil {
		t.Log(err)
		t.Fail()
	}

	config := LoadConfig("./test_folder", "config")

	if config.GetUint("universe") != 42 {
		t.Error("Expected pie to be 42, got", config.GetUint("universe"))
	}

	if config.GetUint("nested.nothing") != 0 {
		t.Error("Expected pie to be 0, got", config.GetUint("nested.nothing"))
	}
}
