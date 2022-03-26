//go:build integration

package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

const (
	baseURL = "http://localhost:8080"
	dbName  = "fintra.test.db"
)

func TestMain(m *testing.M) {
	currDir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

	os.Remove(binPath)
	os.Remove(dbName)

	binPath := filepath.Join(currDir, "api.test")

	if err := buildApp(binPath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		startApp(ctx, binPath)
	}()

	if err := healthcheck(); err != nil {
		fmt.Println("healthcheck failed:", err)
		os.Remove(binPath)
		os.Remove(dbName)
		os.Exit(1)
	}

	fmt.Println("healthcheck successful")
	m.Run()

	cancel()
	os.Remove(binPath)
	os.Remove(dbName)
}

func buildApp(binPath string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "build", "-o", binPath, "../cmd/...")
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%s: %w", out, err)
	}

	return nil
}

func startApp(ctx context.Context, appPath string) error {
	os.Setenv("DB_NAME", dbName)
	cmd := exec.CommandContext(ctx, appPath, "-port", "8080")

	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%s: %w", out, err)
	}
	return nil
}

func healthcheck() error {
	ticker := time.NewTicker(time.Millisecond * 50)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var err error
	var resp *http.Response

	for {
		select {
		case <-ctx.Done():
			if resp == nil {
				return fmt.Errorf("no response from server")
			}
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("status returned: %d", resp.StatusCode)
			}
			return err
		case <-ticker.C:
			resp, err = http.Get(baseURL + "/healthcheck")
			if err == nil && resp.StatusCode == http.StatusOK {
				return nil
			}
		}
	}
}
