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

func TestTuiAppendPromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.AppendPrompt(context.TODO(), abov3.TuiAppendPromptParams{
		Text:      abov3.F("text"),
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

func TestTuiClearPromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.ClearPrompt(context.TODO(), abov3.TuiClearPromptParams{
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

func TestTuiExecuteCommandWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.ExecuteCommand(context.TODO(), abov3.TuiExecuteCommandParams{
		Command:   abov3.F("command"),
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

func TestTuiOpenHelpWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.OpenHelp(context.TODO(), abov3.TuiOpenHelpParams{
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

func TestTuiOpenModelsWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.OpenModels(context.TODO(), abov3.TuiOpenModelsParams{
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

func TestTuiOpenSessionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.OpenSessions(context.TODO(), abov3.TuiOpenSessionsParams{
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

func TestTuiOpenThemesWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.OpenThemes(context.TODO(), abov3.TuiOpenThemesParams{
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

func TestTuiShowToastWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.ShowToast(context.TODO(), abov3.TuiShowToastParams{
		Message:   abov3.F("message"),
		Variant:   abov3.F(abov3.TuiShowToastParamsVariantInfo),
		Directory: abov3.F("directory"),
		Title:     abov3.F("title"),
	})
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTuiSubmitPromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.SubmitPrompt(context.TODO(), abov3.TuiSubmitPromptParams{
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
