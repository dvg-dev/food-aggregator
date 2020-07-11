# Food Aggregator
REST API developed using golang, echo web framework that aggregates foods from various suppliers and provides APIs for consumers to buy foods 

There are 6 API endpoints(GET Methods) in this service
1) /food-aggregator/ping - to check if the service is up and running. It just returns a 200 OK response

2) /food-aggregator/buy-item/:name - to find the item by specifying its name in the path parameter

3) /food-aggregator/buy-item-qty/:name/:quantity - to find the item by specifying the name and quantity as path parameters

4) /food-aggregator/buy-item-qty-price/:name/:quantity/:price - to find the item by specifying the name, quantity and price as path parameters

5) /food-aggregator/show-summary - to display all the cached items

6) /food-aggregator/fast-buy-item/:name - to fetch the item by making parallel calls to Supplier APIs