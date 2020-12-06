# Seminario2020GoLang
API REST GoLang - SQLite 3
## Endpoints de los recursos
### Houses:
```
GET     http://localhost:8080/houses       Obtener todos
GET     http://localhost:8080/houses/:ID   Obtener todos por ID
POST    http://localhost:8080/houses       Insertar uno
PUT     http://localhost:8080/houses/:ID   Actualiza el campo "status" a "Sold"
DELETE  http://localhost:8080/houses/:ID   Eliminar por ID
```

## Estructuras JSON de los recursos
### Houses:
``` 
{
  "Id": 74,
  "Name": "Las Acacias",
  "Status": "For Sale",
  "Rooms": 1,
  "Price": 70.5
}
```

## Tutorial de Instalación
_Estas instrucciones te permitirán obtener una copia del proyecto en funcionamiento en tu máquina local para propósitos de desarrollo y pruebas._

### 1. Instalar GoLang
* [Descargar GoLang](https://golang.org/dl)

### 2. Clonar el repositorio actual
```
git clone https://github.com/martinmassimo/Seminario2020GoLang.git
```

### 3. Comandos Go para ejecutar el proyecto e inicializar la BD de SQLite
```
go run ./cmd/housesForSale/housesForSale.go -config ./config/config.yaml
```

### * Opcionalmente puede descargar Postman para realizar pruebas de funcionamiento

* [Postman](https://www.postman.com/downloads/)

## Autor

* **Martín Massimo** - *Desarrollo y documentación* - [martinmassimo](https://github.com/martinmassimo)
