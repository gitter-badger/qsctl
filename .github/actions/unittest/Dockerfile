FROM archlinux

LABEL Name=Unittest
LABEL Version=0.1.0

RUN pacman -Syu --noconfirm && pacman -S --noconfirm go git make

# Install required tools.
RUN go get -u golang.org/x/lint/golint
RUN go get -u github.com/golang/mock/mockgen

RUN make check
RUN make build
RUN make test
