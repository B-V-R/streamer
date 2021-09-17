# streamer

Input JSON

>{
>        "name": "BVR",
>        "address": {
>            "area": "MAMILLAGUDEM",
>            "city": "HYDERABAD",
>            "state": "TELANGANA",
>            "country": "INDIA"
>        }
>}

CURL Request

>curl --location --request POST 'http://localhost:8080/stream' \
>--header 'Content-Type: application/json' \
>--data-raw ' {
>        "name": "BVR",
>                "address": {
>                            "area": "MAMILLAGUDEM",
>                                        "city": "HYDERABAD",
>            "state": "TELANGANA",
>            "country": "INDIA"
>        }
>}'
