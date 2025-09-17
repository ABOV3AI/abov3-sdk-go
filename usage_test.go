// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package abov3_test

import (
	"context"
	"os"
	"testing"

	"github.com/fajardofahad/abov3-sdk-go"
	"github.com/fajardofahad/abov3-sdk-go/internal/testutil"
	"github.com/fajardofahad/abov3-sdk-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := abov3.NewClient(
		option.WithBaseURL(baseURL),
	)
	sessions, err := client.Session.List(context.TODO(), abov3.SessionListParams{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", sessions)
}
