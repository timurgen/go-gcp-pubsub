FROM iron/base

WORKDIR /app

ARG binaryname
ENV runnable=$binaryname
COPY $binaryname /app/

ENTRYPOINT ./$runnable