# Drop TCP packets on a specific port

Dropping TCP packets on a specific port using eBPF.

## Pre-requisites

1. A relatively recent version of the Linux kernel (>= 4.4) for full eBPF support.
2. The `clang` compiler - `sudo apt-get -y install clang`.
3. The `llvm` tools package - `sudo apt-get -y install llvm-14-tools`.
4. The `libbpf` and `bpftool` library - `sudo apt-get -y install libbpf-dev bpftool`.
5. `go` - The go programming language compiler and tools.

## Running the eBPF program 

1. Clone the repo: 
    
    `git clone https://github.com/navin772/go-ebpf-accuknox.git`.

2. Compile the eBPF program using clang:

    `clang -O2 -g -target bpf -c bpf/xdp_prog.c -o bpf/xdp_prog.o`

3. Compile the go program:

    `sudo go run main.go -port 4040`

    or
    
    `sudo $(which go) run main.go -port 4040`

    ```
    Note: Replace the port on which you want to drop the packets, the default is 4040.
    ```

## Verify the eBPF program
Create 2 new terminal sessions - 1 and 2.

```
Note: The go program is configured to drop packets on the loopback (lo) interface, hence the packets will be dropped only if the server is running on the same machine (localhost).
```

1. Start a python (or any other) http server on port `4040` on terminal 1:

    `python3 -m http.server 4040`

2. Try to access the server from terminal 2:

    `curl --max-time 10 http://localhost:4040`

3. No connection will be made to the server and the connection will timeout, hence the TCP packets are being succesfully dropped.
4. Stop the running eBPF program by pressing `Ctrl+C` and verify that the `curl` command works as expected.

## Demo: Dropping TCP Packets

![Demo Drop TCP Packets](https://github.com/navin772/go-ebpf-accuknox/raw/main/demo_drop_TCP_packets.mp4)

[Click here to view the video](https://github.com/navin772/go-ebpf-accuknox/raw/main/demo_drop_TCP_packets.mp4)