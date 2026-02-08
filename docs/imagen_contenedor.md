# Elección de la imagen del contenedor de pruebas

## Criterios:

- **Seguridad:** Buscamos imágenes con el menor número de vulnerabilidades conocidas.
- **Tamaño:** El tamaño es importante para reducir tiempos de descarga y despliegue pero en nuestro caso, como no queremos desplegar la aplicación, buscamos una imagen que tenga la instrumentación mínima para ejecutar los test,
  desestimando imágenes base que puedan requerir instalaciones adicionales.
- **Soporte:** La imagen debe ser oficial, con mantenimiento por parte sus los desarrolladores, eliminando aquellas opciones mantenidas por la comunidad.
- **No Multi-Stage:** Las imágenes base que obliguen a implementar una parte adicional de build al contenedor quedarán descartadas, ya que, a pesar de ser lo más preciso y que generalmente se usa con Go para compilar e instalar
  dependencias de manera aislada a la ejecución, no es necesario para nuestro caso, donde solo vamos a ejecutar los test.
- **Estabilidad:** Las imágenes con versiones de prueba, con pocas actualizaciones y baja periodicidad, que puedan generar deuda técnica a futuro, serán descartadas del mismo modo.

---
