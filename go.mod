module github.com/alchemicode/jServCore

go 1.21

retract (
	v1.1.0
	v1.2.0
	v1.2.1
)

require (
	github.com/pelletier/go-toml/v2 v2.1.0
	github.com/vmihailenco/msgpack/v5 v5.4.1
	gopkg.in/yaml.v3 v3.0.1
)

require github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
