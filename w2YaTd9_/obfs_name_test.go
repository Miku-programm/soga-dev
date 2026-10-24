package w2YaTd9_

import "testing"

func TestOriginalObfsNames(t *testing.T) {
	for input, want := range map[string]OdAX8q{"": "", "plain": "", "HTTP": "http", "TLS": "tls", "Simple_Obfs_Http": "http", "SIMPLE_OBFS_TLS": "tls", " http": "", "unknown": ""} {
		if got := GBXRYZ(input); got != want {
			t.Fatal(input, got, want)
		}
	}
}
