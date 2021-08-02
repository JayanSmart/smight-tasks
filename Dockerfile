FROM balenalib/%%BALENA_MACHINE_NAME%%-golang:latest-build AS build

WORKDIR /go/src/github.com/JayanSmart/smight-tasks/app

RUN apt install tree

COPY /app ./
COPY /go.mod ./

RUN tree

RUN go mod download

RUN go build

FROM balenalib/%%BALENA_MACHINE_NAME%%-debian:stretch

COPY --from=build /go/src/github.com/JayanSmart/smight-tasks/app/ .

CMD ./app