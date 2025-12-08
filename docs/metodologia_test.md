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
