package config

import "github.com/timest/env"

// Env Env
var Env *EnvConfig

func init() {
	Env = new(EnvConfig)
	env.IgnorePrefix()
	err := env.Fill(Env)
	if err != nil {
		panic(err)
	}
}
