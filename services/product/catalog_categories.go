package product

import "github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"

var dictCategoryTree = []*product.Category{
	{
		Id:   categoryFace,
		Name: "Лицо",
		Categories: []*product.Category{
			{
				Id:         categoryMask,
				Name:       "Маска",
				Categories: nil,
			},
			{
				Id:         categoryCream,
				Name:       "Крем",
				Categories: nil,
			},
			{
				Id:         categoryLotion,
				Name:       "Лосьон",
				Categories: nil,
			},
		},
	},
	{
		Id:   categoryHair,
		Name: "Волосы",
		Categories: []*product.Category{
			{
				Id:         categoryShampoo,
				Name:       "Шампунь",
				Categories: nil,
			},
			{
				Id:         categoryHairConditioner,
				Name:       "Кондиционер",
				Categories: nil,
			},
		},
	},
}

func getSubtree(id string, tree []*product.Category) []*product.Category {
	if len(tree) == 0 || id == "" {
		return tree
	}
	var subtree []*product.Category
	for _, node := range tree {
		if node.Id == id {
			return node.Categories
		}
		subtree = getSubtree(id, node.Categories)
		if len(subtree) > 0 {
			return subtree
		}
	}

	return subtree
}

func isValidPathInTree(path []string, tree []*product.Category) bool {
	if len(path) == 0 {
		return false
	}

	for _, node := range tree {
		if node.Id == path[0] {
			if len(path) == 1 {
				return true
			}
			return isValidPathInTree(path[1:], node.Categories)
		}
	}

	return false
}
