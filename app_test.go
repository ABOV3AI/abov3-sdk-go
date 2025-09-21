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

func TestAppLogWithOptionalParams(t *testing.T) {
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
	_, err := client.App.Log(context.TODO(), abov3.AppLogParams{
		Level:     abov3.F(abov3.AppLogParamsLevelDebug),
		Message:   abov3.F("message"),
		Service:   abov3.F("service"),
		Directory: abov3.F("directory"),
		Extra: abov3.F(map[string]interface{}{
			"foo": "bar",
		}),
	})
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppProvidersWithOptionalParams(t *testing.T) {
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
	_, err := client.App.Providers(context.TODO(), abov3.AppProvidersParams{
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
