const mongoHost = process.env.AMBULANCE_API_MONGODB_HOST
const mongoPort = process.env.AMBULANCE_API_MONGODB_PORT

const mongoUser = process.env.AMBULANCE_API_MONGODB_USERNAME
const mongoPassword = process.env.AMBULANCE_API_MONGODB_PASSWORD

const database = process.env.AMBULANCE_API_MONGODB_DATABASE

const retrySeconds = parseInt(process.env.RETRY_CONNECTION_SECONDS || "5") || 5;

// try to connect to mongoDB until it is not available
let connection;
while(true) {
    try {
        connection = Mongo(`mongodb://${mongoUser}:${mongoPassword}@${mongoHost}:${mongoPort}`);
        break;
    } catch (exception) {
        print(`Cannot connect to mongoDB: ${exception}`);
        print(`Will retry after ${retrySeconds} seconds`)
        sleep(retrySeconds * 1000);
    }
}

// Create database
const db = connection.getDB(database)

// Initialize departments collection
const departmentsCollection = 'departments'
if (!db.getCollectionNames().includes(departmentsCollection)) {
    db.createCollection(departmentsCollection)
    db[departmentsCollection].createIndex({ "id": 1 }, { unique: true })
    db[departmentsCollection].createIndex({ "name": 1 })
    db[departmentsCollection].createIndex({ "floor": 1 })

    // Insert sample departments
    db[departmentsCollection].insertMany([
        {
            "id": "internal-med",
            "name": "Interné oddelenie",
            "description": "Oddelenie internej medicíny",
            "floor": 2,
            "capacity": {
                "maximum_beds": 30,
                "actual_beds": 25,
                "occupied_beds": 20
            },
            "created_at": new Date(),
            "updated_at": new Date()
        },
        {
            "id": "surgery",
            "name": "Chirurgické oddelenie",
            "description": "Oddelenie všeobecnej chirurgie",
            "floor": 3,
            "capacity": {
                "maximum_beds": 25,
                "actual_beds": 20,
                "occupied_beds": 15
            },
            "created_at": new Date(),
            "updated_at": new Date()
        },
        {
            "id": "pediatric",
            "name": "Pediatrické oddelenie",
            "description": "Oddelenie detskej medicíny",
            "floor": 4,
            "capacity": {
                "maximum_beds": 20,
                "actual_beds": 18,
                "occupied_beds": 12
            },
            "created_at": new Date(),
            "updated_at": new Date()
        }
    ])
}

// Initialize beds collection
const bedsCollection = 'beds'
if (!db.getCollectionNames().includes(bedsCollection)) {
    db.createCollection(bedsCollection)
    db[bedsCollection].createIndex({ "id": 1 }, { unique: true })
    db[bedsCollection].createIndex({ "department_id": 1 })
    db[bedsCollection].createIndex({ "bed_type": 1 })

    // Insert sample beds
    db[bedsCollection].insertMany([
        {
            "id": "int-101",
            "department_id": "internal-med",
            "bed_type": "standard",
            "bed_quality": 0.9,
            "status": {
                "patient_id": "",
                "description": "Available"
            },
            "created_at": new Date(),
            "updated_at": new Date()
        },
        {
            "id": "int-102",
            "department_id": "internal-med",
            "bed_type": "intensive",
            "bed_quality": 1.0,
            "status": {
                "patient_id": "",
                "description": "Under maintenance"
            },
            "created_at": new Date(),
            "updated_at": new Date()
        },
        {
            "id": "surg-201",
            "department_id": "surgery",
            "bed_type": "post-op",
            "bed_quality": 0.95,
            "status": {
                "patient_id": "",
                "description": "Available"
            },
            "created_at": new Date(),
            "updated_at": new Date()
        },
        {
            "id": "ped-301",
            "department_id": "pediatric",
            "bed_type": "children",
            "bed_quality": 0.85,
            "status": {
                "patient_id": "",
                "description": "Available"
            },
            "created_at": new Date(),
            "updated_at": new Date()
        }
    ])
}

// Initialize patients collection
const patientsCollection = 'patients'
if (!db.getCollectionNames().includes(patientsCollection)) {
    db.createCollection(patientsCollection)
    db[patientsCollection].createIndex({ "id": 1 }, { unique: true })
    db[patientsCollection].createIndex({ "last_name": 1 })
    db[patientsCollection].createIndex({ "birth_date": 1 })

    // Insert sample patients
    db[patientsCollection].insertMany([
        {
            "id": "pat-001",
            "first_name": "Ján",
            "last_name": "Novák",
            "birth_date": "1975-05-15",
            "gender": "M",
            "phone": "+421903123456",
            "email": "jan.novak@email.sk",
            "hospitalization_records": [
                {
                    "id": "hosp-001",
                    "description": "Hospitalizácia pre srdcové problémy",
                    "admission_date": new Date("2024-01-15T08:00:00Z"),
                    "discharge_date": new Date("2024-01-20T14:00:00Z"),
                    "department_id": "internal-med",
                    "bed_id": "int-101"
                }
            ],
            "created_at": new Date(),
            "updated_at": new Date()
        },
        {
            "id": "pat-002",
            "first_name": "Eva",
            "last_name": "Kováčová",
            "birth_date": "1988-09-23",
            "gender": "F",
            "phone": "+421905789012",
            "email": "eva.kovacova@email.sk",
            "hospitalization_records": [
                {
                    "id": "hosp-002",
                    "description": "Pooperačná starostlivosť",
                    "admission_date": new Date("2024-02-01T10:00:00Z"),
                    "department_id": "surgery",
                    "bed_id": "surg-201"
                }
            ],
            "created_at": new Date(),
            "updated_at": new Date()
        },
        {
            "id": "pat-003",
            "first_name": "Michal",
            "last_name": "Horváth",
            "birth_date": "2018-12-10",
            "gender": "M",
            "hospitalization_records": [
                {
                    "id": "hosp-003",
                    "description": "Pediatrické vyšetrenie",
                    "admission_date": new Date("2024-02-10T09:00:00Z"),
                    "department_id": "pediatric",
                    "bed_id": "ped-301"
                }
            ],
            "created_at": new Date(),
            "updated_at": new Date()
        }
    ])
}

// exit with success
print("Database initialization completed successfully")
process.exit(0);