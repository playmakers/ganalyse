package vendors

import ()

func InspectDocA(productPage []byte, origin string) *Product {
	// sizeMapping := map[string]string {
	//   // "76": "XS",
	//   "1": "S",
	//   "2": "M",
	//   "3": "L",
	//   "4": "XL",
	//   // "31": "XXL",
	//   // "38": "3XL",
	//   // "39": "4XL",
	// }

	doc := Parse(productPage, "utf-8")

	product := &Product{
		Name: doc.Find(".right h1").Text(),
		Origin: origin,
	}

	price := NormPrice(doc.Find(".article_details_price strong").Text())

	product.AddVariant(DEFAULT_SIZE, DEFAULT_COLOR, price, 0)

	return product
}
