### Query Tasks - GraphQL

```
curl -X POST -H "Content-type: application/json" \
  -d '{ "query": "{   
    groups {         
        id           
        name         
        tasks {      
            id       
            title 
            status   
        }            
    }                
  }" }'                \
  'http://localhost:4000/task/query'
```