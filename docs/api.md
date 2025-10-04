# cookbook

## ping 
 
Get service status.

### request

```
GET /ping
```

### response

```json
{
  "message": "pong"
}
```

## search

Search recipes based on a query string and up to three optional ingredients. 

```
GET /search?q=[query-string]&ingredient=[ingredient-01],[ingredient-02],[ingredient-03]
```

> Empty query string returns all recipes.

### request

```
GET /search?q=papas%20barquito
```

```
GET /search?q=papa&ingredient=tomillo,sal,pimienta
```

### response

```json
{
  "query": "papas barquito",
  "recipes": [
    {
      "id": 6,
      "name": "papas barquito",
      "description": "nam-nam pagina 83",
      "direction": "",
      "ingredients": null
    }
    ...
  ]
}
```

## recipes/

```
GET /recipes/
```

Return all recipes.

### request

```
GET /recipes/
```

### response

```json
[
    {
        "id": 6,
        "name": "papas barquito",
        "description": "nam-nam pagina 83",
        "direction": "Antes que nada, prendemos el horno y ponemos la placa o fuente limpia para precalentarla. Mientras, lavamos y cortamos las papas<br><br>En un bol condimentamos las papas con un poquito de aceite de oliva, limon, manteca, curcuma, hierbas, sal y pimienta. Mezclamos bien<br><br>Las ponemos en la placa (muy caliente) y cocinamos en el horno a 200 °C hasta que estén doraditas y crocantes. Las podemos sacar a los 15 minutos y dar vuelta para que salgan más crocantes<br><br>Y si son muy grandes los trozos de papa, una vez en el horno, agregamos media tacita de agua para que resulten tiernas por dentro y doradas por fuera",
        "ingredients": [
          {
            "name": "papa",
            "quantity": 1,
            "unit": "unidad",
            "note": ""
          },
          {
            "name": "aceite de oliva",
            "quantity": 1,
            "unit": "a gusto",
            "note": "o manteca"
          },
          ...
    }
    ...
]
```

## recipes/:id

```
GET /recipes/:id
```

Returns a specific recipe based on and id.

### request

```
GET /recipes/6
```

### response

```json
{
  "id": 6,
  "name": "papas barquito",
  "description": "nam-nam pagina 83",
  "direction": "Antes que nada, prendemos el horno y ponemos la placa o fuente limpia para precalentarla. Mientras, lavamos y cortamos las papas<br><br>En un bol condimentamos las papas con un poquito de aceite de oliva, limon, manteca, curcuma, hierbas, sal y pimienta. Mezclamos bien<br><br>Las ponemos en la placa (muy caliente) y cocinamos en el horno a 200 °C hasta que estén doraditas y crocantes. Las podemos sacar a los 15 minutos y dar vuelta para que salgan más crocantes<br><br>Y si son muy grandes los trozos de papa, una vez en el horno, agregamos media tacita de agua para que resulten tiernas por dentro y doradas por fuera",
  "ingredients": [
    {
      "name": "papa",
      "quantity": 1,
      "unit": "unidad",
      "note": ""
    },
    {
      "name": "aceite de oliva",
      "quantity": 1,
      "unit": "a gusto",
      "note": "o manteca"
    },
    ...
}
```

### error

invalid id.

    HTTP/1.1 404 Not Found

## recipes/count

GET /recipes/count

Returns the quantity of recipes on the database.

## request

    GET /recipes/count

## response

    {
      "count": 28
    }
