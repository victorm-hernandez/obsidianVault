package handlers

import (
	"TakeHomeAssessment/api"
	"TakeHomeAssessment/internal/health"
	"encoding/json"
	"net/http"
)

func GetStatsHandler(locGenData *health.LocationGenerationHealthData) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		genStatus := locGenData.Generator.GetStatus()
		outReqStats := locGenData.OutgoingReqRoundTripper.GetStats()

		response := api.HealthResponse{
			GeneratorStatus:   genStatus,
			OutgoingReqStatus: outReqStats,
		}

		responseBytes, err := json.Marshal(response)

		if err != nil {
			http.Error(w, "Error generating response JSON", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "json/application-text")
		w.Write(responseBytes)

	})
}
