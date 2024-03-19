# Manual Técnico del Proyecto

## Introducción
Este manual técnico proporciona una descripción detallada de todos los componentes utilizados en el proyecto, incluyendo la interfaz de usuario web (Web UI), la base de datos, los módulos, entre otros. El proyecto se ha desarrollado utilizando las siguientes tecnologías y herramientas:

- **Web UI**: Vite (un entorno de desarrollo para aplicaciones web modernas)
- **Base de Datos**: MySQL (un sistema de gestión de bases de datos relacional)
<!-- - **Simulador de Estados**: No se especifica en la información proporcionada. -->
- **Módulos**: Dos módulos escritos en C, uno para extraer información de la CPU y otro para la RAM.
- **Backend**: Go (también conocido como Golang, un lenguaje de programación de sistemas concurrente y compilado)
- **Docker**: Para la contenerización de los componentes del proyecto.
- **Nginx**: Se utilizó como proxy para dirigir el tráfico desde el frontend al backend.

## Componentes del Proyecto

### Web UI (Interfaz de Usuario Web)
El frontend del proyecto se ha desarrollado utilizando Vite, que es un entorno de desarrollo rápido para aplicaciones web modernas. Vite permite una experiencia de desarrollo rápida y eficiente al proporcionar características como la recarga en caliente (hot module replacement) y un rendimiento óptimo durante el desarrollo. La interfaz de usuario web proporciona la interfaz gráfica para que los usuarios interactúen con el sistema.

### Base de Datos
Para el almacenamiento y gestión de datos, se ha utilizado MySQL. MySQL es un sistema de gestión de bases de datos relacional ampliamente utilizado que proporciona una amplia gama de funcionalidades, incluyendo soporte para consultas SQL complejas, escalabilidad y seguridad.

### Módulos en C
Se han desarrollado dos módulos en el lenguaje de programación C. Uno de los módulos se encarga de extraer información del CPU, mientras que el otro módulo se encarga de obtener información de la RAM del sistema. Estos módulos proporcionan funcionalidades específicas para recopilar datos del sistema que luego se utilizan en el backend.

### Backend
El backend del proyecto se ha implementado utilizando Go (Golang), un lenguaje de programación de sistemas concurrente y compilado. Go ofrece un rendimiento excepcional y es ideal para construir aplicaciones de alta concurrencia y eficiencia. El backend se encarga de procesar las solicitudes del cliente, interactuar con la base de datos, utilizar los módulos desarrollados en C para obtener datos del sistema y proporcionar respuestas adecuadas a la interfaz de usuario web.

### Dockerización
Se ha dockerizado todo el proyecto para facilitar el despliegue y la gestión de los componentes. Se ha creado un contenedor para el backend, otro para el frontend y un tercero para la base de datos MySQL. Esto permite una fácil portabilidad y escalabilidad del sistema.

### Nginx como Proxy
Se ha utilizado Nginx como proxy para dirigir el tráfico desde el frontend al backend. Nginx proporciona funcionalidades avanzadas de enrutamiento y balanceo de carga, lo que garantiza un rendimiento óptimo y una alta disponibilidad del sistema.

## Configuración y Despliegue
Para configurar y desplegar el proyecto dockerizado, se deben seguir los siguientes pasos:

1. **Construir los Contenedores Docker**: Se deben construir los contenedores Docker para el backend, frontend y base de datos utilizando los archivos de configuración Dockerfile proporcionados.

2. **Configurar Nginx**: Se debe configurar Nginx como proxy para dirigir el tráfico desde el frontend al backend. Se deben establecer las reglas de enrutamiento y balanceo de carga según sea necesario.

3. **Desplegar los Contenedores**: Se deben desplegar los contenedores Docker en un entorno de producción utilizando herramientas como Docker Compose o Kubernetes. Se deben configurar correctamente las variables de entorno y los volúmenes según las necesidades del proyecto.

4. **Probar y Monitorizar**: Una vez desplegado, se deben realizar pruebas exhaustivas para asegurar que el sistema funciona correctamente. Se debe implementar un sistema de monitorización para supervisar el rendimiento del sistema y detectar posibles problemas.

