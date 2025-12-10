# Metodología de Test (TDD)

Con el objetivo de garantizar la fiabilidad y robustez del software, en este proyecto se ha adoptado el enfoque de Desarrollo Guiado por Test (TDD, 
Test-Driven Development). Esta metodología invierte el flujo de trabajo tradicional, estableciendo como norma la creación de las pruebas unitarias antes que la implementación del
código funcional. De este modo, se garantiza que cada línea de código escrita se refiere directamente a un requisito verificable.

### El ciclo de trabajo adoptado para aplicar TDD consta de los siguientes pasos:

1. Selección de la funcionalidad: Se identifica una historia de usuario o especificación concreta que se desea desarrollar.
2. Creación del test (Fase Roja): Se programa un test que falle inicialmente. Esto confirma que la funcionalidad deseada aún no existe o no funciona, validando la utilidad de la prueba.
3. Implementación mínima (Fase Verde): Se escribe el código estrictamente necesario para que el test pase exitosamente. En esta fase se prioriza la funcionalidad sobre la perfección del código.
4. Verificación de integridad: Se ejecuta la suite completa de tests para asegurar que el nuevo código no ha introducido errores en funcionalidades previas (regresión).
5. Refactorización: Se optimiza el código escrito para mejorar su legibilidad, estructura y eficiencia, sin alterar su comportamiento externo.
6. Reiteración: Se marca el requisito como completado y se reinicia el ciclo seleccionando la siguiente funcionalidad de la lista de pendientes.

La aplicación de este flujo asegura que el producto final cumple con las especificaciones de diseño desde su origen. Además, 
fomenta una arquitectura de software más limpia (Clean Code) y disminuye drásticamente la introducción de errores durante la etapa de desarrollo.

# Principios de Diseño de Test (F.I.R.S.T.)

Para garantizar que la batería de pruebas aporte valor real y no se convierta en un obstáculo para el desarrollo, este proyecto se adhiere estrictamente a los principios F.I.R.S.T. Este acrónimo define las cinco características esenciales que debe poseer un test unitario de calidad:

1. Fast (Velocidad): La ejecución de los tests debe ser inmediata. La rapidez es crucial para no romper el flujo de trabajo del desarrollador, 
permitiendo que se lancen constantemente tras cada pequeña modificación sin tiempos de espera.
2. Independent (Independencia): Cada prueba debe ser una unidad aislada. No debe existir acoplamiento entre tests; el resultado de uno no puede depender 
de la ejecución previa de otro, permitiendo que se ejecuten en cualquier orden o de forma individual.
3. Repeatable (Repetibilidad): Los resultados deben ser deterministas. Un test debe arrojar siempre el mismo resultado (éxito o fallo) independientemente 
del entorno, el sistema operativo o el momento en que se ejecute, eliminando la incertidumbre o "flakiness".
4. Self-validating (Autovalidación): El test debe emitir un veredicto binario y automático (Pasa/Falla). No debe ser necesaria la intervención humana 
para interpretar logs o comparar salidas visualmente para saber si la prueba ha sido exitosa.
5. Timely (Oportunidad): Las pruebas deben escribirse en el momento adecuado, que según la metodología TDD es antes de escribir el código de producción. 
Esto asegura que el diseño del software sea testeable desde su concepción.

La aplicación de estos principios nos permite mantener una suite de pruebas fiable y escalable, asegurando la integridad funcional del código a medida que el proyecto crece.
