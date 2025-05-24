# Hospital Management Module

This module provides CRUD operations for managing hospital resources including departments, beds, and patients with their hospitalization records.

## Models

### Department
Represents a hospital department with capacity management.

```json
{
  "id": "string",
  "name": "string",
  "description": "string", 
  "floor": "integer",
  "capacity": {
    "maximum_beds": "integer",
    "actual_beds": "integer", 
    "occupied_beds": "integer"
  },
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

### Bed
Represents a hospital bed assigned to a department.

```json
{
  "id": "string",
  "department_id": "string",
  "bed_type": "string",
  "bed_quality": "float64",
  "status": {
    "patient_id": "string (optional)",
    "description": "string (optional)"
  },
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

### Patient
Represents a patient with their hospitalization history.

```json
{
  "id": "string",
  "first_name": "string",
  "last_name": "string", 
  "birth_date": "string",
  "gender": "string",
  "phone": "string (optional)",
  "email": "string (optional)",
  "hospitalization_records": [
    {
      "id": "string",
      "description": "string",
      "admission_date": "datetime (optional)",
      "discharge_date": "datetime (optional)",
      "department_id": "string (optional)",
      "bed_id": "string (optional)"
    }
  ],
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

## API Endpoints

### Departments API
- `POST /api/departments` - Create a new department
- `GET /api/departments/:departmentId` - Get department details
- `GET /api/departments` - List all departments (not yet implemented)
- `PUT /api/departments/:departmentId` - Update department
- `DELETE /api/departments/:departmentId` - Delete department

### Beds API
- `POST /api/beds` - Create a new bed
- `GET /api/beds/:bedId` - Get bed details
- `GET /api/beds` - List all beds (not yet implemented)
- `GET /api/departments/:departmentId/beds` - List beds by department (not yet implemented)
- `PUT /api/beds/:bedId` - Update bed
- `DELETE /api/beds/:bedId` - Delete bed

### Patients API
- `POST /api/patients` - Create a new patient
- `GET /api/patients/:patientId` - Get patient details
- `GET /api/patients` - List all patients (not yet implemented)
- `PUT /api/patients/:patientId` - Update patient
- `DELETE /api/patients/:patientId` - Delete patient

#### Hospitalization Records Management
- `POST /api/patients/:patientId/hospitalizations` - Add hospitalization record
- `PUT /api/patients/:patientId/hospitalizations/:recordId` - Update hospitalization record
- `DELETE /api/patients/:patientId/hospitalizations/:recordId` - Delete hospitalization record

## Usage Examples

### Creating a Department
```bash
curl -X POST http://localhost:8080/api/departments \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Kardiológia",
    "description": "Oddelenie kardiológie", 
    "floor": 3,
    "capacity": {
      "maximum_beds": 20,
      "actual_beds": 18,
      "occupied_beds": 10
    }
  }'
```

### Creating a Bed
```bash
curl -X POST http://localhost:8080/api/beds \
  -H "Content-Type: application/json" \
  -d '{
    "department_id": "dept-123",
    "bed_type": "standard",
    "bed_quality": 0.8,
    "status": {
      "description": "available"
    }
  }'
```

### Creating a Patient
```bash
curl -X POST http://localhost:8080/api/patients \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Mária",
    "last_name": "Svobodová",
    "birth_date": "1985-03-15",
    "gender": "F",
    "phone": "+421902345678",
    "email": "maria.svobodova@email.sk"
  }'
```

### Adding Hospitalization Record
```bash
curl -X POST http://localhost:8080/api/patients/patient-123/hospitalizations \
  -H "Content-Type: application/json" \
  -d '{
    "description": "Hospitalizácia pre infekčnú chorobu",
    "department_id": "dept-123",
    "bed_id": "bed-456"
  }'
```

## Implementation Details

The module follows the same architectural patterns as the existing `ambulance_wl` module:

- **Interface Definition**: Each resource has a dedicated API interface (e.g., `DepartmentsAPI`)
- **Implementation**: Concrete implementations follow the `impl*API` naming pattern
- **Database Integration**: Uses the existing `db_service.DbService` interface with generics
- **Error Handling**: Consistent HTTP status codes and error responses
- **Testing**: Comprehensive test suites using testify/suite and mocks

## Testing

Run tests for the hospital management module:

```bash
cd sarsabsim-webapi
go test ./internal/hospital_mgmt/...
```

Individual test suites:
```bash
go test ./internal/hospital_mgmt/ -run TestDepartmentSuite
go test ./internal/hospital_mgmt/ -run TestBedSuite  
go test ./internal/hospital_mgmt/ -run TestPatientSuite
```

## Notes

- All models include automatic ID generation using UUID if not provided
- Timestamps (`created_at`, `updated_at`) are automatically managed
- The implementation preserves creation timestamps during updates
- List operations (GET collections) are marked as not implemented and would require additional database service methods
- The module uses BSON tags for MongoDB integration alongside JSON tags 