# Nostos Album Service

The **Nostos Album Service** enables users to create and manage photo and video albums by grouping media from their trips. It provides flexible privacy settings, location metadata aggregation, and cross-service integration with user and trip services.

---

## 🚀 Features

* Full album creation and management
* Associate multiple trips with a single album
* Privacy settings per album: PUBLIC or PRIVATE
* Location and metadata management (location, date, tags)
* Aggregates location data from associated trips
* MinIO integration for efficient media storage
* Integrates with Trip, Profile, and Auth services
* User-specific and publicly discoverable album collections

---

## 📌 Endpoints

### 🔹 Album Management

* **Create Album**
  `POST /api/albums/`
  Creates a new album.

* **Get My Albums**
  `GET /api/albums/`
  Retrieves albums owned by the authenticated user.

* **Update Album**
  `PUT /api/albums/:id`
  Updates the specified album.

* **Delete Album**
  `DELETE /api/albums/:id`
  Deletes an album by ID.

* **Get Album by ID**
  `GET /api/albums/:id`
  Retrieves details for a specific album.

* **Get Albums by User ID**
  `GET /api/albums/user/:id`
  Lists albums created by a specific user.

* **Get Album Locations**
  `GET /api/albums/:id/locations`
  Returns aggregated locations from all associated trips.

* **Get Public Albums**
  `GET /api/albums/public`
  Lists all albums marked as PUBLIC.

### 🔹 Albums with Trips

* **My Albums with Trips**
  `GET /api/albums/trips/`
  Retrieves all albums (owned by the current user) with their associated trips.

* **Public Albums with Trips**
  `GET /api/albums/trips/public`
  Retrieves all public albums with associated trip data.

* **Get Album with Trips by ID**
  `GET /api/albums/trips/:id`
  Retrieves a specific album and its associated trips.

---

## ⚙️ Installation and Configuration

### Prerequisites

* Go installed
* PostgreSQL
* MinIO instance for media storage
* Docker and Docker Compose (optional, for local development)
* Auth service with JWT support

### Installation

```bash
git clone https://github.com/nostos-globe/NostosAlbums.git
cd NostosAlbums
go mod download
```

### Configuration

Ensure the following environment variables or Vault secrets are set:

* `DATABASE_URL`
* `MINIO_ENDPOINT`, `MINIO_ACCESS_KEY`, `MINIO_SECRET_KEY`
* `JWT_SECRET`

Access to Vault can be configured via token, AppRole, or Kubernetes auth.

---

## ▶️ Running the Application

```bash
go run cmd/main.go
```

---

## 🧱 Technologies Used

* **Language**: Go
* **Framework**: Gin
* **Database**: PostgreSQL (GORM)
* **Storage**: MinIO
* **Authentication**: JWT via Auth Service
* **Orchestration**: Docker

---

## 🏗️ Project Structure

```
NostosAlbums/
├── cmd/                  # Application entry point
│   └── main.go
├── internal/
│   ├── api/              # HTTP route handlers
│   ├── db/               # Database access
│   ├── models/           # Data models
│   └── service/          # Business logic
├── pkg/
│   ├── config/           # Configuration management
│   └── clients/          # External service integrations
├── Dockerfile
├── go.mod
└── README.md
```
