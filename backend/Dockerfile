FROM debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

COPY bin/planner_server /usr/bin/planner_server

CMD ["planner_server"]
