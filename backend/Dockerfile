FROM --platform=${BUILDPLATFORM} golang:1.23-alpine AS build

WORKDIR /src
COPY . .

RUN mkdir /src/binary

ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /src/binary cmd/main.go

FROM scratch AS bin
COPY --from=build /src/binary /
COPY --from=build /src/static /static

CMD ["/main"]
