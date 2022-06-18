package backoff

type BackOff interface {
	BackOffIteration() (int, bool)
}

type ServiceBackOff struct {
	BackOff
}
