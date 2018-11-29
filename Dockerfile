FROM iron/base

WORKDIR /app

ARG binaryname

COPY $binaryname /app/

ENTRYPOINT ["./$binaryname"]