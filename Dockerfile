# Build stage
FROM golang AS build-env
ADD . /src/merchant-service
ENV CGO_ENABLED=0
RUN cd /src/merchant-service/cmd && go build -o /app

# Production stage
FROM scratch
COPY --from=build-env /app /
ARG DB_USER
ARG DB_PASSWORD
ARG DB_HOST
ARG DB_PORT
ARG DB_NAME
ENV POSTGRES_URL=postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable
ENTRYPOINT ["/app"]
