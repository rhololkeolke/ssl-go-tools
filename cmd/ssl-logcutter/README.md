# ssl-logcutter

Cut and rename official log files as described in the [SSL wiki](http://wiki.robocup.org/Small_Size_League/Game_Logs).

The given log files will be shortened from the start of the first half to the start of the post game stage.

Team names and the timestamp of the first used frame will be used for the log file name.

## Installation

Use go-get to install this executable:

```
go get -u github.com/RoboCup-SSL/ssl-go-tools/cmd/ssl-logcutter
```

## Usage

The binary is called `ssl-logcutter`.
Run it with `-h` to print usage information.