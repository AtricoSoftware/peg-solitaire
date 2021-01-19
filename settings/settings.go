// Generated 2021-01-19 11:42:11 by go-framework v1.2.1-1-ge87b524
package settings

import "github.com/atrico-go/container"

type Settings interface {
	// Output all stages of operation
	Verbose() bool
}

// Register the settings
func RegisterSettings(c container.Container) {
	c.Singleton(func() Settings {return theSettings{}})
}

// Stub object for settings interface
type theSettings struct{}
