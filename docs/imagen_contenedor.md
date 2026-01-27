# Elección de la imagen del contenedor de pruebas

## Criterios:

- **Seguridad:** Minimizar la superficie de ataque. El objetivo es **0 vulnerabilidades** críticas o altas.
- **Operabilidad:** Capacidad de diagnóstico. Lo ideal es disponer de una shell (`/bin/sh`) para inspeccionar el sistema de archivos, verificar permisos y depurar errores.
- **Observabilidad y Mantenibilidad:** Capacidad de extender el contenedor en caliente (ej. instalar `curl` o `top` para diagnósticos de red/rendimiento) y facilidad para gestionar usuarios no-root sin configuraciones complejas.
- **Eficiencia:** Reducir el tamaño final para agilizar los tiempos de descarga y despliegue, siempre que no comprometa los criterios anteriores.

---

## Alternativas:

- **`scratch`**:
  - **Descripción:** Es una imagen vacía (sin sistema operativo).
  - **Tamaño y Seguridad:** La opción óptima teórica (mínimo tamaño absoluto, 0 vulnerabilidades).
  - **Operabilidad (Consola):** **NULA**. No dispone de shell (`/bin/sh`), lo que hace imposible acceder al contenedor para revisar logs, ficheros o permisos.
  - **Monitorización:** **NULA**. Al no tener sistema de archivos ni gestor de paquetes, es imposible instalar herramientas de diagnóstico en caliente si surge un problema.

- **`golang:alpine`**:
  - **Descripción:** Distribución de Linux orientada a la seguridad y extremada ligereza.
  - **Tamaño y Seguridad:** Muy ligera (~18MB en su versión runtime) y segura (0 vulnerabilidades). Mantenida activamente.
  - **Operabilidad (Consola):** **ALTA**. Incluye shell nativa (`/bin/sh`), permitiendo la navegación y revisión del sistema de archivos.
  - **Monitorización:** **ALTA**. Dispone del gestor de paquetes `apk`, lo que permite instalar herramientas de red o monitoreo (como `curl`, `htop`) en tiempo de ejecución sin necesidad de reconstruir la imagen.
  - Última actualización: Hace 4 días.

- **`golang:bookworm`**:
  - **Descripción:** Imagen completa basada en Debian Bookworm.
  - **Tamaño:** Excesivo para una imagen final (>800MB), ya que incluye compiladores y librerías innecesarias para la ejecución.
  - **Operabilidad:** **MUY ALTA**. Tiene todas las herramientas imaginables preinstaladas.
    
- **`debian:stable-slim`**:
  - **Descripción:** Versión recortada de Debian estable.
  - **Tamaño y Seguridad:** Significativamente más pesada que otras opciones como Alpine y suele presentar vulnerabilidades de prioridad baja/media inherentes a sus librerías base.
  - **Operabilidad (Consola):** **ALTA**. Incluye `bash`, un estándar muy familiar para desarrolladores.
  - **Monitorización:** **ALTA**. Utiliza `apt`, permitiendo una fácil instalación de herramientas.
  - Última actualización: Hace 2 días.

- **`bitnami/golang`**:
  - **Descripción:** Imagen mantenida por Bitnami, enfocada en actualizaciones rápidas de seguridad.
  - **Tamaño:** Pesada en comparación con las opciones nativas de Alpine.
  - **Operabilidad:** **ALTA**.
  - Última actualización: Hace 1 día.

---

## Métricas:

### 1. Eficiencia: Tamaño de imagen final

Medición del tamaño total de la imagen construida (Sistema operativo + Binario compilado).
*Objetivo: < 50 MB.*

| Imagen | Tamaño |
|--------|--------|
| `scratch` | **12.4 MB** |
| `golang:alpine` (Runtime) | **18.1 MB** |
| `debian:stable-slim` | 42.5 MB |
| `golang:bookworm` | > 800 MB |
| `bitnami/golang` | > 900 MB |

### 2. Rendimiento: Velocidad de ejecución de tests

Tiempo de ejecución de la suite de tests unitarios (promedio de 50 iteraciones).
*Objetivo: Minimizar latencia en CI/CD.*

| Imagen | Media | Desv.Est | Comparativa |
|--------|-------|----------|-------------|
| `scratch` | **45.10ms** | ±0.80ms | Referencia |
| `golang:alpine` | **46.50ms** | ±1.20ms | +1.4ms |
| `debian:stable-slim` | 48.20ms | ±1.15ms | +3.1ms |
| `golang:bookworm` | 55.40ms | ±3.50ms | +10.3ms |

### 3. Seguridad: Análisis de Vulnerabilidades (CVEs)

Escaneo realizado con **Trivy** (Fecha: 2026-01-27). Se prioriza la ausencia de vulnerabilidades Críticas o Altas.

| Imagen base | CRITICAL | HIGH | MEDIUM | LOW | Estado |
|-------------|----------|------|--------|-----|--------|
| `scratch` | **0** | **0** | 0 | 0 | Seguro |
| `golang:alpine` | **0** | **0** | 0 | 0 | Seguro |
| `bitnami/golang` | 0 | 0 | 4 | 8 | Requiere parches |
| `debian:stable-slim` | 0 | 0 | 2 | 16 | Requiere parches |
| `golang:bookworm` | 2 | 14 | 45 | 120 | Inseguro |

### 4. Operabilidad y Monitorización (Cualitativo)

Evaluación de la capacidad para gestionar, depurar y extender el contenedor en un entorno productivo o de desarrollo.

| Criterio | `scratch` | `golang:alpine` | `debian:stable-slim` | `golang:bookworm` |
| :--- | :---: | :---: | :---: | :---: |
| **Acceso a Consola** | **No**<br>(Caja negra) | **Sí**<br>(Nativa) | **Sí**<br>(Bash/Sh) | **Sí**<br>(Completa) |
| **Monitorización** | **No**<br>(Sin gestor paquetes) | **Sí**<br>(`apk add`) | **Sí**<br>(`apt install`) | **Sí**<br>(`apt install`) |
| **Gestión Usuarios** | **Difícil**<br>(Manual /etc/passwd) | **Nativa**<br>(`adduser`) | **Nativa**<br>(`useradd`) | **Nativa**<br>(`useradd`) |
| **Depuración** | **Imposible** | **Fácil** | **Fácil** | **Fácil** |

---


## Elección de la imagen

Para tomar esta decisión debemos atender a los criterios establecidos, pues, de momento, solo necesitamos un entorno que nos permita ejecutar una aplicación segura y mantenible con apenas dependencias. Todas las alternativas nos ofrecen estabilidad en mayor o menor medida, pero hay diferencias críticas en cuanto a operabilidad y tamaño. Las dos más afines a lo que buscamos serían scratch y alpine:3.21, decantándome finalmente por alpine:3.21 debido a que ofrece el equilibrio necesario entre ligereza extrema y capacidad de gestión.

A pesar de que scratch lidera marginalmente en métricas puras de tamaño, elijo alpine:3.21 (implementada vía Multi-Stage Build) como la opción final.

Justificación: Aunque es ligeramente más pesada que scratch (una diferencia de apenas 6 MB), sigue siendo extremadamente ligera para los estándares actuales (18.1 MB vs los 42.5 MB de Debian) y ofrece una ventaja crucial: Operabilidad. A diferencia de scratch, Alpine incluye una consola (/bin/sh) y un gestor de paquetes, permitiendo diagnosticar errores y monitorizar el contenedor en caso de fallo, algo imposible en una imagen vacía. Además, mantiene el estándar de seguridad con 0 vulnerabilidades detectadas, igualando a scratch en este aspecto crítico.
