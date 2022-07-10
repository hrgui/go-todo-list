FROM cosmtrek/air

WORKDIR /app

COPY . ./

EXPOSE 3000

CMD air
