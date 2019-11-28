module GoLab

go 1.13

require (
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/garyburd/redigo v1.6.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/micro/go-micro v1.13.1 // indirect
	github.com/olzhy/quote v1.0.0
	github.com/satori/go.uuid v1.2.0
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.0.0-20191011234655-491137f69257
	google.golang.org/grpc v1.24.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

// replace "go.uber.org/zap" => "./vendor/go.uber.org/zap"
