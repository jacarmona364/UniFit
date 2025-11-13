# Gestor de Tareas

A diferencia de otros lenguajes, como Go no incorpora un gestor de tareas nativo necesitamos recurrir a herramientas externas para automatizar pasos como:
- Compilación
- Ejecución de tests
- Formateo y lint
- Ejecutar tareas previas a despliegue o CI/CD

## Criterios para designar el gestor de tareas
1. La herramienta debe estar actualizada como máximo en los últimos 6 meses, de esta forma reduciremos la deuda técnica.
2. Capacidad para automatizar comandos del toolchain (build, test, vet...), independientemente de si la herramienta usa sintaxis propia o genera llamadas externas.
3. Rendimiento en la ejecución de tareas: velocidad en ejecuciones repetidas, paralelización, capacidad de builds incrementales y tiempo reducido de resolución y ejecución de flujos complejos.
4. Portabilidad y facilidad de integración en entornos dockerizados y CI/CD.

## Opciones de gestor de tareas

**Mage**: Al estar completamente escrito en Go, permite definir tareas usando un lenguaje ya conocido, lo que facilita la integración con el proyecto sin aprender nuevas sintaxis ni lidiar con archivos YAML. 
Además, las tareas son código Go real, lo que permite aprovechar librerías existentes.
Inconvenientes: Su mantenimiento no es muy activo actualmente, lo que puede generar cierta deuda técnica si surgen errores o se requiere compatibilidad con nuevas versiones de Go. 
También carece de una comunidad tan amplia como otras herramientas, por lo que la disponibilidad de ejemplos y soporte es limitada.
[Documentación oficial](https://github.com/magefile/mage)

**Make**: Es una herramienta que permite definir tareas complejas y dependencias entre ellas, lo que la hace muy versátil y compatible con casi cualquier entorno. 
Inconvenientes: Su sintaxis basada en Makefiles puede ser compleja y difícil de mantener, especialmente para desarrolladores acostumbrados a Go. Requiere conocimientos adicionales de shell scripting, y en sistemas Windows puede ser más difícil de usar sin herramientas adicionales (como WSL). 
Además, no ofrece integración con el lenguaje Go, lo que puede ralentizar el desarrollo y generar deuda técnica si el equipo tiene que adaptar constantemente la sintaxis a sus necesidades.
[Documentación no oficial](https://earthly.dev/blog/golang-makefile/)

**Task**: Es un gestor de tareas que usa archivos YAML para definir los objetivos y comandos, lo que facilita la lectura y organización de tareas simples y no tan simples. 
Es multiplataforma y puede integrarse con pipelines de CI/CD.
Inconvenientes: El uso de YAML introduce una capa de abstracción que puede limitar la flexibilidad en proyectos grandes o con lógica más compleja. Además, no permite la integración con Go.
[Documentación oficial](https://taskfile.dev/)

**Just**: Es una herramienta de automatización que simplifica la creación de Makefiles. 
Los justfile son fáciles de leer y escribir, y permiten definir tareas y dependencias de forma clara, lo que reduce la curva de aprendizaje y mejora la mantenibilidad. Cuenta además con la ventaja de estar en constante proceso de actualización y soporte.
Inconvenientes: Aunque su sintaxis es más sencilla que la de Make, no está integrada directamente con Go, por lo que las tareas siguen ejecutando comandos externos. Esto puede limitar el aprovechamiento del tipado y la reutilización de librerías dentro de las tareas. 
Además, la comunidad es más pequeña que la de Make o Task, lo que puede afectar la disponibilidad de ejemplos y documentación avanzada.
[Documentación oficial](https://github.com/casey/just)

**Sage**: Sage está inspirado en Mage y permite definir tareas directamente en Go, reemplazando archivos tradicionales como Makefile o Justfile. Al igual que Mage, permite aprovechar tipado, autocompletado y librerías Go, lo que hace que las tareas sean sólidas y fáciles de mantener en proyectos grandes.
Inconvenientes: Es relativamente nuevo y su comunidad es pequeña, lo que puede limitar el soporte y los recursos disponibles. Al ser reciente, puede carecer de algunas funcionalidades avanzadas de otras herramientas más maduras. Su adopción en proyectos existentes puede requerir tiempo de aprendizaje adicional.
[Documentación oficial](https://github.com/einride/sage)

---

## Justificación de la elección
Se ha decidido usar Just como gestor de tareas principal del proyecto debido a su simplicidad y facilidad de mantenimiento. 
Al permitir definir tareas de forma clara y legible, reduce la posibilidad de errores humanos y facilita la automatización de tareas rutinarias como compilación, pruebas y despliegue del servidor. Además, al estar orientado a la ejecución de comandos de manera directa, no requiere dependencias externas ni cambios considerables en la estructura del proyecto.

Aunque herramientas como Mage y Sage permiten programar tareas directamente en Go, Just se ha elegido por su curva de aprendizaje más rápida y su sintaxis sencilla, especialmente útil para tareas repetitivas que no requieren lógica compleja dentro de la tarea. Esta elección contribuye a reducir la deuda técnica y facilita la escalabilidad y el mantenimiento a largo plazo del proyecto.

---

## Resumen de las tareas:
- `just build`: Compila el proyecto y genera el binario en la carpeta bin/.
- `just install-deps`: Instala las dependencias del proyecto utilizando Go Modules.
- `just clean`: Elimina el binario y otros archivos generados durante la compilación.
- `just check`: Verifica que el código fuente esté correctamente formateado y que cumpla con las reglas de sintaxis definidas.
