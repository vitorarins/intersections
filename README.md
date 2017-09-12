<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [intersections](#intersections)
    - [Installation](#installation)
    - [Running](#running)
    - [Output](#output)
    - [Development](#development)

<!-- markdown-toc end -->
# intersections
This is a simple application, written in [Golang](https://golang.org/), that simulates a set of traffic lights at an intersection.

The traffic lights are designated (N, S) and (E, W) like a compass.

## Installation

If you have [Go](https://golang.org/doc/install), go get it:
```sh
go get github.com/vitorarins/intersections
```

If you don't want to have Go installed on your system you can download the executable binary from the [Releases](https://github.com/vitorarins/intersections/releases) page.

Just choose the latest version and click one of the Downloads links according to your Operating System. Currently there are versions for *Linux*, *MacOS* and *Windows*. All versions are for 64bit systems.

## Running

After you've installed the `intersections` binary, you just have to run it.
If it's in your `PATH` environment variable then you can just call it:
```sh
[user@pc]$ intersections
```

If it's not in your `PATH` then just give the full path of where you saved it:
```sh
[user@pc]$ ./path/to/intersections
```

## Output

The output of this application is a set of *strings* showing the time the lights will change, starting at `00:00.00`. Whenever the right time has passed, the lights changing event is reported.

Example:
```
(N, S): Green, (E, W): Red - 00:00.00
(N, S): Yellow, (E, W): Red - 04:30.00
(N, S): Red, (E, W): Green - 05:00.00
(N, S): Red, (E, W): Yellow - 09:30.00
(N, S): Green, (E, W): Red - 10:00.00
```

This goes on for 30 minutes.

## Development

If you want to look and change the code, you can clone, build and run this repository:

```sh
[user@pc]$ git clone https://github.com/vitorarins/intersections.git
[user@pc]$ cd intersections
[user@pc]$ go build
[user@pc]$ ./intersections
```

The code was made trying to keep things simple and easy to change. So for example, if you want to change the time duration of color switching in the lights:

 - open `main.go`
 - change the `colorDurations` variable
 - `go build`it again
 - run it again: `./intersections`
