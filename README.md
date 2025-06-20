# Golang REST API for Student Management

A robust REST API built in Golang for managing student data, featuring SQLite as the database for seamless data storage and retrieval.

## 🌟 Features

* ✅ RESTful Endpoints for Student Operations (Create, Retrieve, List)
* 🐳 Integrated SQLite for Persistent Storage
* 🛠️ Thoroughly Tested Endpoints using Postman
* 💻 Verified and Explored Database using TablePlus
* 🚀 Clean Architecture with Strong Error Handling and Input Validation

## 🛠️ Tech Stack

* **Golang**: Main programming language for the REST API
* **SQLite**: Lightweight relational database for data storage
* **Postman**: API testing and validation
* **TablePlus**: Tool for database inspection and management

## ⚡️ Getting Started

### Prerequisites

* Golang installed
* SQLite installed
* Postman installed for testing
* TablePlus (optional) for database inspection

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/student-rest-api.git
   cd student-rest-api
   ```
2. Install required dependencies:

   ```bash
   go mod tidy
   ```
3. Run the application:

   ```bash
   go run main.go
   ```

### Usage

* **Create Student**: `POST /api/students` (JSON body with `name`, `email`, `age`)
* **Get Student by ID**: `GET /api/students/{id}`
* **List All Students**: `GET /api/students`

### ⚡️ Testing

Use Postman to test the endpoints. Import the available endpoints and run requests for:

* Creating a student
* Retrieving a student by ID
* Listing all students

### 📄 Database

The database is configured using SQLite (`storage/storage.db`) and can be explored with TablePlus for seamless inspection and management.

## 🙌 Acknowledgments

A deep dive into **modern backend practices** — focusing on building REST APIs that are structured, resilient, and easy to extend. This project showcases coding craftsmanship and best practices in Golang.

## 👋 Contact

For questions or collaboration:

* Email: swamiaman9988@gmail.com
* LinkedIn: https://www.linkedin.com/in/aman-swami-6062a3289/

**Thank you for checking out this project! 🚀**
