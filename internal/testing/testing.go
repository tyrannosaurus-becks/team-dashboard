package testing

import (
	"os"
	"testing"
)

const envVarIsAcc = "DASHBOARD_ACC"

func RequireAcceptanceTestFlag(t *testing.T) {
	if os.Getenv(envVarIsAcc) == "" {
		t.Logf("skipped because %s is unset", envVarIsAcc)
		t.SkipNow()
	}
}
