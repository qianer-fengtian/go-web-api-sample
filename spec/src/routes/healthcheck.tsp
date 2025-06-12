import "@typespec/http";

using TypeSpec.Http;

namespace NoteService.Routes;

model HealthCheckResponse {
  status: "healthy";
  timestamp: utcDateTime;
  version?: string;
}

@route("/healthcheck")
namespace HealthCheck {
  @get
  @summary("Health check endpoint")
  @doc("Returns the health status of the API")
  op getHealthCheck(): HealthCheckResponse;
}