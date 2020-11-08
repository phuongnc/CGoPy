FROM ubuntu:18.04

WORKDIR /app

# Update and upgrade repo
RUN apt-get update -y -q
RUN apt-get install -y curl wget nano htop

# Download Go 1.14.2 and install it to /usr/local/go
RUN curl -s https://storage.googleapis.com/golang/go1.14.2.linux-amd64.tar.gz| tar -v -C /usr/local -xz

# Let's people find our Go binaries
ENV PATH $PATH:/usr/local/go/bin
ENV LD_LIBRARY_PATH $LD_LIBRARY_PATH:/root/.pyenv/versions/3.8.0/lib:/app

# Set environmental variables
ENV PYENV_ROOT="/root/.pyenv" \
	PATH="/root/.pyenv/shims:/root/.pyenv/bin:${PATH}" \
	PIPENV_YES=1 \
	PIPENV_DONT_LOAD_ENV=1 \
	LC_ALL="C.UTF-8" \
	LANG="en_US.UTF-8"

# Install pyenv
RUN apt-get install -y git mercurial build-essential libssl-dev libbz2-dev zlib1g-dev libffi-dev libreadline-dev libsqlite3-dev curl && \
curl -L https://raw.githubusercontent.com/yyuu/pyenv-installer/master/bin/pyenv-installer | bash

RUN PYTHON_CONFIGURE_OPTS="--enable-shared" pyenv install 3.8.0
RUN pyenv global 3.8.0
RUN pyenv rehash

ENV PYTHONUNBUFFERED 1

COPY . .

EXPOSE 8080

RUN make build

CMD ["make", "run"]


