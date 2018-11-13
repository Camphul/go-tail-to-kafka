# go-tail-to-kafka
Tail log files and submit them to kafka in go.

Simply 1file script.

## Build

```
go build main.go
```

## CLI options

- `v` enable verbose printing
- `brokers <brokers****>` comma-separated kafka brokers
- `topic <name>` topic to write to 
- `tail <filename>` file to tail
- `maxRetry <amt>` kafka max retry limit
