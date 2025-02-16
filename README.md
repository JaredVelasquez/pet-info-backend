# Pet Info Backend

## Descripción del Proyecto

Pet Info Backend es una aplicación backend desarrollada en Go que proporciona una API para gestionar información sobre mascotas. Utiliza la arquitectura hexagonal para una mayor separación de preocupaciones y una mejor mantenibilidad del código. La aplicación está diseñada para ejecutarse en un contenedor Docker para facilitar su despliegue y administración.

## Arquitectura Hexagonal

La arquitectura hexagonal, también conocida como Puertos y Adaptadores, es un patrón de arquitectura de software que permite una separación clara entre el núcleo de la aplicación y los detalles de implementación. Esta arquitectura facilita la prueba de la lógica de negocio y la adaptabilidad de la aplicación a diferentes interfaces de usuario y tecnologías de infraestructura.

### Estructura del Proyecto

```plaintext
.
└── pet-info-backend-docker
    ├── cmd
    │   ├── main.go
    │   └── wait-for-db.sh
    ├── pkg
    │   └── db
    │       └── database.go
    ├── infrastructure
    │   ├── handler
    │   │   ├── request
    │   │   │   ├── fields.go
    │   │   │   ├── parameter.go
    │   │   │   └── token.go
    │   │   ├── response
    │   │   │   ├── message.go
    │   │   │   └── response.go
    │   │   ├── router.go
    │   ├── user
    │   │   ├── model
    │   │   │   └── user.go
    │   │   ├── repository
    │   │   │   └── user_repository.go
    │   │   ├── service
    │   │   │   └── user_service.go
    │   │   ├── handler
    │   │   │   └── user_handler.go
    │   │   └── router.go
    │   ├── util
    │   │   └── query_util.go
    ├── go.mod
    ├── go.sum
    └── docker-compose.yml
```

## Componentes Principales

* cmd: Contiene el punto de entrada de la aplicación.
* pkg: Contiene el paquete de configuración y conexión a la base de datos.
* infrastructure: Contiene la implementación de la infraestructura, incluyendo handlers, rutas, repositorios y servicios.
* infrastructure/handler: Contiene los handlers para manejar las solicitudes HTTP.
* infrastructure/user: Contiene los modelos, repositorios y servicios relacionados con los usuarios.
* infrastructure/util: Contiene utilidades y funciones de ayuda.

## Pasos para Levantar la Aplicación con Docker

### Requisitos Previos

* Docker
* Docker Compose

## Instrucciones

### Clonar el Repositorio

```
    $ git clone https://github.com/tu-usuario/tu-repositorio.git
    $ cd tu-repositorio
```

### Configurar las Variables de Entorno

Asegúrate de que las variables de entorno necesarias para la conexión a la base de datos estén configuradas. Puedes crear un archivo .env con el siguiente contenido:

    DB_USER=tu_usuario
    DB_PASSWORD=tu_contraseña
    DB_NAME=mi_base_de_datos
    DB_HOST=db

### Construir la Imagen Docker:

- docker-compose build

### Levantar los Contenedores:

- docker-compose up

### Verificar la API

Una vez que los contenedores estén levantados, puedes verificar que la API esté funcionando correctamente utilizando curl o Postman.

- Probar la obtención de usuarios:

``` curl -X GET http://localhost:8080/users?username=johndoe
```
