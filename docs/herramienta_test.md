# Análisis y Elección de Herramientas de Testing

Para garantizar la calidad del código en UniFit, se ha realizado un análisis de las herramientas disponibles en el ecosistema de Go, evaluando criterios 
como el mantenimiento, la deuda técnica generada y la adopción por la comunidad.

### El Estándar: Paquete testing

Go incluye una potente librería estándar llamada testing.

Ventajas: Es la herramienta oficial, no requiere instalación, es extremadamente rápida y garantiza compatibilidad futura. Al no añadir dependencias 
externas, mantiene el proyecto ligero. Es muy utilizada por la comunidad y tiene actualizaciones muy frecuentes (en torno a unas 2 al mes).

Limitaciones: Como biblioteca de aserciones es muy primitiva. Obliga a escribir mucha lógica condicional repetitiva (if got != expected { t.Error(...) }), 
lo que puede hacer que los tests sean difíciles de mantener si crecen mucho.

[Documentación oficial](https://pkg.go.dev/testing)

---

### Análisis de Alternativas

1. Testify

Testify es el conjunto de utilidades de testing más adoptado por la comunidad en Go fuera del estándar, complementa el runner de Go añadiendo
una capa de aserciones (assert.Equal, assert.NotNil). Permite escribir tests mucho más declarativos y legibles sin introducir un framework pesado, lo que lo hace
más rápido. Tiene actualizaciones más o menos cada trimestre.

[Documentación oficial](https://github.com/stretchr/testify)

---

2. Ginkgo

Ginkgo es un framework BDD (Behavior-Driven Development) completo, utiliza su propio Lenguaje Específico de Dominio (DSL) con bloques Describe e It. Aunque 
es muy potente, introduce una capa de complejidad y abstracción muy alta, lo que lo hace más difícil de mantener. Para un proyecto como UniFit, aprender su 
DSL y depender de su estructura añade una carga cognitiva y técnica innecesaria que pueden derivar en deuda técnica futura. Tiene actualizaciones irregulares pero 
de una media mensual y la comunidad lo tiene considerablemente adoptado, aunque no es el favorito.

[Documentación oficial](https://github.com/onsi/ginkgo)

---

3. Goconvey

Goconvey destaca por ofrecer una interfaz web en tiempo real para visualizar los tests. Aunque visualmente es atractiva, su repositorio muestra una 
actividad reducida recientemente (pocas actualizaciones en los últimos meses), lo que plantea dudas sobre su mantenimiento a largo plazo y la posible deuda técnica.

[Documentación oficial](https://github.com/smartystreets/goconvey)

---

# Elección Final: Estrategia Híbrida

Tras el análisis, se ha decidido desacoplar el concepto de Test Runner del de Biblioteca de Aserciones para obtener lo mejor de ambos.

1. Test Runner: go test (Librería estándar) Se utilizará el comando nativo de Go como ejecutor de pruebas.
Justificación: Es el estándar de la industria, garantiza ejecución paralela eficiente y evita depender de binarios externos para compilar y lanzar los tests.

2. Biblioteca de Aserciones: testify Se importará el paquete github.com/stretchr/testify/assert.
Justificación: Se elige para cumplir con los principios de Código Limpio. Mientras que go test gestiona la ejecución, testify nos permite validar la lógica de negocio
con un lenguaje natural y claro, facilitando la lectura de los tests y la identificación rápida de errores.

Conclusión: Esta combinación (go test + testify) maximiza la mantenibilidad y rapidez sin introducir la complejidad de frameworks pesados como Ginkgo, eliminando así deuda técnica futura.
