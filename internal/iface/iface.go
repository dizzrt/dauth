package iface

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewExampleHandler)
