module github.com/superseriousbusiness/gotosocial

go 1.18

require (
	codeberg.org/gruf/go-errors v1.0.5
	codeberg.org/gruf/go-runners v1.2.0
	codeberg.org/gruf/go-store v1.3.6
	github.com/ReneKroon/ttlcache v1.7.0
	github.com/buckket/go-blurhash v1.1.0
	github.com/coreos/go-oidc/v3 v3.1.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/gzip v0.0.5
	github.com/gin-contrib/sessions v0.0.4
	github.com/gin-gonic/gin v1.7.7
	github.com/go-fed/httpsig v1.1.0
	github.com/go-playground/validator/v10 v10.9.0
	github.com/google/uuid v1.3.0
	github.com/gorilla/websocket v1.4.2
	github.com/h2non/filetype v1.1.3
	github.com/jackc/pgconn v1.11.0
	github.com/jackc/pgx/v4 v4.15.0
	github.com/microcosm-cc/bluemonday v1.0.16
	github.com/mitchellh/mapstructure v1.4.3
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/oklog/ulid v1.3.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/russross/blackfriday/v2 v2.1.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.10.0
	github.com/stretchr/testify v1.7.0
	github.com/superseriousbusiness/activity v1.0.1-0.20220405135100-18e8f86a760a
	github.com/superseriousbusiness/exif-terminator v0.2.0
	github.com/superseriousbusiness/oauth2/v4 v4.3.2-SSB
	github.com/tdewolff/minify/v2 v2.9.22
	github.com/uptrace/bun v1.1.3
	github.com/uptrace/bun/dialect/pgdialect v1.0.19
	github.com/uptrace/bun/dialect/sqlitedialect v1.0.19
	github.com/wagslane/go-password-validator v0.3.0
	golang.org/x/crypto v0.0.0-20211209193657-4570a0811e8b
	golang.org/x/net v0.0.0-20211209124913-491a49abca63
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	golang.org/x/text v0.3.7
	gopkg.in/mcuadros/go-syslog.v2 v2.3.0
	modernc.org/sqlite v1.14.6
	mvdan.cc/xurls/v2 v2.3.0
)

require (
	codeberg.org/gruf/go-bytes v1.0.2 // indirect
	codeberg.org/gruf/go-fastcopy v1.1.1 // indirect
	codeberg.org/gruf/go-fastpath v1.0.2 // indirect
	codeberg.org/gruf/go-format v1.0.3 // indirect
	codeberg.org/gruf/go-hashenc v1.0.1 // indirect
	codeberg.org/gruf/go-mutexes v1.1.2 // indirect
	codeberg.org/gruf/go-pools v1.0.2 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dsoprea/go-exif/v3 v3.0.0-20210625224831-a6301f85c82b // indirect
	github.com/dsoprea/go-iptc v0.0.0-20200610044640-bc9ca208b413 // indirect
	github.com/dsoprea/go-logging v0.0.0-20200710184922-b02d349568dd // indirect
	github.com/dsoprea/go-photoshop-info-format v0.0.0-20200610045659-121dd752914d // indirect
	github.com/dsoprea/go-png-image-structure/v2 v2.0.0-20210512210324-29b889a6093d // indirect
	github.com/dsoprea/go-utility/v2 v2.0.0-20200717064901-2fccff4aa15e // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-errors/errors v1.4.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-xmlfmt/xmlfmt v0.0.0-20211206191508-7fd73a941850 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/geo v0.0.0-20210211234256-740aa86cb551 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.2.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.10.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/quasoft/memstore v0.0.0-20191010062613-2bce066d2b0b // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20200410134404-eec4a21b6bb0 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/superseriousbusiness/go-jpeg-image-structure/v2 v2.0.0-20220321154430-d89a106fdabe // indirect
	github.com/tdewolff/parse/v2 v2.5.23 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/ugorji/go/codec v1.2.6 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/sys v0.0.0-20220328115105-d36c6a25d886 // indirect
	golang.org/x/tools v0.1.8 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	lukechampine.com/uint128 v1.1.1 // indirect
	modernc.org/cc/v3 v3.35.22 // indirect
	modernc.org/ccgo/v3 v3.15.13 // indirect
	modernc.org/libc v1.14.5 // indirect
	modernc.org/mathutil v1.4.1 // indirect
	modernc.org/memory v1.0.5 // indirect
	modernc.org/opt v0.1.1 // indirect
	modernc.org/strutil v1.1.1 // indirect
	modernc.org/token v1.0.0 // indirect
)
