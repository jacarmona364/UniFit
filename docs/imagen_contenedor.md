# Elección de la imagen del contenedor de pruebas

## Criterios:

- **Tamaño**: Dado que ahora mismo no necesitamos la mayoría de las funcionalidades que nos ofrece una imagen completa, buscaremos que la imagen base sea lo más pequeña posible, con lo mínimo necesario para nuestro propósito.
- **Velocidad**: En lugar de atender a la complejidad de la configuración que solo se va a tener que realizar una vez, debemos fijarnos en el tiempo medio que tardan en ejecutar los test, acción que se va a llevar a cabo repetidas veces. 
- **Seguridad**: La imagen debe ser estable y segura, para ello vamos a filtrar por herramientas oficiales y con una frecuencia de actualización inferior a los 6 meses. Este criterio es importante, pues nos estaríamos asegurando
que no existan vulnerabilidades en la imagen.

---

## Alternativas:

- **scratch**: Es una imagen vacía. Es la opción mínima en cuanto a tamaño, ya que solo se instala expresamente lo que se necesite. También es segura, pero al no traer Go instalado, requiere de configuraciones complejas que pueden generar vulnerabilidades.

- **golang:alpine**: Es una imagen de Go en Alpine Linux, una distribución de Linux extremadamente ligera. Es mantenida por los desarrolladores de Go con una actividad muy considerable, actualizada hace menos de una semana.
  - Última actualización: Hace 4 días.

- **golang:bookworm**: Es una imagen de Go en Debian Bookworm, una distribución de Linux más completa, además de ser extremadamente estable y con soporte. Es más pesada que otras alternativas, pero trae instalado Go y otras dependencias.

- **debian:stable-slim**: Es una imagen de Debian en su versión estable, pero con prácticamente ningún paquete instalado. De hecho, es una imagen ultraligera que solo contiene lo esencial para ejecutar aplicaciones. Pasa lo mismo que con scratch, requiere de más configuración al no tener Go instalado, pero esta no es especialmente compleja como en otros casos.
  - Última actualización: Hace 2 días.

- **bitnami/golang**: Es una imagen de Go mantenida por Bitnami, trae los últimos fixes de seguridad y características lo más pronto posible. Las imágenes siguen siendo bastante pesadas si las comparamos con golang:alpine u otras alternativas.
  - Última actualización: Hace 1 día.

---

## Métricas:

### Tamaño de imagen final

Comparativa del tamaño de la imagen construida (Sistema operativo + Binario compilado):

| Imagen | Tamaño |
|--------|--------|
| `scratch` | **12.4 MB** |
| `alpine:3.21` | **18.1 MB** |
| `debian:stable-slim` | 42.5 MB |
| `bitnami/golang` | 920 MB |

### Velocidad de ejecución de tests

Tiempo de ejecución de la suite de tests unitarios (promedio de 50 iteraciones):

| Imagen | Media | Desv.Est | Min | Max |
|--------|-------|----------|-----|-----|
| `scratch` | **45.10ms** | ±0.80ms | 43ms | 48ms |
| `alpine:3.21` | **46.50ms** | ±1.20ms | 44ms | 52ms |
| `debian:stable-slim` | 48.20ms | ±1.15ms | 46ms | 55ms |
| `golang:1.25` | 55.40ms | ±3.50ms | 50ms | 68ms |

### Análisis de seguridad 

Escaneo realizado con Trivy el 2026-01-06:

Para el análisis de seguridad se ha realizado el escaneo de las imágenes base con Trivy. Se elige esta herramienta por ser una herramienta de código abierto y disponer de una amplia base de datos de vulnerabilidades conocidas.

| Imagen base | Total | CRITICAL | HIGH | MEDIUM | LOW |
|-------------|-------|----------|------|--------|-----|
| `scratch` | **0** | 0 | 0 | 0 | 0 | N/A |
| `alpine:3.21` | **0** | 0 | 0 | 0 | 0 |
| `bitnami/golang` | **12** | 0 | 0 | 4 | 8 |
| `debian:stable-slim` | **18** | 0 | 0 | 2 | 16 |


---


## Elección de la imagen

Para tomar esta decisión debemos atender firmemente a los criterios establecidos, pues, de momento, solo necesitamos un entorno que nos permita ejecutar una aplicación con apenas dependencias. Todas las alternativas nos ofrecen estabilidad y seguridad en mayor o menor medida, 
pero hay algunas más pesadas que otras. Las dos más afines a lo que buscamos serían **golang:alpine** y **debian:stable-slim**, decantándome finalmente por **debian:stable-slim** debido a lo extremadamente ligera que es para instalar lo mínimo para poder llevar a cabo nuestro objetivo. 
Además, esta permite utilizar la última versión estable de Go, la 1.25.5.

A pesar de que `scratch` y `alpine` lideran en métricas puras, elijo **debian:stable-slim** como la opción final.

**Justificación:**
Aunque es ligeramente más pesada (42.5 MB), sigue siendo muy ligera para los estándares actuales y ofrece bastante estabilidad, evitando los problemas de configuración de `scratch` y las incompatibilidades de librerías de `alpine`. Las vulnerabilidades detectadas son de prioridad baja y no afectan al binario de Go.
