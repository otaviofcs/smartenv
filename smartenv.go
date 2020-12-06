package smartenv

import "os"

func Env(name string, dflt ...string) string {
  value := os.Getenv(name)
  if(value != "") {
    return value
  }
  if(len(dflt) > 0) {
    return dflt[0]
  }
  return ""
}
