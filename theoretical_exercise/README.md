# Desafío Teórico

## Procesos, hilos y corrutinas

*Un caso en el que usarías procesos para resolver un problema y por qué*

Respuesta:

<span style="color:#5564eb">
Para las copias de seguridad de db.<br>
La generación de reportes (semanal. mensual, diario, etc) que la data venga de una API.<br>
Limpieza de archivos temporales, etc.<br>
Usaría procesos como el famoso CRON por su aislamiento, lo gestiona el sistema operativo.
</span>

<br>

<br>

*Un caso en el que usarías threads para resolver un problema y por qué*

Respuesta:

<span style="color:#5564eb">
Cuando un cliente realiza una compra, hay varios procesos que deben llevarse a cabo, como verificar la disponibilidad de los productos, procesar el pago y actualizar el inventario. Estos procesos pueden ejecutarse en paralelo utilizando hilos separados. Por ejemplo, un hilo puede encargarse de verificar la disponibilidad de los productos mientras otro hilo procesa el pago. Esto agiliza el proceso de compra y reduce el tiempo de espera para el usuario.
</span>

<br>

<br>

*Un caso en el que usarías corrutinas para resolver un problema y por qué*

Respuesta:

<span style="color:#5564eb">
Donde se requiere multiples solicitudes a servicios externos, por ejemplo necesito recuperar data de 3 servicios y por cada uno procesar la información. Al final, combinamos todo y se la presentamos al cliente o salvamos en db respectivamente.
</span>

## Optimización de recursos del sistema operativo
Si tuvieras 1.000.000 de elementos y tuvieras que consultar para cada uno de
ellos información en una API HTTP. ¿Cómo lo harías? Explicar.

### Respuesta:
<span style="color:#5564eb">
Hacer consultas por lotes más pequeños, es más controlable. Con ello reducimos la carga en el API y evitamos errores de timeout/sobrecarga. <br>
Implemantar un sistema de colas almacenando elementos que se consultarán para que de forma asícrona se vaya ejecutando, con ésto se aprovecha el tiempo de respuesta en las peticiones http <br>
Mecanismo de reintento <br>
Monitorear los cuellos de botella y optimización de consulta a nivel código o como consumir la API <br>
Mecanismo de cache donde sea posible
</span>

## Análisis de complejidad
1. Dados 4 algoritmos A, B, C y D que cumplen la misma funcionalidad, con
complejidades O(n2), O(n3), O(2n) y O(n log n), respectivamente, <br> ¿Cuál de los
algoritmos favorecerías y cuál descartarías en principio? Explicar por qué.

2. Asume que dispones de dos bases de datos para utilizar en diferentes
problemas a resolver. <br> La primera llamada AlfaDB tiene una complejidad de O(1)
en consulta y O(n2) en escritura. <br> La segunda llamada BetaDB que tiene una
complejidad de O(log n) tanto para consulta, como para escritura. <br> ¿Describe en
forma sucinta, qué casos de uso podrías atacar con cada una?