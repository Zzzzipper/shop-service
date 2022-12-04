# Build stage
FROM golang AS build-env
ADD ./app /src/merchant-service
ENV CGO_ENABLED=0
ARG GITLAB_DEPENDENCIES_TOKEN
RUN git config --global url."https://oauth2:${GITLAB_DEPENDENCIES_TOKEN}@gitlab.mapcard.pro/external-map-team/api-proto".insteadOf https://gitlab.mapcard.pro/external-map-team/api-proto
RUN go env -w GOPRIVATE=gitlab.mapcard.pro/external-map-team/*
RUN cd /src/merchant-service/cmd && go build -o /app

# Production stage
FROM scratch
COPY --from=build-env /app /
ENTRYPOINT ["/app"]
