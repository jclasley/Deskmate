############################
# STEP 1 build executable binary
############################
FROM golang@sha256:0991060a1447cf648bab7f6bb60335d1243930e38420bee8fec3db1267b84cfa as zendesk


WORKDIR $GOPATH/src/github.com/circleci/Deskmate/graphql
COPY graphql .
# Fetch dependencies.
# Using go get.
RUN go mod download
RUN go mod verify
# Build the binary.
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/zendesk

##############################
# STEP 2 build out REST server
##############################

FROM golang as server

WORKDIR $GOPATH/src/github.com/circleci/Deskmate/server
COPY server .
# Fetch dependencies.
# Using go get.


# # need ssh key here to do this part -- it gets removed because this is an intermediate container
# # default value is the contents of a file where the access token is stored:
#     # in my case, this is "$HOME/pwds/Dockerfile-access"
# ARG ACCESS_TOKEN=$(cat $HOME/pwds/Dockerfile-access.txt)
# # use a github access token to access the repositories
# RUN git config --global url."https://$(ACCESS_TOKEN):@github.com/".insteadOf "https://github.com/"

RUN go get ./
# Build the binary.
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/server


#################################
# STEP 3 create permissioned user
#################################
FROM alpine AS user
# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001


# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

############################
# STEP 4 build a small image
############################
# Can't run from sratch, need to be able to `chmod`
FROM alpine

# Install curl
RUN apk update && apk --no-cache add curl

# Copy our static executable.
# Import from builder.
COPY --from=user /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=user /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=user /etc/passwd /etc/passwd
COPY --from=user /etc/group /etc/group

COPY --from=zendesk /go/bin/zendesk /go/bin/zendesk
COPY --from=server /go/bin/server /go/bin/server
# Run the zendesk binary.

# Use an unprivileged user.
USER appuser:appuser
EXPOSE 8090
EXPOSE 8080

# Copy the command to run both servers at startup
COPY serverStart.sh /scripts/serverStart.sh
RUN ["chmod", "+x", "/scripts/serverStart.sh"]

ENTRYPOINT ["/scripts/serverStart.sh"]