package metrics

import (
	"fmt"
	"sync"
)

type Registry struct {
	mu      sync.Mutex
	Counters map[string]float64
	Gauges   map[string]float64
}

var DefaultRegistry = &Registry{
	Counters: make(map[string]float64),
	Gauges:   make(map[string]float64),
}

func (r *Registry) IncrementCounter(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Counters[name]++
}

func (r *Registry) SetGauge(name string, value float64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Gauges[name] = value
}

func (r *Registry) Report() {
	r.mu.Lock()
	defer r.mu.Unlock()
	fmt.Println("--- SYSTEM METRICS ---")
	for k, v := range r.Counters {
		fmt.Printf("COUNTER %s: %.0f\n", k, v)
	}
	for k, v := range r.Gauges {
		fmt.Printf("GAUGE   %s: %.2f\n", k, v)
	}
}
