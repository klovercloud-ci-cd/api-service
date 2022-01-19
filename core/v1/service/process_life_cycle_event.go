package service


// ProcessLifeCycleEvent Process Life Cycle Event operations.
type ProcessLifeCycleEvent interface {
	Store(events interface{}) (httpCode int, error error)
	PullNonInitializedAndAutoTriggerEnabledEventsByStepType(count , stepType string) (httpCode int, body interface{})
	PullPausedAndAutoTriggerEnabledResourcesByAgentName(count , agent string) (httpCode int, body interface{})
}
