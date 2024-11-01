FROM homebrew/ubuntu22.04:latest
RUN brew tap filebrowser/tap && brew install filebrowser
WORKDIR /app/
ENTRYPOINT ["filebrowser"]
CMD ["-a", "0.0.0.0", "-p", "8080", "-r", "/var/log/"]
