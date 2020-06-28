module github.com/geekymedic/neon

go 1.12

require (
	github.com/aliyun/aliyun-oss-go-sdk v2.1.0+incompatible
	github.com/beanstalkd/go-beanstalk v0.0.0-20190515041346-390b03b3064a
	github.com/coreos/etcd v3.3.20+incompatible // indirect
	github.com/fortytw2/leaktest v1.3.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-playground/locales v0.12.1 // indirect
	github.com/go-playground/universal-translator v0.16.0 // indirect
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/leodido/go-urn v1.1.0 // indirect
	github.com/mailru/easyjson v0.0.0-20190626092158-b2ccc519800e // indirect
	github.com/olivere/elastic v6.2.21+incompatible
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/prometheus/client_golang v1.1.0
	github.com/shamaton/msgpack v1.1.1
	github.com/shima-park/agollo v1.1.7
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.5.0
	github.com/stretchr/testify v1.4.0
	github.com/ugorji/go v1.1.7 // indirect
	github.com/zentures/cityhash v0.0.0-20131128155616-cdd6a94144ab
	go.etcd.io/etcd v3.3.18+incompatible
	golang.org/x/crypto v0.0.0-20191119213627-4f8c1d86b1ba // indirect
	golang.org/x/net v0.0.0-20191119073136-fc4aabc6c914 // indirect
	golang.org/x/sys v0.0.0-20191120155948-bd437916bb0e // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/grpc v1.22.1
	gopkg.in/go-playground/validator.v8 v8.18.2
	gopkg.in/go-playground/validator.v9 v9.29.1
	gopkg.in/yaml.v2 v2.2.4
	sigs.k8s.io/yaml v1.1.0 // indirect
)

replace github.com/beanstalkd/go-beanstalk => github.com/geekymedic/go-beanstalk v0.0.0-20191210081744-8aff47a77476
