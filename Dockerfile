FROM iron/base

COPY pubsubservice /opt/service/

WORKDIR /opt/service

RUN chmod +x /opt/service/pubsubservice

EXPOSE 8000:8000

CMD /opt/service/pubsubservice
