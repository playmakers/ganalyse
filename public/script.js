function renderProduct(product) {
  var variants = {},
  mainNode = $('<table border="1"></table>'),
  header = $("<tr>").appendTo(mainNode),
  sizes = {},
  defaultSizes = [];

  $.each(product.Variants, function(key, variant) {
    key1 = (variant.Color != "") ? variant.Color : variant.Position;
    key2 = (variant.Size != "") ? variant.Size : variant.Position;

    if (!variants[key1]) {
      variants[key1] = {};
    }
    variants[key1][key2] = variant;
    sizes[key2] = true;
  });

  defaultSizes = Object.keys(sizes);

  header.append($('<th><a target="_blank" href="' + product.Origin + '">origin</a></th>'));
  $.each(defaultSizes, function(key, size) {
    header.append($("<th>" + size + "</th>"));
  });

  $.each(variants, function(color, d) {
    node = $("<tr>");
    node.append($("<td><b>" + color + "</b></td>"));

    $.each(defaultSizes, function(key, size) {
      var variant = d[size],
      value, className;

      if(variant) {
        outofstock = (variant.Quantity < 1) ? " outofstock" : "";
        if (variant.Quantity || variant.Quantity == 0) {
          value = variant.Quantity;
          className = 'p'+ variant.Policy + outofstock;
        } else {
          value = variant.Price;
          className = 'av'+ variant.Availability;
        }
        node.append($('<td class="' + className + '">' + value + '</td>'));
      } else {
        node.append($('<td>-</td>'));
      }

    });
    node.appendTo(mainNode)
  });

  return mainNode;
}
