package colony

import "fmt"

type MockSDK struct{}

func NewMockSDK() *MockSDK {
	return &MockSDK{}
}

func (s *MockSDK) SubmitWorkflow(specJSON []byte) error {
	fmt.Printf("[MockSDK] Simulating submission to ColonyOS...\n")
	fmt.Printf("[MockSDK] Payload check: %d bytes\n", len(specJSON))
	return nil
}
