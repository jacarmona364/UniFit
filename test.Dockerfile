# Fase de build: Instalamos Task y dependencias
FROM golang:bookworm AS builder

# Fuerzo la instalación de binarios en go/bin
ENV GOBIN=/go/bin

# Instalo el gestor de tareas
RUN go install github.com/go-task/task/v3/cmd/task@latest

# Importo el gestor de tareas para llamar a mi objetivo install, encargado de instalar las dependencias en el builder
COPY Taskfile.yml go.mod go.sum ./

# Instalo las dependencias
RUN task install

# Fase de Runner
FROM golang:bookworm

# Copiamos task, gotestsum y las librerías descargadas de go mod
COPY --from=builder /go/bin/task /usr/local/bin/task
COPY --from=builder /go/bin/gotestsum /usr/local/bin/gotestsum
COPY --from=builder /go/pkg/mod /go/pkg/mod

# Creo el ususario sin permisos
RUN useradd -m tester

# Redirijo la caché a /tmp y apunto los módulos a donde los copié
ENV GOCACHE=/tmp/gocache
ENV GOMODCACHE=/go/pkg/mod

# Doy permisos a los directorios temporales.
RUN mkdir -p /tmp/gocache && \
    chmod -R 777 /tmp/gocache && \
    chmod -R 777 /go/pkg/mod && \
    chmod +x /usr/local/bin/task && \
    chmod +x /usr/local/bin/gotestsum

# Apunto al directorio de trabajo
WORKDIR /app/test

# Cambio al usuarios in peromisos
USER tester

# El gestor de tareas ejecuta los test
ENTRYPOINT ["task", "test"] 
