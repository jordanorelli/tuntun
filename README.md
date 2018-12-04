tuntun
===

this is a proof of concept proxy music messaging server. tuntun acts as a
bridge between a udp message and a tcp socket. note that tcp is not in general
recommended for music applications of this sort; this project targets a
specific piece of hardware that only speaks tcp. if you have the option of
avoiding tcp for anything latency or time-sensitive, avoid it in favor of udp.

anyway, usage:

  tuntun [--listen addr] [--forward addr]

starting tuntun listens for incoming OSC messages on the provided listen
address, and forwards them to a tcp server on the forward address.

the one understood message at the time of reading is as follows:

  /example,i N

(where N is some integer)

this can be verified using a combination of oscsend and netcat:
- open three terminals
- in terminal one, start tuntun
- in terminal two, start a netcat tcp listener. i recommend doing this in a
  loop since netcat will only process one tcp socket, then quit.
    nc -l 9225
- in terminal three, send an OSC message to tuntun with oscsend:
    oscsend localhost 9220 /example i 15
- you should see "15" printed in terminal two, and some debug info in terminal
  one.

there's also a chuck client example provided, in ex/ex.ck

installing
====

there's a release in the releases folder. You can untar that and put the
binary anywhere on your PATH and you'll be good to go. I strong recommend
putting the binary on your PATH so that you can invoke it from ChucK directly.

building
====
if you want to build from source (maybe you're not on Linux and can't use that
release for example), all you need is a Go compiler and the OSC library.
- [Install Go](https://golang.org/dl/) 
- install the Go OSC library with `go get github.com/hypebeast/go-osc/osc`
- build and install the binary to your $GOPATH with `go install` -or- build
  the binary in the current directory with `go build` and move it to your
  $PATH manually.
