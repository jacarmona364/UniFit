# Gestor de Tareas

A diferencia de otros lenguajes, como Go no incorpora un gestor de tareas nativo necesitamos recurrir a herramientas externas para automatizar pasos como:
- Compilación
- Ejecución de tests
- Formateo y lint
- Ejecutar tareas previas a despliegue o CI/CD

## Criterios para designar el gestor de tareas
1. La herramienta debe estar actualizada como máximo en los últimos 6 meses, de esta forma reduciremos la deuda técnica.
2. Rendimiento en la ejecución de tareas (Crítico para CI/CD): El rendimiento es un criterio esencial, ya que el gestor se ejecutará con frecuencia en el sistema de Integración Continua (CI). La ganancia de rendimiento, aunque sean milisegundos por tarea, se acumulará significativamente en cada pipeline, reduciendo el coste de hosting y el tiempo de feedback al desarrollador.
3. El toolchain permite ejecutar tareas directamente en Go.

# Opciones de gestor de tareas

## Mage

Mage presenta un **mantenimiento** moderado, con una actividad menor que otras herramientas, lo cual puede generar algo de deuda técnica a largo plazo pues su última actualización es de hace más de 6 meses. En cuanto al **rendimiento**, su primera ejecución es más lenta debido a la compilación inicial, tardando entre 100-500ms, aunque las posteriores son rápidas, en torno a loa 5ms. En **entornos dockerizados y CI/CD**, funciona correctamente siempre que se incluya la instalación del binario en la imagen base. Respecto al **toolchain**, permite ejecutar tareas directamente en Go.

[Documentación oficial](https://github.com/magefile/mage)

---

## Make

Make destaca por su **mantenimiento** extremadamente estable y una comunidad muy madura. El **toolchain** permite ejecutar cualquier comando de Go sin restricciones. En términos de **rendimiento**, resulta eficiente, aunque no alcanza la velocidad de sistemas especializados como Ninja, ofreciendo tiempos medios de ejecución de entre 50 y 200 ms. Su uso en **Docker y CI/CD** es muy favorable, ya que suele venir preinstalado en imágenes base de Linux.

[Documentación no oficial](https://earthly.dev/blog/golang-makefile/)

---

## Task

Task ofrece un **mantenimiento** muy activo con actualizaciones frecuentes (a día de hoy se agregan mejoras de manera diaria-semanal) y una comunidad creciente. Su **toolchain** permite ejecutar comandos de Go de manera declarativa y sin necesidad de scripts adicionales. En cuanto al **rendimiento**, es suficientemente rápido para la mayoría de entornos y pipelines modernos, el overhead es prácticamente nulo y el tiempo medio de ejecución por tarea de menos de 8 ms. En **Docker y CI/CD**, es especialmente sólido, siendo común en flujos cloud. Por ello, resulta una herramienta equilibrada, rápida y fácil de mantener.

[Documentación oficial](https://taskfile.dev/)
[Documentación GitHub](https://github.com/go-task/task)

---

## Just

Just cuenta con un **mantenimiento** activo y constante, la última actualización es de hace dos días pero suele tener mejoras con una frecuencia mensual. El **toolchain** permite ejecutar comandos de Go. En términos de **rendimiento**, es rápido y ligero, presentando un tiempo medio de ejecución por tarea inferior a los 10 ms. Es apta en **Docker y CI/CD**, sin requerir configuraciones complejas.

[Documentación oficial](https://github.com/casey/just)

---

## Sage

Sage presenta un **mantenimiento** activo aunque, al ser una herramienta reciente, su comunidad aún es reducida. Ofrece un **toolchain** que permite ejecutar tareas en Go. En cuanto al **rendimiento**, se comporta de manera similar a Mage: una primera ejecución más lenta por la compilación (100-500ms) y ejecuciones posteriores rápidas (5ms). Permite su uso en **entornos dockerizados y CI/CD**.

[Documentación oficial](https://github.com/einride/sage)

---

## Ninja

Ninja destaca por su **mantenimiento** altamente activo. Aunque su orientación principal no es el **toolchain** de Go, permite ejecutar tareas sin problemas. Su principal fortaleza es el **rendimiento**, siendo la herramienta más rápida de todas, con un tiempo medio de ejecución por tarea de 8 ms, esto se debe a dos razones: que minimiza la sobrecarga y que no hace detección de cambios compleja. En cuanto a **Docker y CI/CD**, permite su uso de manera sólida, proporcionando tiempos extremadamente bajos de compilación. Es ideal cuando la prioridad absoluta es la velocidad, aunque menos apropiado para tareas generales como linting o testing.

[Documentación oficial](https://pkg.go.dev/github.com/Duncaen/go-ninja)
---

## Justificación de la elección
Aunque inicialmente la herramienta que puede parecer más interesante es Mage, el hecho de que el proyecto lleve sin actualizarse más de 6 meses puede indicar que ha sido abandonado, incurriendo así en deuda técnica.

Por lo tanto, debido a los criterios establecidos, la herramienta que consideramos más adecuada para el proyecto es **Task**, ya que ofrece el equilibrio óptimo entre mantenimiento activo, **toolchain** que permite ejecutar tareas en Go y **rendimiento** suficiente para ejecutar las tareas pertinentes. A diferencia de Ninja, que está centrado en el rendimiento puro, o de Make, que puede volverse difícil de mantener, Task proporciona una solución moderna y flexible que permite cubrir todas las necesidades del proyecto sin añadir complejidad innecesaria.

---

## Resumen de las tareas:
- `task build`: Compila el proyecto y genera el binario en la carpeta bin/.
- `task install-deps`: Instala las dependencias del proyecto utilizando go mod tidy.
- `task clean`: Elimina el binario y otros archivos generados durante la compilación.
- `task check`: Ejecuta comprobaciones de calidad, incluyendo el formateo (go fmt), el análisis estático (go vet) y cualquier otra regla de sintaxis definida.
