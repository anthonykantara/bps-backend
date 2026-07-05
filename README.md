# Ghost Hive Backend

Backend system for the Ghost Hive network - a NATO-compatible defensive system against drone swarms.

## Features
- **C2 Dashboard Service**: Manages hives, threats, and missions.
- **Engagement Engine**: Intelligent planning for interception dispatch.
- **NATO Compatibility**: STANAG 4586 translation layer.
- **Hive Control**: State management, BMS, and autonomous launch.
- **Interceptor Guidance**: Jamming-resistant, terrain-following, and CV-assisted terminal guidance.
- **Security**: AES-256-GCM encryption for all communications.

## Structure
- `cmd/simulator`: Simulation runner.
- `internal/nato`: NATO STANAG interface logic.
- `internal/engagement`: Mission planning and scoring.
- `internal/security`: Encryption and data protection.
- `internal/data`: Drone specification database.

## Running the Simulation
```bash
go run cmd/simulator/main.go
```
