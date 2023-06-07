FROM    golang:alpine as builder
WORKDIR /app
ADD . /app/
RUN go build -o booklib

FROM    alpine
WORKDIR /app
RUN apk add --no-cache bash
COPY    --from=builder /app/booklib .
COPY    --from=builder /app/static/* ./static/
COPY    --from=builder /app/templates/* ./templates/
COPY    --from=builder /app/app.env .
CMD [ "./booklib" ]