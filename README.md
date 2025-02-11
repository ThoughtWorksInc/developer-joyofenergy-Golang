# Joy of Energy - Golang Edition

## Project Overview
This is a Golang implementation of the Joy of Energy project, which provides functionality for managing meter readings, calculating energy usage, and comparing price plans.

## Project Structure
- `src/`: Source code for the application
  - `config/`: Configuration files for price plans and test data
  - `meters/`: Meter-related functionality
  - `pricePlans/`: Price plan comparison logic
  - `readings/`: Reading storage and retrieval
  - `usage/`: Usage calculation services
- `tests/`: Unit tests for each module

## Prerequisites
- Go 1.21 or higher

## Running the Project
1. Clone the repository
2. Navigate to the project directory
3. Run `go mod tidy` to download dependencies
4. Run the application with `go run src/app.go`

## Running Tests
Execute all tests with:
```bash
go test ./tests/...
```

## Project Components
- **Meter Service**: Manages meter creation and retrieval
- **Reading Service**: Stores and retrieves meter readings
- **Price Plan Comparator**: Calculates costs for different price plans
- **Usage Service**: Calculates total energy usage for a meter

## Example Usage
The main application demonstrates:
- Creating meters
- Storing meter readings
- Calculating total usage
- Comparing price plans

## License
[Add your license information here]
