package data

import "github.com/gin-gonic/gin"

// ประกาศ Recipes เป็น public variable
var Recipes = map[int64]gin.H{
	1: {"id": 1, "name": "Spaghetti Carbonara", "ingredients": []string{"spaghetti", "eggs", "pancetta", "parmesan cheese", "black pepper"}, "instructions": "Cook spaghetti. In a bowl, mix eggs and cheese. Fry pancetta until crispy. Combine everything and season with black pepper."},
	2: {"id": 2, "name": "Chicken Curry", "ingredients": []string{"chicken", "curry powder", "coconut milk", "onions", "garlic", "ginger"}, "instructions": "Sauté onions, garlic, and ginger. Add chicken and curry powder. Pour in coconut milk and simmer until chicken is cooked."},
	3: {"id": 3, "name": "Caesar Salad", "ingredients": []string{"romaine lettuce", "croutons", "parmesan cheese", "Caesar dressing"}, "instructions": "Toss lettuce with croutons and dressing. Top with parmesan cheese."},
}
