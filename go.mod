module github.com/Aoi-hosizora/scut-academic-notifier

go 1.18

require (
	github.com/Aoi-hosizora/ahlib v0.0.0-00010101000000-000000000000
	github.com/Aoi-hosizora/ahlib-db/xgorm v0.0.0-00010101000000-000000000000
	github.com/Aoi-hosizora/ahlib-db/xredis v0.0.0-00010101000000-000000000000
	github.com/Aoi-hosizora/ahlib-more v0.0.0-00010101000000-000000000000
	github.com/Aoi-hosizora/ahlib-web v0.0.0-00010101000000-000000000000
	github.com/Aoi-hosizora/ahlib/xgeneric v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis/v8 v8.4.11
	github.com/jinzhu/gorm v1.9.16
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/pflag v1.0.5
	gopkg.in/tucnak/telebot.v2 v2.5.0
	gopkg.in/yaml.v2 v2.3.0
)

require (
	github.com/VividCortex/mysqlerr v1.0.0 // indirect
	github.com/bytedance/sonic v1.8.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.9.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.1.1 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-sqlite3 v1.14.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.9 // indirect
	go.opentelemetry.io/otel v0.16.0 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/Aoi-hosizora/ahlib => ../ahlib
	github.com/Aoi-hosizora/ahlib-db/xgorm => ../ahlib-db/xgorm
	github.com/Aoi-hosizora/ahlib-db/xredis => ../ahlib-db/xredis
	github.com/Aoi-hosizora/ahlib-more => ../ahlib-more
	github.com/Aoi-hosizora/ahlib-web => ../ahlib-web
	github.com/Aoi-hosizora/ahlib/xgeneric => ../ahlib/xgeneric
)
