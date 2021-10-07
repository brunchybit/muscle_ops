FROM ubuntu:21.10 as builder
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y wget gnupg
RUN wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -

RUN echo "deb http://apt.postgresql.org/pub/repos/apt/ `lsb_release -cs`-pgdg main" | tee  /etc/apt/sources.list.d/pgdg.list

RUN apt-get install -y \
    pgbackrest \
    postgresql-13 \

