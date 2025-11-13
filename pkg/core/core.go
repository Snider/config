package core

// ServiceRuntime is a helper struct embedded in services to provide access to the core application.
type ServiceRuntime[T any] struct {
	core *Core
	opts T
}

// NewServiceRuntime creates a new ServiceRuntime instance for a service.
func NewServiceRuntime[T any](c *Core, opts T) *ServiceRuntime[T] {
	return &ServiceRuntime[T]{
		core: c,
		opts: opts,
	}
}

// Core returns the central core instance.
func (r *ServiceRuntime[T]) Core() *Core {
	return r.core
}

// Config returns the registered Config service from the core application.
func (r *ServiceRuntime[T]) Config() ConfigService {
	return r.core.Config()
}

type Core struct {
	config ConfigService
}

func (c *Core) SetConfig(config ConfigService) {
	c.config = config
}

func New() (*Core, error) {
	c := &Core{}
	return c, nil
}

func (c *Core) Config() ConfigService {
	return c.config
}

type ConfigService interface {
	Save() error
	Get(key string, out any) error
	Set(key string, v any) error
	SaveStruct(key string, data interface{}) error
	LoadStruct(key string, data interface{}) error
}
