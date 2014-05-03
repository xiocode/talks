$ curl -XPUT 'http://localhost:9200/twitter/tweet/_mapping' -d '{
    "tweet" : {
        "properties" : {
            "message" : {
                "type" : "string",
                "store" : true,
                "index" : "analyzed",
                "analyzer" : "mmseg"
            }
        }
    }
}'
