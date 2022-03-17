package cache

import (
	"sync"

	"github.com/evanw/esbuild/internal/config"
)

type PluginCache struct {
	loadEntries    map[string]*config.OnLoadResult
	resolveEntries map[string]*config.OnResolveResult
	mutex          sync.Mutex
	resolveMutex   sync.Mutex
}

func (c *PluginCache) GetLoadCache(path string) *config.OnLoadResult {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.loadEntries[path]
}

func (c *PluginCache) SetLoadCache(path string, res *config.OnLoadResult) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.loadEntries[path] = res
}

func (c *PluginCache) GetResolveCache(path string) *config.OnResolveResult {
	c.resolveMutex.Lock()
	defer c.resolveMutex.Unlock()
	return c.resolveEntries[path]
}

func (c *PluginCache) SetResolveCache(path string, res *config.OnResolveResult) {
	c.resolveMutex.Lock()
	defer c.resolveMutex.Unlock()
	c.resolveEntries[path] = res
}
