# Elección de la imagen del contenedor de pruebas

## Criterios:

- **Seguridad:** Minimizar la superficie de ataque. El objetivo es **0 vulnerabilidades** críticas o altas.
- **Tamaño:** Reducir el tamaño final para agilizar los tiempos de descarga y despliegue, siempre que no comprometa los criterios anteriores y disponga de elementos mínimos para poder trabajar con ella, como certificados CA, estructura de ficheros o alguna de las dependencias necesarias como glibc para soportar binarios dinámicos. 
- **Gestión de directorios:** Daremos preferencia a las imágenes que permitan la gestión de directorios, pues evitamos tener que montarlos en el builder y hacer una copia en la imagen final.
- **Gestión de usuarios:** Como necesitamos crear un usuario que ejecute los test y no tenga permisos, necesitamos una imagen que nos permita crearlo o traerlo creado.
   
Además, para este análisis descartaremos directamente imágenes que no cumplan con estos requisitos:
- **Oficial:** La imagen debe ser oficial, es decir, o bien propia del lenguaje o bien mantenidad por sus desarrolladores y no únicamente por la comunidad.
- **Estable:** Descartaremos versiones experimentales o que no hayan sido probadas y verificadas, que puedan comprometer el entorno a futuro generando deuda técnica.
- **Evitar Copy-Go-From-Builder:** Algunas de las imágenes no traen go instalado, lo cual imposibilita instalar las dependencias necesarias.

---

## Alternativas:

- **`scratch`**:
  - **Descripción:** Es una imagen vacía (sin sistema operativo).
  - **Seguridad:** **Máxima**. Superficie de ataque inexistente.
  - **Tamaño:** **Óptima**. Tamaño mínimo absoluto.
  - **Directorios:** Con builder
  - **Usuarios:** No

- **`golang:bookworm`**:
  - **Descripción:** Imagen completa basada en Debian Bookworm.
  - **Seguridad:** **Baja**. Aunque no es la imagen más segura, cumple con los mínimos requeridos para nuestro objetivo.
  - **Tamaño:** **Baja**. Tamaño considerable (>800MB), integra muchos elementos que no son necesarios para la ejecución de test, por lo que crece considerablemente.
  - **Directorios:** Sí
  - **Usuarios:** Sí


- **`bitnami/golang`**:
  - **Descripción:** Imagen mantenida por Bitnami, enfocada en actualizaciones rápidas de seguridad, es similar a bookworm.
  - **Seguridad:** Bitnami Mantiene un perfil de vulnerabilidades bajo, pero superior al de una imagen mínima (se detectaron 12 en el análisis).
  - **Tamaño** La imagen es **muy pesada** (>900MB) en comparación con las demás opciones. De nuevo es una imagen completa con utilidades innecesarias para ejecutar test.
  - **Directorios:** Sí
  - **Usuarios:** Sí
 
- **`gcr.io/distroless/static-debian12`**:
  - **Descripción:** Imagen mínima de Google. Contiene solo lo imprescindible para la ejecución sobre una base Debian.
  - **Seguridad:** **Excelente**. Mantiene una política estricta de vulnerabilidades y se ejecuta sin privilegios.
  - **Tamaño:** Excelente (~2MB + binario).
  - **Directorios:** Con builder
  - **Usuarios:** Nonroot existente en la imagen 

---

| Imagen | Tamaño |
|--------|--------|
| `scratch` | **12.4 MB** |
| `gcr.io/distroless/static-debian12` | **15.1 MB** |
| `debian:stable-slim` | 42.5 MB |
| `golang:bookworm` | > 800 MB |
| `bitnami/golang` | > 900 MB |


## Elección de la imagen

Seleccionamos **`golang:bookworm`**, ya que es la imagen que tiene todo lo que necesitamos a pesar de ser menos ligera que otras. 

**Justificación:**
Nos quedamos finalmente con bookworm, ya que es, junto con bitnami (salvo por musl), la que trae todo lo que necesitamos para poder ejecutar los test con los requerimientos para ello. Permite gestionar usuarios, directorios, trae go instalado y es más ligera que bitnami.
