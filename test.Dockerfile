# Fase de build: Instalamos Task y dependencias
FROM golang:bookworm AS builder

# Instalo el gestor de tareas
RUN go install github.com/go-task/task/v3/cmd/task@latest

# Fase de Runner
FROM golang:bookworm

# Copiamos task, gotestsum y las librerías descargadas de go mod
COPY --from=builder /go/bin/task /bin/task

# Creo el ususario sin permisos
RUN useradd -m tester

# Redirijo la caché a /tmp y apunto los módulos a donde los copié
ENV GOCACHE=/tmp/gocache

# Doy permisos a los directorios temporales.
RUN mkdir -p /tmp/gocache && \
    chmod -R 777 /tmp/gocache && \

# Apunto al directorio de trabajo
WORKDIR /app/test

# Cambio al usuarios in peromisos
USER tester

# El gestor de tareas ejecuta los test
ENTRYPOINT ["task", "test"] 
