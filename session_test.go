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

func TestSessionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.New(context.TODO(), abov3.SessionNewParams{
		Directory: abov3.F("directory"),
		ParentID:  abov3.F("parentID"),
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

func TestSessionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Update(
		context.TODO(),
		"id",
		abov3.SessionUpdateParams{
			Directory: abov3.F("directory"),
			Title:     abov3.F("title"),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.List(context.TODO(), abov3.SessionListParams{
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

func TestSessionDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Delete(
		context.TODO(),
		"id",
		abov3.SessionDeleteParams{
			Directory: abov3.F("directory"),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionAbortWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Abort(
		context.TODO(),
		"id",
		abov3.SessionAbortParams{
			Directory: abov3.F("directory"),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionChildrenWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Children(
		context.TODO(),
		"id",
		abov3.SessionChildrenParams{
			Directory: abov3.F("directory"),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionCommandWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Command(
		context.TODO(),
		"id",
		abov3.SessionCommandParams{
			Arguments: abov3.F("arguments"),
			Command:   abov3.F("command"),
			Directory: abov3.F("directory"),
			Agent:     abov3.F("agent"),
			MessageID: abov3.F("msgJ!"),
			Model:     abov3.F("model"),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Get(
		context.TODO(),
		"id",
		abov3.SessionGetParams{
			Directory: abov3.F("directory"),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionInitWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Init(
		context.TODO(),
		"id",
		abov3.SessionInitParams{
			MessageID:  abov3.F("messageID"),
			ModelID:    abov3.F("modelID"),
			ProviderID: abov3.F("providerID"),
			Directory:  abov3.F("directory"),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionMessageWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Message(
		context.TODO(),
		"id",
		"messageID",
		abov3.SessionMessageParams{
			Directory: abov3.F("directory"),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionMessagesWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Messages(
		context.TODO(),
		"id",
		abov3.SessionMessagesParams{
			Directory: abov3.F("directory"),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionPromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Prompt(
		context.TODO(),
		"id",
		abov3.SessionPromptParams{
			Parts: abov3.F([]abov3.SessionPromptParamsPartUnion{abov3.TextPartInputParam{
				Text:      abov3.F("text"),
				Type:      abov3.F(abov3.TextPartInputTypeText),
				ID:        abov3.F("id"),
				Synthetic: abov3.F(true),
				Time: abov3.F(abov3.TextPartInputTimeParam{
					Start: abov3.F(0.000000),
					End:   abov3.F(0.000000),
				}),
			}}),
			Directory: abov3.F("directory"),
			Agent:     abov3.F("agent"),
			MessageID: abov3.F("msgJ!"),
			Model: abov3.F(abov3.SessionPromptParamsModel{
				ModelID:    abov3.F("modelID"),
				ProviderID: abov3.F("providerID"),
			}),
			System: abov3.F("system"),
			Tools: abov3.F(map[string]bool{
				"foo": true,
			}),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionRevertWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Revert(
		context.TODO(),
		"id",
		abov3.SessionRevertParams{
			MessageID: abov3.F("msgJ!"),
			Directory: abov3.F("directory"),
			PartID:    abov3.F("prtJ!"),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionShareWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Share(
		context.TODO(),
		"id",
		abov3.SessionShareParams{
			Directory: abov3.F("directory"),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionShellWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Shell(
		context.TODO(),
		"id",
		abov3.SessionShellParams{
			Agent:     abov3.F("agent"),
			Command:   abov3.F("command"),
			Directory: abov3.F("directory"),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionSummarizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Summarize(
		context.TODO(),
		"id",
		abov3.SessionSummarizeParams{
			ModelID:    abov3.F("modelID"),
			ProviderID: abov3.F("providerID"),
			Directory:  abov3.F("directory"),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionUnrevertWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Unrevert(
		context.TODO(),
		"id",
		abov3.SessionUnrevertParams{
			Directory: abov3.F("directory"),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionUnshareWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Unshare(
		context.TODO(),
		"id",
		abov3.SessionUnshareParams{
			Directory: abov3.F("directory"),
		},
	)
	if err != nil {
		var apierr *abov3.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
