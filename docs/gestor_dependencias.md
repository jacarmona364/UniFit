## Gestor de Dependencias

Desde Go 1.11 se introdujo el sistema de módulos (Go Modules) como forma oficial de gestionar dependencias. 
A partir de Go 1.16, el modo de módulos pasó a ser el comportamiento por defecto, quedando obsoletos los antiguos mecanismos como gopath o gestores externos como Glide y Dep. 

---

En Go no existe un gestor de dependencias externo más allá del Go Modules, el cual consta de dos ficheros: **go.mod**, que define el nombre del módulo, la versión de Go utilizada y las dependencias necesarias; y **go.sum**, 
encargado de gestionar las versiones y garantizar la integridad de las dependencias.
