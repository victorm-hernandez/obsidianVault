package health

type LocationGenerationHealthData struct {
	Generator               *InstrumentedGenerator
	OutgoingReqRoundTripper *InstrumentedRoundTripper
}
