## GO Kafka Example
This is example repository how to use kafka in golang, with sarama, goavro, and gogen avro

- plain : example publish and consume using string
- avro : example publish and consume using Avro data. to generate avro schame -> struct, please go folder `avro/schema` and run following command
```gotemplate
 gogen-avro . test.avsc
```

### Run App
#### plain
will send/publish 10 message 
- producer : please go folder `cmd/plain/producer` and run command `go run main.go`
- consumer : please go folder `cmd/plain/consumer` and run command `go run main.go`

#### avro
will send/publish avro data (name, age)
- producer : please go folder `cmd/avro/producer` and run command `go run main.go`
- consumer : please go folder `cmd/avro/consumer` and run command `go run main.go`
