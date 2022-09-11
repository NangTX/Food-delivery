# Food-delivery
[1. Cài đặt MySQL bằng Docker](#thetieude)
   
    - docker run -d --name mysql --privileged=true -e MYSQL_ROOT_PASSWORD="ead8686ba57479778a76e"  -e MYSQL_USER="food_delivery"  -e MYSQL_PASSWORD="19e5a718a54a9fe0559dfbce6908"  -e MYSQL_DATABASE="food_delivery"  -p 3307:3306  bitnami/mysql:5.7

[2. Dowload Table Plus](#thetieude)
    
    - https://tableplus.com/
   
[3. Khởi chạy câu lệnh SQL trong Table Plus](#thetieude)
  
    - Food-delivery/restaurant.sql
    
[4. Chạy các request](#thetieude)

    - Get ping:  
        curl --location --request GET 'localhost:8080/ping'

    - Get restarant:  
        curl --location --request GET 'localhost:8080/v1/restaurants/:id'

    - Get restaurant all: 
        curl --location --request GET 'localhost:8080/v1/restaurants?limit=2'

    - Create restaurant:
        curl --location --request POST 'localhost:8080/v1/restaurants/' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "name": "test",
            "addr":"somewhere",
            "logo":{
                "url":"https://banner2.cleanpng.com/20180401/lkq/kisspng-pizza-hut-buffet-restaurant-logo-hut-5ac10f63394143.9292678315226018272345.jpg",
                "width":900,
                "height":580

            },
            "cover" : [
                {
                "url":"https://banner2.cleanpng.com/20180401/lkq/kisspng-pizza-hut-buffet-restaurant-logo-hut-5ac10f63394143.9292678315226018272345.jpg",
                "width":900,
                "height":580

            }
            ]
        }'

    - Upload image: 

        curl --location --request POST 'localhost:8080/v1/upload' \
        --form 'file=@"/C:/Users/Admin/Downloads/a.png"'
    

    - Upload restaurant by id:

        curl --location --request PATCH 'localhost:8080/v1/restaurants/5' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "name": "New restaurant"
        }'

    - Delete restaurant by id:

        curl --location --request DELETE 'localhost:8080/v1/restaurants/:id' \


            
