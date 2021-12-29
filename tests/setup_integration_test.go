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

const baseURL = "http://localhost:8090"

func TestMain(m *testing.M) {
	currDir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

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
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("healthcheck successful")
	m.Run()

	cancel()
	os.Remove(binPath)
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
	cmd := exec.CommandContext(ctx, appPath)

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
