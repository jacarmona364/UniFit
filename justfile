# Variables
BINARY_NAME := unifit
BIN := bin
CODE_FOLDERS := ./internal

# Construir el proyecto
build:
	@echo "Construyendo el proyecto..."
	mkdir -p {{BIN}}
	go build -o {{BIN}}/{{BINARY_NAME}} {{CODE_FOLDERS}}

# Instalar dependencias
install-deps:
	@echo "Instalando dependencias..."
	go mod tidy

# Limpiar binarios generados
clean:
	@echo "Limpiando binarios..."
	rm -f {{BIN}}/{{BINARY_NAME}}
	go clean {{CODE_FOLDERS}}

# Comprobar estilo de código / formato
check:
	@echo "Comprobando sintaxis del proyecto..."
	gofmt -e {{CODE_FOLDERS}} > /dev/null
