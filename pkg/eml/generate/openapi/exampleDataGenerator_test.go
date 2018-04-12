package openapi

import (
	"testing"
)

func TestMustCreateGUIDForStreamType(t *testing.T) {
	streamID := parameterGenerator("invoiceId", "invoice")
	if len(streamID) == 0 {
		t.Error("expected a UUID")
	}
	streamID2 := parameterGenerator("invoiceId", "invoice")
	if streamID != streamID2 {
		t.Error("expected no new stream ID UUID to be generated for existing stream type")
	}

}
