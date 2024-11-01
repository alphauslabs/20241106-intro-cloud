FROM homebrew/ubuntu22.04:latest
RUN brew install ttyd
RUN sudo apt-get update && sudo apt-get install telnet -y
WORKDIR /app/
ENTRYPOINT ["ttyd"]
CMD ["-W", "-p", "8080", "telnet", "mapscii.me"]
