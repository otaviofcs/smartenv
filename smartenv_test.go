package smartenv

import (
	"testing"
	"os"
)

func TestEnvWithDefault(t *testing.T) {
  t.Run("Check unsetted env without default", func(t *testing.T) {
    os.Unsetenv("RANDOM_ENV")
    random_env := Env("RANDOM_ENV")
    if random_env != "" {
      t.Errorf("Should be an empty env var. Got: %s", random_env)
    }
  })
  t.Run("Check unsetted env with default", func(t *testing.T) {
    os.Unsetenv("RANDOM_ENV")
    random_env := Env("RANDOM_ENV", "my default")
    if random_env != "my default" {
      t.Errorf("Should return default env 'my default'. Got: %s", random_env)
    }
  })
  t.Run("Return env value despite of existing default", func(t *testing.T) {
    os.Setenv("RANDOM_ENV", "my env")
    random_env := Env("RANDOM_ENV", "my default")
    if random_env != "my env" {
      t.Errorf("Should return env value 'my env'. Got: %s", random_env)
    }
  })
}
