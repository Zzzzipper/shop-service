# Build stage
FROM golang AS build-env
ADD ./app /src/merchant-service
ENV CGO_ENABLED=0
RUN cd /src/merchant-service/cmd && go build -o /app

# Production stage
FROM scratch
COPY --from=build-env /app /
ENTRYPOINT ["/app"]
