# Elección de la imagen del contenedor de pruebas

## Criterios:

- **Seguridad:** Minimizar la superficie de ataque. El objetivo es **0 vulnerabilidades** críticas o altas.
- **Eficiencia:** Reducir el tamaño final para agilizar los tiempos de descarga y despliegue, siempre que no comprometa los criterios anteriores.
- **Permite montar el contenedor en una etapa**

---

## Alternativas:

- **`scratch`**:
  - **Descripción:** Es una imagen vacía (sin sistema operativo).
  - **Seguridad:** **Máxima**. Superficie de ataque inexistente.
  - **Eficiencia:** **Óptima**. Tamaño mínimo absoluto.
  - **Permite montar el contenedor en una etapa:** No.


- **`golang:alpine`**:
  - **Descripción:** Distribución Linux mínima basada en Alpine.
  - **Seguridad:** Alta, aunque introduce riesgos de compatibilidad al usar musl libc en lugar de glibc que es el estándar.
  - **Eficiencia:** Muy buena (~18MB).
  - **Permite montar el contenedor en una etapa:** Si.


- **`golang:bookworm`**:
  - **Descripción:** Imagen completa basada en Debian Bookworm.
  - **Seguridad:** **Baja**. Aunque no es la imagen más segura, cumple con los mínimos requeridos para nuestro objetivo.
  - **Eficiencia:** **Baja**. Tamaño considerable (>800MB).
  - **Permite montar el contenedor en una etapa:** Si.

    
- **`debian:stable-slim`**:
  - **Descripción:** Versión recortada de Debian estable.
  - **Seguridad:** **Media**. Aunque reduce la superficie de ataque respecto a la versión completa, sigue incluyendo un sistema operativo base con utilidades estándar que suelen presentar vulnerabilidades de prioridad baja/media periódicamente.
  - **Eficiencia:** **Media**. Es más ligera que la versión completa (~75MB), pero sigue siendo mucho más pesada que las opciones estáticas (~2MB).
  - **Permite montar el contenedor en una etapa:** Si.


- **`bitnami/golang`**:
  - **Descripción:** Imagen mantenida por Bitnami, enfocada en actualizaciones rápidas de seguridad.
  - **Seguridad:** Bitnami Mantiene un perfil de vulnerabilidades bajo, pero superior al de una imagen mínima (se detectaron 12 en el análisis).
  - **Eficiencia** La imagen es **muy pesada** (>900MB) en comparación con las demás opciones.
  - **Permite montar el contenedor en una etapa:** Si. 

 
- **`gcr.io/distroless/static-debian12`**:
  - **Descripción:** Imagen mínima de Google. Contiene solo lo imprescindible para la ejecución (certificados CA y usuario nonroot) sobre una base Debian.
  - **Seguridad:** **Excelente**. Mantiene una política estricta de vulnerabilidades y se ejecuta sin privilegios.
  - **Eficiencia:** Excelente (~2MB + binario).
  - **Permite montar el contenedor en una etapa:** No.


---


## Elección de la imagen

Seleccionamos **`golang:bookworm`**.

Tras redefinir los criterios en base a la necesidad de un entorno de pruebas robusto y estándar, nos decantamos por la imagen oficial basada en Debian 12 (Bookworm). Durante el análisis evaluamos alternativas como bitnami/golang (imágenes de terceros) o alpine (imágenes mínimas), pero decidimos descartarlas. Alpine introduce complejidades de compatibilidad con librerías C (musl vs glibc) y las imágenes de terceros añaden una capa de dependencia externa innecesaria. Lo principal es buscar una imagen que ofrezca un equilibrio entre reproducibilidad y seguridad, permitiendo la gestión explícita de usuarios y permisos (useradd).

**Justificación:**

Al ser la imagen oficial basada en Debian (glibc), asegura la máxima compatibilidad con el toolchain de Go y las herramientas de sistema estándar. Esto nos permite implementar una gestión de seguridad explícita (creación de usuarios no privilegiados y gestión de cachés) utilizando comandos nativos de Linux.
