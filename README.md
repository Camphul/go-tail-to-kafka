# go-tail-to-kafka
Tail log files and submit them to kafka in go.

All written in one file. Compatible with kafka >= 1.0.0

## Build

```
git clone https://github.com/Camphul/go-tail-to-kafka.git
cd go-tail-to-kafka
go build main.go
```

## CLI options

- `v` enable verbose printing
- `brokers <brokers>` comma-separated kafka brokers
- `topic <name>` topic to write to 
- `tail <filename>` file to tail
- `maxRetry <amt>` kafka max retry limit

# Example

Please build first.

```
./main -v -brokers localhost:9092,localhost:9093 -topic myapplogs -tail myapplication.log
```
