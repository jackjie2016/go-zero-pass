package config

// AuthConf is a JWT config
type AuthConf struct {
	AccessSecret string `json:",optional"`
	AccessExpire int64  `json:",optional"`
}