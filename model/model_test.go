package model
import "testing"
func TestModelValidation(t *testing.T){if !NewRecord("r","oil","paint",2).Valid(){t.Fatal("invalid")};if AllowedTransition("archived","available"){t.Fatal("transition")}}
