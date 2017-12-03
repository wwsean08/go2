FROM scratch

EXPOSE 4646
ENTRYPOINT ["/go2"]

ADD etc/config-defaults.yaml /etc/go2/config.yaml
ADD bin/linux_amd64/go2 /
