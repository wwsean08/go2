FROM scratch

EXPOSE 4646
ENTRYPOINT ["/go2"]
CMD ["run"]

ADD bin/linux_amd64/go2 /
