# Questions 


## Operators

## It is possible to call several cases in the switch case construction
Yes, using keyword `fallthrogh`

```
func main() {
	animals := []string{"bear", "bear", "rabbit", "wolf"}

	for _, animal := range animals {
		switch animal {
      
		case "rabbit":
			fmt.Println(animal, "is so weak!")
			fallthrough
      
		case "bear", "wolf":
			fmt.Println(animal, "is so strong!")
		}
	}
}
```
**output:**
bear is so strong!
bear is so strong!
rabbit is so weak!
rabbit is so strong!
wolf is so strong!

