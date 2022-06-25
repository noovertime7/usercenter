FROM alpine
ADD usercenter-service /usercenter-service
ENTRYPOINT [ "/usercenter-service" ]
