// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package abov3_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/fajardofahad/abov3-sdk-go"
	"github.com/fajardofahad/abov3-sdk-go/internal/testutil"
	"github.com/fajardofahad/abov3-sdk-go/option"
)

func TestAgentListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Agent.List(context.TODO(), abov3.AgentListParams{
		Directory: abov3.F("directory"),
	})
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
