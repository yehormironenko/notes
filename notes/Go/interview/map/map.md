## Map

What is the map in golang

Map - this is structure implementing hash functions. 
Map should be initialize. Because map refers to the buckets. 

Each bucket contains: 
 - 8 bits
 - link to the next collision bucket
 - 8 key-value pairs stored in an array



 ## Questions

 - Why we cannot use link for a value by key in map 
  Hash map links can be changed. 


  Each bucket can contains max 8 values. When it's more thant 6.5 then go make a resize our map.


  - "What are the syntax features for getting and writing values ​​in a map?"
  we cannot write without allocation -> panic

  If key doesn't found in the map we get a default value for int - 0 , for string - "". 
  To check the value better to use special syntax:

 ` if _, ok := mm[1]; ok {...}`

 ## How we are looking for a value in the map

 - the hash of key calculated
 - calculated the bucket for this hash 
 - calculated additional hash - 8 bit if hash 
 - then in the bucket we compare each from 8 additional hashes this our addition hash
 - if additional hash is equal our founding hash  -> return the link
 - if our hashes didn't equal then then we move to the next bucket (this link contain current bucket)
 - if link is not here and we cannot found value then return default