package colony

import "fmt"

type MockColonyClient struct{}

func NewMockColonyClient() *MockColonyClient {
	return &MockColonyClient{}
}

func (c *MockColonyClient) SubmitWorkflow(specJSON []byte) error {
	fmt.Printf("[ColonyOS] Submitting Workflow:\n%s\n", string(specJSON))
	return nil
}

func (c *MockColonyClient) RegisterFunction(specJSON []byte) error {
	fmt.Printf("[ColonyOS] Registering Function:\n%s\n", string(specJSON))
	return nil
}
