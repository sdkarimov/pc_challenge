package backoff

type BackOff interface {
	BackOffIteration() (bool, int)
}

type ServiceBackOff struct {
	BackOff
}
