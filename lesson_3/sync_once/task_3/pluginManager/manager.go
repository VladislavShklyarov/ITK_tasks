package pluginManager

import (
	"fmt"
	"sync"
)

type Plugin interface {
	Execute() string
}

type PluginManager struct {
	plugins map[string]*pluginEntry
	mu      sync.RWMutex
}

type pluginEntry struct {
	once     sync.Once // Добавили поле once для однократной инициализации
	initFn   func() (Plugin, error)
	err      error  // для ошибки
	instance Plugin // сам плагин
}

func NewPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make(map[string]*pluginEntry),
	}
}

func (pm *PluginManager) RegisterPlugin(name string, initFn func() (Plugin, error)) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.plugins[name] = &pluginEntry{
		initFn: initFn,
	}
}

func (pm *PluginManager) GetPlugin(name string) (Plugin, error) {

	// 1. Проверка существования плагина
	pm.mu.RLock()
	if plugin, ok := pm.plugins[name]; !ok {
		return nil, fmt.Errorf("not implemented")
	} else {
		plugin.once.Do(
			// 2. Потокобезопасная однократняа инициализацию + 	обработка и кэширование ошибок
			func() {
				plugin.instance, plugin.err = plugin.initFn()
			})
		pm.mu.RUnlock()
		return plugin.instance, plugin.err
	}
}

type DemoPlugin struct{}

func (p *DemoPlugin) Execute() string {
	return "DemoPlugin executed successfully!"
}
