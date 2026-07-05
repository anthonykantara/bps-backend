# Ghost Hive Backend Instructions

## Architecture Overview
The system follows a micro-modular architecture:
- `internal/c2`: Orchestrates the global state and engagement logic.
- `internal/engagement`: High-performance dispatching engine.
- `internal/interceptor`: Autonomous agent logic including 961 Mesh coordination and terminal guidance.
- `internal/nato`: Binary and protocol compatibility for STANAG standards.
- `internal/geo`: 3D Pathfinding and terrain-following logic.

## Key Protocols
- **961 Mesh**: Interceptor-to-interceptor communication for target deconfliction.
- **STANAG 4586**: Data Link Interface for UAV control.
- **STANAG 4607**: Ground Moving Target Indicator (GMTI) for radar fusion.

## Development Rules
1. **Concurrency**: Use Go channels and sync.RWMutex for high-frequency telemetry.
2. **Security**: All mission data must be ephemeral or encrypted with AES-256-GCM. Implement Dead Man's Switch logic in `internal/security/resilience.go`.
3. **Simulation**: Always verify changes by running `go run cmd/simulator/main.go` to ensure system latency remains within military specs (<20ms per engagement).
