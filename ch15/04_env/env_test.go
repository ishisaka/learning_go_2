package env

import (
	"testing"
)

func TestEnvVarProcess(t *testing.T) {
	// テストの中だけで環境変数を設定
	t.Setenv("OUTPUT_FORMAT", "JSON")
	cfg := ProcessEnvVars()
	if cfg.OutputFormat != "JSON" {
		t.Error("OutputFormat not set correctly")
	}
	// value of OUTPUT_FORMAT is reset when the test function exits
}
