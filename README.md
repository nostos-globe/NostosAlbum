# Album Service (Servicio de Álbumes e Imágenes)

## Descripción
El servicio de álbumes permite la gestión y almacenamiento de fotos y videos, configurando su privacidad y metadatos. Además, optimiza el acceso a los álbumes populares mediante Redis.

## Características
- Creación y gestión de álbumes de fotos y videos.
- Configuración de privacidad (público/privado).
- Almacenamiento de metadatos (ubicación, fecha, etiquetas).
- Integración con MinIO para el almacenamiento de imágenes y videos.
- Uso de Redis para mejorar el rendimiento en la carga de álbumes populares.

## Tecnologías Utilizadas
- **Lenguaje:** Go
- **Base de Datos:** PostgreSQL
- **Almacenamiento:** MinIO
- **Cache:** Redis
- **Orquestación:** Docker

## Instalación
1. Clona el repositorio:
   ```sh
   git clone <repo-url>
   cd album-service
   ```
2. Configura las variables de entorno en un archivo `.env`.
3. Construye y ejecuta el servicio con Docker:
   ```sh
   docker-compose up --build -d
   ```

## Endpoints

### Álbumes

| Método  | Ruta               | Descripción                          |
|---------|------------------|----------------------------------|
| POST    | `/albums`         | Crea un nuevo álbum             |
| GET     | `/albums/:id`     | Obtiene información de un álbum |
| PUT     | `/albums/:id`     | Actualiza un álbum              |
| DELETE  | `/albums/:id`     | Elimina un álbum                |
| GET     | `/albums/popular` | Obtiene los álbumes más populares |

## Seguridad
- **Autenticación:** Implementada mediante JWT.
- **Control de acceso:** Basado en permisos.
- **Optimización:** Uso de Redis para la carga rápida de álbumes populares.
