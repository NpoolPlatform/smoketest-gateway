FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY smoketest-gateway /app

ENTRYPOINT ["/app/smoketest-gateway"]

EXPOSE 50760 50761 50762 50763

CMD ["run"]
