# Gestor de Tareas

A diferencia de otros lenguajes, como Go no incorpora un gestor de tareas nativo necesitamos recurrir a herramientas externas para automatizar pasos como:
- Compilación
- Ejecución de tests
- Formateo y lint
- Ejecutar tareas previas a despliegue o CI/CD

## Criterios para designar el gestor de tareas
1. La herramienta debe estar actualizada como máximo en los últimos 6 meses, de esta forma reduciremos la deuda técnica.
2. Esencialmente, el gestor de tareas debe ser compatible con el lenguaje elegido, como todos lo son, no hay problema en ese aspecto.
3. Rendimiento en la ejecución de tareas: aunque es una aplicación que no va a ejecutarse, a priori, más de 2 veces por semana, no es un criterio tan relevante, pero es conveniente evitar pérdidas de tiempo innecesarias en el entorno de desarrollo.
4. Facilidad de uso y mantenimiento: aunque, por lo general y dada la complejidad del proyecto, ninguno de los gestores propuestos va a preseentar dificultades muy superiores al resto, hay algunos más fáciles de usar que otros, lo cuál en términos de tiempo de desarrollo se traduce en una mejora de eficiencia tan notable como para considerarlo un criterio a tener en cuenta.
5. Cualquier detalle que ayude a reducir la deuda técnica deberá ser tomado en cuenta de cara la elección.

# Opciones de gestor de tareas

## Mage

Mage presenta un **mantenimiento** moderado, con una actividad menor que otras herramientas, lo cual puede generar algo de deuda técnica a largo plazo. En cuanto al **rendimiento**, su primera ejecución es más lenta debido a la compilación inicial, aunque las posteriores son rápidas. En **entornos dockerizados y CI/CD**, funciona correctamente siempre que se incluya la instalación del binario en la imagen base. Respecto al **toolchain**, permite ejecutar tareas directamente en Go.

**Conclusión:** En conjunto, es el más adecuado si se desea escribir tareas en Go, aunque dado el poco soporte que recibe, debemos descartarlo para no generar deuda técnica.

[Documentación oficial](https://github.com/magefile/mage)

---

## Make

Make destaca por su **mantenimiento** extremadamente estable y una comunidad muy madura. Su flexibilidad para automatizar el **toolchain** es alta, ya que puede ejecutar cualquier comando de Go sin restricciones. En términos de **rendimiento**, resulta eficiente, aunque no alcanza la velocidad de sistemas especializados como Ninja. Su uso en **Docker y CI/CD** es muy favorable, ya que suele venir preinstalado en imágenes base de Linux. Constituye una herramienta potente y versátil, aunque su sintaxis puede dificultar el mantenimiento en equipos poco familiarizados. Además, Make no es un gestor de tareas específico para Go, lo que puede hacer que la deuda técnica aumente.

**Conclusión:** muy potente y versátil, con mucho recorrido y muy estandarizada pero menos optimizada que otras, el no ser específico de Go sería el único inconveniente.

[Documentación no oficial](https://earthly.dev/blog/golang-makefile/)

---

## Task

Task ofrece un **mantenimiento** muy activo con actualizaciones frecuentes y una comunidad creciente. Su capacidad para trabajar con el **toolchain** es excelente, al ejecutar comandos de Go de manera declarativa y sin necesidad de scripts adicionales. En cuanto al **rendimiento**, es suficientemente rápido para la mayoría de entornos y pipelines modernos, aunque no está orientado a builds incrementales avanzados. En **Docker y CI/CD**, es especialmente sólido, siendo común en flujos cloud. Por ello, resulta una herramienta equilibrada, clara y fácil de mantener. No es un gestor específico de Go.

**Conclusión:** una de las opciones más equilibradas, cumple con los objetivos de rendimiento, reducción de deuda técnica y soporte muy activo de la comunidad.

[Documentación oficial](https://taskfile.dev/)

---

## Just

Just cuenta con un **mantenimiento** activo y constante. Su manejo del **toolchain** es sencillo, ya que permite ejecutar comandos externos de Go sin complicaciones. En términos de **rendimiento**, es rápido y ligero, aunque sin las optimizaciones avanzadas de herramientas específicas para compilación. Es apta en **Docker y CI/CD**, sin requerir configuraciones complejas. Es una opción adecuada para quienes buscan una sintaxis simple y directa sin recurrir a YAML.

**Conclusión:** buena opción si se prioriza una sintaxis ligera sin recurrir a YAML. Ofrece una sintaxis similar a Make, pero más simple, aunque tampoco es específico para Go. Su comunidad es muy activa.

[Documentación oficial](https://github.com/casey/just)

---

## Sage

Sage presenta un **mantenimiento** activo aunque, al ser una herramienta reciente, su comunidad aún es reducida. Ofrece un **toolchain** que permite implementar tareas en Go de forma tipada y programática. En cuanto al **rendimiento**, se comporta de manera similar a Mage: una primera ejecución más lenta por la compilación y ejecuciones posteriores rápidas. Su uso en **entornos dockerizados y CI/CD** es correcto, aunque requiere instalar su binario. Es una alternativa potente para tareas complejas escritas en Go, aunque menos probada en grandes proyectos.

**Conclusión:** potente para tareas complejas escritas en Go, pero menos probado en proyectos grandes. Lo poco madurado que aún está este gestor no obliga a descartarlo para no generar deuda técnica innecesaria.

[Documentación oficial](https://github.com/einride/sage)

---

## Ninja

Ninja destaca por su **mantenimiento** altamente activo y su adopción en proyectos de gran escala como Chromium o LLVM. Aunque su orientación principal no es el **toolchain** de Go, permite ejecutar comandos externos sin problemas. Su principal fortaleza es el **rendimiento**, siendo la herramienta más rápida de todas, especialmente en builds incrementales. En cuanto a **Docker y CI/CD**, permite su uso de manera sólida, proporcionando tiempos extremadamente bajos en compilaciones complejas. Es ideal cuando la prioridad absoluta es la velocidad, aunque menos apropiado para tareas generales como linting o testing.

**Conclusión:** es la herramienta con el mejor rendimiento de todas, especialmente en compilaciones incrementales. Es ideal para proyectos donde la prioridad absoluta es la velocidad de build, aunque su enfoque especializado lo hace menos adecuado para tareas genéricas del flujo de desarrollo.

[Documentación oficial](https://pkg.go.dev/github.com/Duncaen/go-ninja)
---

## Justificación de la elección
Aunque inicialmente la herramienta que puede parecer más interesante es Mage, el hecho de que el proyecto lleve sin actualizarse más de un año puede indicar que ha sido abandonado, incurriendo así en deuda técnica.

Por lo tanto, debido a los criterios establecidos, la herramienta que consideramos más adecuada para el proyecto es **Task**, ya que ofrece el equilibrio óptimo entre mantenimiento activo, facilidad de uso del toolchain, rendimiento suficiente para pipelines en Docker y CI/CD, y una sintaxis clara y de bajo coste de mantenimiento. A diferencia de Ninja, que está centrado en el rendimiento puro, o de Make, que puede volverse difícil de mantener, Task proporciona una solución moderna y flexible que permite cubrir todas las necesidades del proyecto sin añadir complejidad innecesaria.

---

## Resumen de las tareas:
- `task build`: Compila el proyecto y genera el binario en la carpeta bin/.
- `task install-deps`: Instala las dependencias del proyecto utilizando go mod tidy.
- `task clean`: Elimina el binario y otros archivos generados durante la compilación.
- `task check`: Ejecuta comprobaciones de calidad, incluyendo el formateo (go fmt), el análisis estático (go vet) y cualquier otra regla de sintaxis definida.
