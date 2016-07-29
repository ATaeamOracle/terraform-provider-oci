package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

// This test runs Provider sanity checks.
func TestProvider(t *testing.T) {
	// Real client for the sanity check. Makes this more of an acceptance test.
	client := &Client{}
	if err := Provider(client).(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
