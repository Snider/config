package core

import (
	"testing"
)

type mockConfigService struct{}

func (m *mockConfigService) Save() error {
	return nil
}
func (m *mockConfigService) Get(key string, out any) error {
	return nil
}
func (m *mockConfigService) Set(key string, v any) error {
	return nil
}
func (m *mockConfigService) SaveStruct(key string, data interface{}) error {
	return nil
}
func (m *mockConfigService) LoadStruct(key string, data interface{}) error {
	return nil
}

func TestCore(t *testing.T) {
	t.Run("New core", func(t *testing.T) {
		c, err := New()
		if err != nil {
			t.Fatalf("New() failed: %v", err)
		}
		if c == nil {
			t.Fatalf("New() returned a nil instance")
		}
	})
}

func TestServiceRuntime(t *testing.T) {
	t.Run("NewServiceRuntime", func(t *testing.T) {
		c, _ := New()
		runtime := NewServiceRuntime(c, "test")
		if runtime.core != c {
			t.Errorf("Expected core to be the same instance")
		}
		if runtime.opts != "test" {
			t.Errorf("Expected opts to be 'test'")
		}
	})

	t.Run("Core", func(t *testing.T) {
		c, _ := New()
		runtime := NewServiceRuntime(c, "test")
		if runtime.Core() != c {
			t.Errorf("Expected Core() to return the same instance")
		}
	})

	t.Run("Config", func(t *testing.T) {
		c, _ := New()
		c.SetConfig(&mockConfigService{})
		runtime := NewServiceRuntime(c, "test")
		if runtime.Config() == nil {
			t.Errorf("Expected Config() to return a non-nil instance")
		}
	})
}
