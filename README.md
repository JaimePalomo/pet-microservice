# pet-microservice

- Microservicio en Go de arquitectura hexagonal para la creación y obtención de mascotas a través de su API REST.
- API documentada en el archivo swagger.json.
- Microservicio desplegado en AWS con dirección http://3.124.185.228:80

### API
- [POST] /creamascota: Endpoint para la creación de una nueva mascota.
- [GET] /lismascotas: Endpoint para obtener todas las mascotas
- [GET] /kpidemascotas: Endpoint para obtener edad media y desviación típica de una especie dada.

### Despliegue en local
- Colocarse en la carpeta raíz del directo y ejecutar el script start.sh
- Para pararlo ejecutar script stop.sh