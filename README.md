# cambiomz
Mozambique Unofficial Exchange API

***

#### 1. What is CambioMZ ?

A: CambioMz is a api service that return exchange values for currencies in relation with Mozambique Metica (MZN). The data is taken from the website form of the Mozambican Bank BCI in real-time.
on the request having the accurate data.

#### 2. How to use it?

A: To access the data you have to make an http request to:

For one currency : 

https://cambiomz.herokuapp.com/api/v1.0.0/exchange?query={currency(currency:"currency"){column}

Available columns:
- country
- currency
- buy
- sell
    
Available currencies:
- USD
- ZAR
- SEK
- GBP
- EUR
- NOK
- NAD
- JPY
- DKK
- CNY
- CHF
- CAD

Example: 
https://cambiomz.herokuapp.com/api/v1.0.0/exchange?query={currency(currency:"usd"){currency,buy}}

Returns :
```json
{"data":{"currency":{"buy":"64.25","currency":"USD"}}}
```

For all currencies available:

https://cambiomz.herokuapp.com/api/v1.0.0/exchange?query={currencies{column}}

Available columns:
- country
- currency
- buy
- sell

Example: 
https://cambiomz.herokuapp.com/api/v1.0.0/exchange?query={currencies{currency,buy}}

Returns :
```json
{"data":
  {"currencies":[
    {"buy":"4.26","currency":"ZAR"},
    {"buy":"64.25","currency":"USD"},
    {"buy":"6.56","currency":"SEK"},
    {"buy":"6.88","currency":"NOK"},
    {"buy":"4.26","currency":"NAD"},
    {"buy":"0.57","currency":"JPY"},
    {"buy":"83.12","currency":"GBP"},
    {"buy":"69.41","currency":"EUR"},
    {"buy":"9.29","currency":"DKK"},
    {"buy":"9.12","currency":"CNY"},
    {"buy":"65.42","currency":"CHF"},
    {"buy":"48.47","currency":"CAD"}
  ]
 }
}
```


Made with :heart: by Myself for the Community. Djamal dos Santos <djamal.dos.santos@gmail.com>
