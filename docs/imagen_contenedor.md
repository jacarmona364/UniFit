# Elección de la imagen del contenedor de pruebas

## Criterios:

- **Seguridad:** Minimizar la superficie de ataque. El objetivo es **0 vulnerabilidades** críticas o altas.
- **Eficiencia:** Reducir el tamaño final para agilizar los tiempos de descarga y despliegue, siempre que no comprometa los criterios anteriores.
- **Inmutabilidad** El entorno de pruebas debe ser estéril e inmutable. No debe permitir la modificación del sistema en tiempo de ejecución, ya sea para instalación de paquetes manuales o cambios de configuración en caliente.

---

## Alternativas:

- **`scratch`**:
  - **Descripción:** Es una imagen vacía (sin sistema operativo).
  - **Seguridad:** **Máxima**. Superficie de ataque inexistente.
  - **Inmutabilidad:** **Absoluta**. Al no tener shell ni librerías, es imposible modificar el entorno manualmente; si funciona, es únicamente por el binario.
  - **Eficiencia:** **Óptima**. Tamaño mínimo absoluto.

- **`golang:alpine`**:
  - **Descripción:** Distribución Linux mínima basada en Alpine.
  - **Seguridad:** Alta, aunque introduce riesgos de compatibilidad al usar musl libc en lugar de glibc que es el estándar.
  - **Inmutabilidad:** **Baja (Mutable)**. Incluye gestor de paquetes apk y shell /bin/sh. Esto permite la instalación de dependencias no declaradas y la modificación del entorno en tiempo de ejecución, lo cual es un riesgo para la fiabilidad de los tests.
  - **Eficiencia:** Muy buena (~18MB).

- **`golang:bookworm`**:
  - **Descripción:** Imagen completa basada en Debian Bookworm.
  - **Seguridad:** **Baja**. Superficie de ataque masiva debido a cientos de librerías y binarios preinstalados que no necesitamos.
  - **Inmutabilidad:** **Nula**. Entorno totalmente mutable con herramientas de compilación y gestión de paquetes completas como apt, permitiendo cualquier modificación del sistema.
  - **Eficiencia:** **Pésima**. Tamaño excesivo (>800MB) que ralentiza drásticamente los tiempos de descarga en CI/CD.
    
- **`debian:stable-slim`**:
  - **Descripción:** Versión recortada de Debian estable.
  - **Seguridad:** **Media**. Aunque reduce la superficie de ataque respecto a la versión completa, sigue incluyendo un sistema operativo base con utilidades estándar que suelen presentar vulnerabilidades de prioridad baja/media periódicamente.
  - **Inmutabilidad:** **Baja (Mutable)**. Mantiene el gestor de paquetes apt y una shell completa (bash), lo que permite instalar dependencias no declaradas y alterar el contenedor, violando el criterio de inmutabilidad.
  - **Eficiencia:** **Media**. Es más ligera que la versión completa (~75MB), pero sigue siendo mucho más pesada que las opciones estáticas (~2MB).

- **`bitnami/golang`**:
  - **Descripción:** Imagen mantenida por Bitnami, enfocada en actualizaciones rápidas de seguridad.
  - **Seguridad:** Bitnami Mantiene un perfil de vulnerabilidades bajo, pero superior al de una imagen mínima (se detectaron 12 en el análisis).
  - **Inmutabilidad:** **Nula**. Entorno totalmente mutable con herramientas de compilación y gestión de paquetes completas como apt, permitiendo cualquier modificación del sistema.
  - **Eficiencia** La imagen es **muy pesada** (>900MB) en comparación con las demás opciones.
 
- **`gcr.io/distroless/static-debian12`**:
  - **Descripción:** Imagen mínima de Google. Contiene solo lo imprescindible para la ejecución (certificados CA y usuario nonroot) sobre una base Debian.
  - **Seguridad:** **Excelente**. Mantiene una política estricta de vulnerabilidades y se ejecuta sin privilegios.
  - **Inmutabilidad:** **Muy Alta**. Carece de shell y gestor de paquetes (apt eliminado), impidiendo modificaciones manuales o instalaciones en caliente dentro del entorno de pruebas.
  - **Eficiencia:** Excelente (~2MB + binario).

---

## Métricas:

### 1. Eficiencia: Tamaño de imagen final

Medición del tamaño total de la imagen construida (Sistema operativo + Binario compilado).
*Objetivo: < 50 MB.*

| Imagen | Tamaño |
|--------|--------|
| `scratch` | **12.4 MB** |
| `gcr.io/distroless` | **16.5 MB** |
| `golang:alpine` (Runtime) | **18.1 MB** |
| `debian:stable-slim` | 42.5 MB |
| `golang:bookworm` | > 800 MB |
| `bitnami/golang` | > 900 MB |

### 2. Seguridad: Análisis de Vulnerabilidades (CVEs)

Escaneo realizado con **Trivy** (Fecha: 2026-01-27). Se prioriza la ausencia de vulnerabilidades Críticas o Altas.

| Imagen base | CRITICAL | HIGH | MEDIUM | LOW | Estado |
|-------------|----------|------|--------|-----|--------|
| `scratch` | **0** | **0** | 0 | 0 | Seguro |
| `gcr.io/distroless` | **0** | **0** | 0 | 0 | Seguro |
| `golang:alpine` | **0** | **0** | 0 | 0 | Seguro |
| `bitnami/golang` | 0 | 0 | 4 | 8 | Requiere parches |
| `debian:stable-slim` | 0 | 0 | 2 | 16 | Requiere parches |
| `golang:bookworm` | 2 | 14 | 45 | 120 | Inseguro |

### 3. Inmutabilidad: Superficie de Modificación

Análisis de la capacidad de alteración del entorno en tiempo de ejecución. Se busca un entorno estéril.

| Imagen Base | Shell | Gestor Paquetes | Nivel de Inmutabilidad |
|-------------|-----------------------------|---------------------------------------|------------------------|
| `scratch` | No | No | Absoluta |
| `gcr.io/distroless/static-debian12`| No | No | Muy Alta |
| `golang:alpine` | sh | apk | Baja (Mutable) |
| `debian:stable-slim` | bash | apt | Baja (Mutable) |
| `bitnami/golang` | bash | apt | Nula |
| `golang:bookworm` | bash | apt | Nula |


## Elección de la imagen

Seleccionamos **`gcr.io/distroless/static-debian12`**.

Tras redefinir los criterios en base a los requerimientos mínimos de la imagen, nos volvemos a decantar por gcr.io/distroless/static-debian12. Durante la búsqueda de documentación hemos encontrado nuevas alternativas como cgr.dev/chainguard/static que decidimos descartar por ser muy similares a los ejemplos que ya hemos estudiado y no aportarían nada nuevo. Lo principal es buscar una imagen que tenga lo estrictamente necesario parea ejecutar los test, que no conceda permisos de usuario innecesarios y reduzca el plano de vulnerabilidades eliminando toda dependencia que no sea estrictamente necesaria. Esto incluye terminal o gestores de paquetes, que pueden derivar en actuaciones que hagan pasar los test debido a modificaciones de la imagen y no estrictamente del código fuente, haciéndola mutable.

**Justificación:**

Elegimos **Distroless** por ser el estándar de facto para contenedores *distroless* y por basarse en **Debian (glibc)**. Esto garantiza la máxima compatibilidad con los binarios compilados en Go sin los problemas de enlazado que a veces surgen con otras librerías C, cumpliendo con todos los requisitos de seguridad e inmutabilidad establecidos.
