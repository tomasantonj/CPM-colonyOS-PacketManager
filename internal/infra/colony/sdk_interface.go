package colony

// ColonySDK abstracts the ColonyOS client operations
type ColonySDK interface {
	SubmitWorkflow(specJSON []byte) error
}
