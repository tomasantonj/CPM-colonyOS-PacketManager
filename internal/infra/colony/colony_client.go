package colony

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ColonyClient struct {
	serverHost string
	serverPort int
	colonyID   string
	prvKey     string
	httpClient *http.Client
}

func NewColonyClient(host string, port int, colonyID, prvKey string) *ColonyClient {
	return &ColonyClient{
		serverHost: host,
		serverPort: port,
		colonyID:   colonyID,
		prvKey:     prvKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *ColonyClient) SubmitWorkflow(specJSON []byte) error {
	url := fmt.Sprintf("http://%s:%d/api/workflows", c.serverHost, c.serverPort)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(specJSON))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Colony-ID", c.colonyID)

	// Sign payload if private key is present
	if c.prvKey != "" {
		signature, err := c.sign(specJSON)
		if err != nil {
			return fmt.Errorf("failed to sign payload: %w", err)
		}
		req.Header.Set("X-Colony-Signature", signature)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to submit workflow: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("server returned error %d: %s", resp.StatusCode, string(body))
	}

	fmt.Println("[ColonyClient] Workflow submitted successfully")
	return nil
}

func (c *ColonyClient) RegisterFunction(specJSON []byte) error {
	// TODO: Implement real registration
	fmt.Println("[ColonyClient] Registering function to", c.serverHost)
	return nil
}

func (c *ColonyClient) sign(msg []byte) (string, error) {
	// Assuming prvKey is hex encoded 64-byte Ed25519 private key
	// Note: Ed25519 private key is usually 64 bytes (seed + public key) or 32 bytes (seed).
	// Go's crypto/ed25519 uses 64 bytes.

	keyBytes, err := hex.DecodeString(c.prvKey)
	if err != nil {
		return "", fmt.Errorf("invalid private key hex: %w", err)
	}

	if len(keyBytes) != ed25519.PrivateKeySize {
		return "", fmt.Errorf("invalid private key length: got %d, want %d", len(keyBytes), ed25519.PrivateKeySize)
	}

	sig := ed25519.Sign(ed25519.PrivateKey(keyBytes), msg)
	return hex.EncodeToString(sig), nil
}
