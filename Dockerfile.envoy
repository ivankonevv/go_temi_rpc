FROM envoyproxy/envoy-dev:03d86fcde78d066a0d1dd7a48c8f1bb3d10859cf
RUN apt-get update
COPY envoy.yaml /etc/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy.yaml


#docker run -d -p 8080:8080 envoy:v1