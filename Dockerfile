FROM golang AS build

WORKDIR /app

COPY . .

RUN make build

FROM alpine

COPY --from=build /app/argo-cd-demo /argo-cd-demo

CMD ["/argo-cd-demo"]

EXPOSE 8888
