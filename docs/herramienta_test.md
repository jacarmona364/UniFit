# Análisis y Elección de Herramientas de Testing

## Biblioteca de Aserciones

#### Criterios: 
- Expresividad: capacidad del código para comunicar la intención del test ("quiero comprobar que esto es igual a aquello") en lugar de la implementación ("si A es distinto de B, entonces imprime error"). También incluye la calidad del mensaje cuando falla el test.
- Concisión: cantidad de código repetitivo necesario para realizar una comprobación unitaria. Poder escribir test rápido, baja fricción, para incentivar el desarrollo guiado por pruebas (TDD).
- Autovalidable: La biblioteca debe permitir que el test revele la causa del fallo sin ambigüedad y sin necesidad de depuración manual.

### El Estándar: Paquete testing

- Expresividad: Baja, la lógica no muestra la finalidad del test.
- Concesión: Baja, ya que requiere bloques if de 3-4 líneas.
- Autovalidación: Deficiente porque depende de que el desarrollador redacte manualmente un mensaje de error útil y consistente en cada fallo.

[Documentación oficial](https://pkg.go.dev/testing)

---

### Testify

- Expresividad: Nativa, funciones con nombres semánticos como ErrorIs 
- Concesión: Alta, aserciones de una línea.
- Autovalidación: Alta, genera automáticamente diffs entre el valor esperado y el obtenido, eliminando la ambigüedad.

[Documentación oficial](https://github.com/stretchr/testify)

----

### Gomega

- Expresividad: Muy alta, se lee casi como lenguaje natural.
- Concesión: Alta, aserciones de una línea.
- Autovalidación: Alta, utiliza matchers específicos.

[Documentación oficial](https://github.com/onsi/gomega)

## Otras alternativas

### QuickTest

- Expresividad: Alta, pero requiere pasar el comparador como argumento.
- Concesión: Alta, aserciones de una línea.
- Autovalidación: Alta, utiliza go-cmp internamente para diffs muy precisos.

[Documentación oficial](https://github.com/frankban/quicktest)



### Goconvey (Biblioteca de aserciones)

- Expresividad: Alta, pero requiere de un DSL anidado.
- Concesión: Baja, pues se requiere anidar bloques convey.
- Autovalidación: Alta, visualmente buena (UI web), pero introduce "funciones mágicas" como So que ocultan la lógica real de Go.

[Documentación oficial](https://github.com/smartystreets/goconvey)

----

## Biblioteca de aserciones seleccionada:

Tras el análisis comparativo, hemos seleccionado Testify como biblioteca de aserciones.

Esta decisión se justifica por las siguientes razones:

- Maximiza la Concisión: Reduce el código repetitivo de los tests considerablemente frente a la biblioteca estándar, eliminando bloques if repetitivos y facilitando la escritura rápida de pruebas ("Timely").
- Garantiza la Autovalidación: Al estandarizar los mensajes de error y generar "Diffs" automáticos, elimina la dependencia de que el desarrollador redacte manualmente mensajes de log descriptivos, asegurando que cada fallo sea autoexplicativo.
- Expresividad: A diferencia de Gomega o GoConvey, Testify utiliza funciones estándar de Go (assert.Equal) en lugar de introducir un DSL complejo, manteniendo el código legible.

---

## Test runner

Tras seleccionar Testify para las aserciones, necesitamos elegir la herramienta encargada de orquestar la ejecución, el descubrimiento de tests y el reporte de resultados.

#### Criterios: 

- Modelo de Ejecución (DSL vs. Nativo): Si el runner nos obliga a reescribir los tests usando un DSL o ejecuta funciones Go estándar
- Integración con Aserciones: Cómo se comporta con la biblioteca elegida (Testify)
- Calidad del Reporte: Capacidad para mostrar logs, panics y tiempos de ejecución.

### Gotestsum

- Modelo de Ejecución (Nativo): Ejecuta los tests estándar de Go (func TestX) sin requerir DSLs ni cambios en el código.
- Integración con Aserciones (Testify): Total. Interpreta la salida JSON de los tests y formatea los errores de Testify de manera legible.
- Calidad del Reporte: Alta. Muestra resúmenes compactos, gestiona reintentos de tests inestables y genera XML (JUnit) automáticamente.

[Documentación oficial](https://github.com/gotestyourself/gotestsum)

### Ginkgo

El ejecutor del framework Ginkgo. Aunque está pensado para BDD, puede ejecutar tests, pero con mucha sobrecarga.

- Modelo de Ejecución (DSL): Diseñado para ejecutar especificaciones BDD (Describe/It).
- Integración con Aserciones (Testify): Baja. Optimizado para Gomega. Usarlo solo para correr tests estándar es matar moscas a cañonazos.
- Calidad del Reporte: Verbosa y estructurada en árbol.

[Documentación oficial](https://github.com/onsi/ginkgo)

---

### Goconvey Runner

Runner enfocado en el desarrollo local con feedback visual.

- Modelo de Ejecución (Híbrido): Centrado en el modo "Watch" (ejecución automática al guardar).
- Integración con Aserciones (Testify): Media. Funciona, pero su valor principal está en su propia interfaz web.
- Calidad del Reporte: Visual (Semáforos en navegador). Pobre en consola.

[Documentación oficial](https://github.com/smartystreets/goconvey)

---

## Test runner seleccionado:
Hemos descartado go test por su salida primitiva y elegimos Gotestsum como herramienta de ejecución

### Justificación:

#### Herramienta Seleccionada: Gotestsum

Elegimos Gotestsum porque cumple estrictamente la función de runner sin añadir complejidad innecesaria:

- Modelo de Ejecución (Nativo): Mantiene la simplicidad al ejecutar funciones Go estándar, evitando la deuda técnica de adoptar un DSL (como Ginkgo).
- Integración con Aserciones: Complementa perfectamente a Testify, formateando sus mensajes de error de forma legible sin interferir en la lógica.
- Calidad del Reporte: Mejora sustancialmente la salida de consola (colores, agrupación de fallos) y añade capacidades profesionales (XML) que el CLI básico no tiene.

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


## Aplicación de Principios F.I.R.S.T.

Para asegurar la calidad del código en los modelos (`Rutina`, `Ejercicio`) y en el `scraper`, hemos aplicado los principios F.I.R.S.T. de la siguiente forma:

* **Fast (Rápidos):** Las pruebas son unitarias y se ejecutan en memoria. No realizamos peticiones HTTP reales a la web ni conexiones a bases de datos, lo que permite correr toda la suite en milisegundos.
* **Independent (Independientes):** Cada test prepara sus propios datos (leyendo un HTML de prueba) y no comparte estado con los demás. Si falla la validación de un ejercicio, no afecta a las pruebas de las rutinas.
* **Repeatable (Repetibles):** Eliminamos el factor aleatorio. Al usar archivos locales (`testdata`) y estructuras de datos fijas, los tests dan siempre el mismo resultado, sin depender de si la web externa está caída o ha cambiado.
* **Self-Validating (Autovalidables):** No es necesario revisar logs manualmente. El propio test compara el resultado obtenido con el esperado (por ejemplo, que un ejercicio sea "Beginner") y reporta automáticamente si hay un fallo o un error de estructura.
* **Timely (Oportunos):** Los tests se han desarrollado a la par que el código. Esto nos ha permitido definir primero los posibles errores (como estructuras HTML rotas o datos inválidos) y programar la solución para manejarlos correctamente desde el principio.
