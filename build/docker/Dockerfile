FROM golang:latest AS build

WORKDIR /app

# download dependencies - low frequency of changes
COPY go.mod go.sum ./
RUN go mod download

# copy sources - higher frequency of changes
COPY internal/ internal/
COPY cmd/ cmd/
COPY api/ api/

# ensure tests are passing
RUN go test ./...

# create executable - hospital-webapi-srv
# we want to use scratch image so setting
# the build options in the way that will link all dependencies statically
RUN CGO_ENABLED=0 GOOS=linux \
      go build \
      -ldflags="-w -s" \
      -installsuffix 'static' \
      -o ./hospital-webapi-srv ./cmd/ambulance-api-service

############################################

FROM scratch

# see https://github.com/opencontainers/image-spec/blob/main/annotations.md for details
LABEL org.opencontainers.image.authors="Your Name"
LABEL org.opencontainers.image.title="Hospital Management WebAPI Service"
LABEL org.opencontainers.image.description="WEBAPI for managing hospital departments, beds, and patients"

# list all variables and their default values for clarity
ENV AMBULANCE_API_ENVIRONMENT=production
ENV AMBULANCE_API_PORT=8080
ENV AMBULANCE_API_MONGODB_HOST=mongo
ENV AMBULANCE_API_MONGODB_PORT=27017
ENV AMBULANCE_API_MONGODB_DATABASE=cv1sarsabsim-ambulance-wl
ENV AMBULANCE_API_MONGODB_COLLECTION=ambulance
ENV AMBULANCE_API_MONGODB_USERNAME=root
ENV AMBULANCE_API_MONGODB_PASSWORD=
ENV AMBULANCE_API_MONGODB_TIMEOUT_SECONDS=5

# to avoid connection errors in standalone case otel exporters are disabled by default
ENV OTEL_TRACES_EXPORTER=none
ENV OTEL_METRICS_EXPORTER=none
ENV OTEL_LOGS_EXPORTER=none
ENV LOG_LEVEL=info

ENV OTEL_SERVICE_NAME=ambulance-wl-api

COPY --from=build /app/hospital-webapi-srv ./

# Actual port may be changed during runtime
# Default using for the simple case scenario
EXPOSE 8080
ENTRYPOINT ["./hospital-webapi-srv"]