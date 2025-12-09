package colony

import "fmt"

type ColonyClient struct {
	sdk ColonySDK
}

func NewColonyClient(sdk ColonySDK) *ColonyClient {
	return &ColonyClient{
		sdk: sdk,
	}
}

func (c *ColonyClient) SubmitWorkflow(specJSON []byte) error {
	// Here we would handle specific ColonyOS logic if needed (e.g. wrapping response)
	// For now, just pass through to SDK
	return c.sdk.SubmitWorkflow(specJSON)
}

func (c *ColonyClient) RegisterFunction(specJSON []byte) error {
	// Not supported in SDK interface yet, stubbing
	fmt.Printf("[ColonyClient] Registering function not implemented in SDK yet.\n")
	return nil
}
