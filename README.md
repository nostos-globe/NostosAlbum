# Nostos Album Service

## Description

The Nostos Album Service manages the creation and organization of photo and video albums, allowing users to group media content from their trips. It provides privacy controls, metadata management, and optimized access to popular albums through Redis caching.

---

## Features

- Complete album creation and management  
- Association of trips with albums for organized media collections  
- Configurable privacy settings (PUBLIC/PRIVATE) for each album  
- Metadata storage and management (location, date, tags)  
- Location aggregation from associated trips  
- MinIO integration for efficient media storage  
- Redis caching for improved performance of popular albums  
- Cross-service integration with Trip and Profile services  
- User-specific album collections and public album discovery  

---

## Technologies Used

- **Language**: Go  
- **Framework**: Gin  
- **Database**: PostgreSQL with GORM  
- **Storage**: MinIO for media files  
- **Cache**: Redis for popular albums  
- **Authentication**: JWT via Auth Service  
- **Orchestration**: Docker  

---

## Architecture

The service follows a clean architecture pattern with the following components:

- **API Controllers**: Handle HTTP requests and responses  
- **Services**: Implement business logic  
- **Repositories**: Handle database operations  
- **Models**: Define data structures  
- **Clients**: Communicate with other microservices  
- **Configuration**: Manage environment settings  

---

## Database Schema

The service uses multiple schemas in PostgreSQL:

- `albums.albums`: Stores album information  
- `albums.album_trips`: Stores album-trip relationships  

---

## Album Features

### Trip Association

Albums can be associated with multiple trips, allowing users to organize their media content by themes, events, or any other criteria beyond just individual trips.

### Location Aggregation

The service aggregates location data from all trips associated with an album, providing a comprehensive view of all places featured in the album.

### Privacy Controls

Each album can have one of the following visibility settings:

- **PUBLIC**: Visible to all users  
- **PRIVATE**: Visible only to the owner  

### Cross-Service Integration

The Album Service integrates with:

- **Trip Service**: To fetch trip details and associated media  
- **Profile Service**: To manage user relationships and access controls  
- **Auth Service**: For authentication and authorization  

---

## Security

- **Authentication**: Implemented using JWT tokens from the Auth Service  
- **Access Control**: Based on user permissions and album visibility settings  
- **Media Access**: Secured through proper service integration  

---

## Structure

```plaintext
NostosAlbum/
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── api/              # HTTP controllers
│   ├── db/               # Database repositories
│   ├── models/           # Data models
│   └── service/          # Business logic and service clients
├── pkg/
│   ├── config/           # Configuration management
│   └── db/               # Database connection
├── Dockerfile            # Container definition
├── go.mod                # Go module definition
└── README.md             # This file
