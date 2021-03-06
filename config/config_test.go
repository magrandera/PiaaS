package config_test

import (
	"testing"

	"github.com/magrandera/SPaaS/config"
)

func TestReadConfig(t *testing.T) {
	v, err := config.ReadConfig("../test/test.json", map[string]interface{}{
		"test": true,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	if v.GetInt("port") != 9000 {
		t.Fatalf("expected port %v, but got %v", 9000, v.GetInt("port"))
	}
	if v.GetBool("test") != true {
		t.Fatalf("expected %v, but got %v", true, v.GetBool("test"))
	}
}
