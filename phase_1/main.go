package main


import (
	"fmt"
	"os"
	"encoding/json"
)


type Product struct {
	ID int  `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Stock int `json:"stock"`
}


func loadProducts(path string) ([]Product, error){
	b, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var ps []Product
	return ps, json.Unmarshal(b, &ps)
}

func main() {
	products, err := loadProducts("products.json")

	if err !=  nil {panic(err)}
	
	cart := map[int]int {}
	// TODO: simple loop: list/add/remove/checkout
	fmt.Println("mini shop ready. commands: list, add <id> <qty>, remove <id>, checkout, exit")
	_ = products; _ = cart

	for  {
		var cmd string
		fmt.Printf("> ")
		fmt.Scan(&cmd) 

		switch cmd {
		case "list" :
			for _, p := range products {
				fmt.Printf("%d: %s - $%.2f (Stock: %d)\n", p.ID, p.Name, p.Price, p.Stock)
			}
		
		case "add" :
			var id, qty int 
			fmt.Scan(&id, &qty)
			cart[id]+= qty
			fmt.Printf("Added %d of product %d\n", qty, id)

		case "remove":
			var id int
			fmt.Scan(&id)
			delete(cart, id)
			fmt.Printf("Removed product %d from cart\n", id)

		case "checkout":
			var total float64
			fmt.Printf("Your bill:")
			for id, qty := range cart {
				for _, p := range products {
					if p.ID == id {
						if p.Stock >= qty {
							fmt.Printf("- %s x%d = ₹%.2f\n", p.Name, qty, p.Price*float64(qty))
							total += p.Price * float64(qty)
							p.Stock -= qty
						} else {
							fmt.Printf("Not enough stock for %s\n", p.Name)
						}
					}
				}
			}
			fmt.Printf("Total = ₹%.2f\n", total)
			cart = map[int]int{} 
		case "exit":
			fmt.Println("Goodbye!!!!")
			return 
		
		default: 
			fmt.Println("Unknown command!!")
	}
}
}