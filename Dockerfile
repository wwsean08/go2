FROM scratch

ENTRYPOINT ["/go2"]
CMD ["version"]
ADD bin/linux_amd64/go2 /
