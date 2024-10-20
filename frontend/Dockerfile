FROM --platform=${BUILDPLATFORM} golang:1.23-alpine AS build

WORKDIR /src
COPY . .

RUN mkdir /src/bin

ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /src/bin frontend/cmd/frontend/main.go

FROM scratch AS bin
COPY --from=build /src/bin /
COPY --from=build /src/static /static

CMD ["/main"]
