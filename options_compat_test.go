package redis

import "testing"

func TestOptionsDefaultToRESP2ForMaintNotificationsCompatibility(t *testing.T) {
	client := NewClient(&Options{Addr: "127.0.0.1:1"})
	defer client.Close()

	if got := client.Options().Protocol; got != 2 {
		t.Fatalf("default protocol = %d, want RESP2 to avoid unsupported CLIENT MAINT_NOTIFICATIONS handshakes", got)
	}
	if client.Options().MaintNotificationsConfig == nil {
		t.Fatal("maintnotifications config should still be defaulted so explicit RESP3 clients keep the feature available")
	}
}
