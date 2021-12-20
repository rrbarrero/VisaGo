# VisaGO
Prueba de concepto: Microservicio para autenticar contra un servidor ldap y que devuelve tokens JWT con los datos de usuario.

![Esquema VisaGo](visago_esquema.jpg)

## Todo
* La librería de JWT permite autorizar (jwt-middleware línea #70). Por ahora no lo voy a usar. Solo autenticar. En un futuro implementar que hable con otro microservicio para autorizaciones.


## Anotaciones:
* curl --location --request GET 'localhost:8080/auth/ping_with_token' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjg3Njc1OTksImlkIjoicm9iZXJ0by5yb2RyaWd1ZXoiLCJvcmlnX2lhdCI6MTYyODc2Mzk5OX0.Gvma5-9CjAOl_go04v_Jubq-GzEEsO7Qh2oaGANsWgQ'

* curl --location --request POST 'localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "user.name", "password": "password"
}'
