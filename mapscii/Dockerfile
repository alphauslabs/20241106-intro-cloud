FROM homebrew/ubuntu22.04:latest
RUN brew install ttyd
WORKDIR /app/
ENTRYPOINT ["ttyd"]
CMD ["-W", "telnet", "mapsii.me"]
