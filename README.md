
# API de Personajes de Béisbol

Una API RESTful simple para gestionar información sobre personajes de béisbol, desarrollada con **Go** y el framework **Echo**.

-----

## 🚀 Cómo Empezar

Sigue estos pasos para poner la API en funcionamiento en tu máquina local para desarrollo y pruebas.

### Prerrequisitos

Asegúrate de tener instalado Go en tu sistema. Puedes descargarlo desde [golang.org/dl](https://golang.org/dl/).

  * **Go** (versión 1.16 o superior recomendada)

### Instalación y Ejecución

1.  **Clona el repositorio** (o crea la estructura de archivos si lo estás haciendo desde cero):

    ```bash
    git clone https://github.com/tu-usuario/beisbol-api.git # Si tienes un repo
    cd beisbol-api
    ```

    Si no tienes un repositorio, asegúrate de estar en el directorio `beisbol-api` que creaste.

2.  **Inicializa el módulo Go e instala las dependencias**:

    ```bash
    go mod init beisbol-api # Solo si no lo has hecho ya
    go mod tidy
    ```

    Esto descargará el framework Echo y otras dependencias necesarias.

3.  **Ejecuta la aplicación**:

    ```bash
    go run main.go
    ```

    La API estará disponible en `http://localhost:8080`.

-----

## 🛠️ Endpoints de la API

La API expone los siguientes endpoints para gestionar los personajes de béisbol. Todos los endpoints de `POST` y `PUT` esperan un cuerpo de solicitud JSON con el tipo de contenido `application/json`.

### Estructura de un Personaje (`Character`)

Los personajes se representan con la siguiente estructura JSON:

```json
{
  "id": 1,          // int (Solo lectura, asignado por la API)
  "name": "Fernando Tatis Jr.", // string
  "team": "Padres", // string
  "position": "Right Fielder", // string
  "battingAvg": 0.280 // float64
}
```

-----

### ⚾ `GET /characters`

  * **Descripción**: Recupera una lista de todos los personajes de béisbol.
  * **Método**: `GET`
  * **URL**: `http://localhost:8080/characters`
  * **Parámetros**: Ninguno
  * **Respuestas Posibles**:
      * `200 OK`: Una lista JSON de objetos `Character`.
        ```json
        [
          {
            "id": 1,
            "name": "Fernando Tatis Jr.",
            "team": "Padres",
            "position": "Right Fielder",
            "battingAvg": 0.28
          },
          {
            "id": 2,
            "name": "Juan Soto",
            "team": "Yankees",
            "position": "Left Fielder",
            "battingAvg": 0.31
          }
        ]
        ```

-----

### ⚾ `GET /characters/:id`

  * **Descripción**: Recupera los detalles de un personaje específico por su ID.
  * **Método**: `GET`
  * **URL**: `http://localhost:8080/characters/{id}`
  * **Parámetros de URL**:
      * `id` (requerido): El ID numérico del personaje.
  * **Respuestas Posibles**:
      * `200 OK`: Un objeto `Character` JSON.
        ```json
        {
          "id": 1,
          "name": "Fernando Tatis Jr.",
          "team": "Dodgers",
          "position": "Right Fielder",
          "battingAvg": 0.28
        }
        ```
      * `400 Bad Request`: Si el `id` proporcionado no es un número válido.
      * `404 Not Found`: Si no se encuentra un personaje con el `id` dado.

-----

### ⚾ `POST /characters`

  * **Descripción**: Crea un nuevo personaje de béisbol. El ID será asignado automáticamente por la API.
  * **Método**: `POST`
  * **URL**: `http://localhost:8080/characters`
  * **Cuerpo de la Solicitud (JSON)**:
    ```json
    {
      "name": "Mookie Betts",
      "team": "Dodgers",
      "position": "Right Fielder",
      "battingAvg": 0.307
    }
    ```
      * `id` no debe ser incluido en la solicitud.
  * **Respuestas Posibles**:
      * `201 Created`: El objeto `Character` recién creado, incluyendo su `id` asignado.
        ```json
        {
          "id": 3,
          "name": "Mookie Betts",
          "team": "Dodgers",
          "position": "Right Fielder",
          "battingAvg": 0.307
        }
        ```
      * `400 Bad Request`: Si el cuerpo de la solicitud es un JSON inválido o faltan campos requeridos (ej. `name`).

-----

### ⚾ `PUT /characters/:id`

  * **Descripción**: Actualiza completamente un personaje existente por su ID. Debes proporcionar todos los campos del personaje, incluso si no cambian.
  * **Método**: `PUT`
  * **URL**: `http://localhost:8080/characters/{id}`
  * **Parámetros de URL**:
      * `id` (requerido): El ID numérico del personaje a actualizar.
  * **Cuerpo de la Solicitud (JSON)**:
    ```json
    {
      "name": "Fernando Tatis Jr.",
      "team": "San Diego Padres",
      "position": "Right Fielder",
      "battingAvg": 0.285
    }
    ```
      * Si un campo se omite o se envía con un valor por defecto (ej. `0` para `battingAvg` si no se quiere cambiar), no se actualizará a menos que sea un valor intencional. Para una actualización parcial (PATCH), se necesitaría otro endpoint o lógica.
  * **Respuestas Posibles**:
      * `200 OK`: El objeto `Character` actualizado.
        ```json
        {
          "id": 1,
          "name": "Fernando Tatis Jr.",
          "team": "San Diego Padres",
          "position": "Right Fielder",
          "battingAvg": 0.285
        }
        ```
      * `400 Bad Request`: Si el `id` es inválido o el cuerpo de la solicitud es un JSON inválido.
      * `404 Not Found`: Si no se encuentra un personaje con el `id` dado.

-----

### ⚾ `DELETE /characters/:id`

  * **Descripción**: Elimina un personaje de béisbol por su ID.
  * **Método**: `DELETE`
  * **URL**: `http://localhost:8080/characters/{id}`
  * **Parámetros de URL**:
      * `id` (requerido): El ID numérico del personaje a eliminar.
  * **Respuestas Posibles**:
      * `204 No Content`: El personaje fue eliminado exitosamente.
      * `400 Bad Request`: Si el `id` proporcionado no es un número válido.
      * `404 Not Found`: Si no se encuentra un personaje con el `id` dado.

-----

## 📚 Consideraciones para Producción

Este proyecto utiliza una base de datos en memoria (`map`). Para una aplicación de producción, deberías integrar una base de datos persistente como PostgreSQL, MySQL, MongoDB, u otra.

-----

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Siéntete libre de abrir un *issue* o enviar un *pull request*.

-----

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo `LICENSE` para más detalles.

-----

