# Análisis y Elección de Herramientas de Testing

## Biblioteca de Aserciones

#### Criterio: Elimina la necesidad de herramientas externas

### El Estándar: Paquete testing

Go incluye una potente librería estándar llamada testing, que aunque no es una biblioteca de aserciones como tal, permite el uso de test y errores integrados mediante el propio lenguaje Go, con elementos
de su propia sintaxis como *if* o *switch*.

Es la herramienta oficial, no requiere instalación y garantiza compatibilidad futura. No añade dependencias 
externas, mantiene el proyecto ligero.

[Documentación oficial](https://pkg.go.dev/testing)

---

### Testify

Testify es el conjunto de utilidades de testing más adoptado por la comunidad en Go fuera del estándar, complementa el runner de Go añadiendo
una capa de aserciones (assert.Equal, assert.NotNil). Requiere de una instalación extra, pues no viene integrada con Go.

[Documentación oficial](https://github.com/stretchr/testify)

---

### Gomega

Gomega es una biblioteca de aserciones que funciona mediante matchers con una sintaxis tal que (Expect(x).To(Equal(y))). Requiere la instalación de dependencias externas, no viene integrada en Go.

[Documentación oficial](https://github.com/onsi/gomega)

---

### QuickTest

Biblioteca de aserciones enfocada en proporcionar comparaciones rápidas y mensajes de error muy detallados (qt.Assert). Requiere la instalación de una dependencia externa.

[Documentación oficial](https://github.com/frankban/quicktest)

---

## Biblioteca de aserciones seleccionada:
Nos quedamos con el estándar de Go, *testing*, ya que simplifica muy considerablemente su uso e integración con el lenguaje al venir por defecto con Go.

---

## Test runner

#### Criterios: Elimina la necesidad de herramientas externas y No requiere código de inicialización (Zero Boilerplate)

### El Estándar: go test
Viene instalado por defecto de manera nativa con Go. No requiere código de setup.

[Documentación oficial](https://pkg.go.dev/testing)

### Ginkgo

Ginkgo es un framework BDD completo. Requiere instalar y un binario CLI externo además de escribir explícitamente un Hook de conexión y una función TestMain en cada paquete, añadiendo código de
inicialización que no aporta valor de negocio.

[Documentación oficial](https://github.com/onsi/ginkgo)

---

### Goconvey

Goconvey destaca por ofrecer una interfaz web en tiempo real para visualizar los tests. Requiere la instalación de binarios adicionales para levantar el servidor web de reportes. Introduce un DSL (Convey(...)) que actúa como una capa de abstracción sobre el test. 

[Documentación oficial](https://github.com/smartystreets/goconvey)

---

## Test runner seleccionado:
Nos quedamos con el estándar *go test*. Es la única opción que garantiza "Zero Config": el desarrollador solo escribe el test y funciona. Evitamos la deuda técnica de mantener archivos de configuración (boilerplate) y la gestión de binarios externos en el entorno de desarrollo. Desechamos también la necesidad de llevar a cabo instalaciones de herramientas externas

---

## Herramienta CLI de ejecución de tests

Nos quedamos con el estándar *go test*.

Esta decisión se fundamenta en la eficiencia operativa y la limpieza del código, cumpliendo estrictamente con los dos criterios de selección del test runner:

Independencia Total: Al estar integrado en el SDK, garantiza la máxima portabilidad. Cualquier entorno (local o CI/CD) capaz de compilar el proyecto es capaz de testearlo inmediatamente, sin gestionar binarios externos ni versiones incompatibles.

Mantenibilidad (Zero Boilerplate): Es la única herramienta que permite una arquitectura "Zero Config". Eliminamos la deuda técnica de mantener archivos de configuración, hooks de inicialización o estructuras de frameworks complejos, manteniendo el código fuente puro y centrado exclusivamente en la lógica de negocio.

# Elección Final:

Tras el análisis, se ha decidido desacoplar el concepto de Test Runner del de Biblioteca de Aserciones para obtener lo mejor de ambos.

1. Test Runner: go test (Librería estándar)

2. Biblioteca de Aserciones: testing

3. Herramienta CLI: Go
