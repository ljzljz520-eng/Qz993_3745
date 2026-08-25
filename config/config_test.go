package config
import "testing"
func TestConfig(t *testing.T){c:=Default();if c.Limit(0)!=1||c.Limit(1000)!=100{t.Fatal(c)}}
