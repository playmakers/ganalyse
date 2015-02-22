#!/usr/bin/env rails runner

require 'staccato'

@tracker = Staccato.tracker('UA-45979504-2')

def hit(shop, type, vendor, name, size, color, position,  product_id, variant_id, anz, price = nil)
  hit = Staccato::Pageview.new(@tracker, path: '/')
  hit.add_custom_dimension(1, shop)
  hit.add_custom_dimension(2, type)
  hit.add_custom_dimension(3, name)
  hit.add_custom_dimension(4, size)
  hit.add_custom_dimension(5, color)
  hit.add_custom_dimension(6, position)
  hit.add_custom_dimension(7, vendor)
  hit.add_custom_dimension(8, product_id)
  hit.add_custom_dimension(9, variant_id)
  hit.add_custom_metric(1, price) if price
  hit.add_custom_metric(2, anz)
  hit.track!
end

# puts hit('Forelle', 'Helm', 'Riddell', '360', 'L', 'Schwarz', nil, 8, nil)
# .group("DATE(created_at), HOUR(created_at)")

date = Time.parse('2014-04-12').utc
range = date...(date+12.hours)
query = WholesalerVariantQuantity.where(:created_at => range).includes(:wholesaler_variant => :product)

query.all.each do |q|
  puts  q.wholesaler_variant.product.title

  hit('Forelle',
    q.wholesaler_variant.product.type,
    q.wholesaler_variant.product.vendor,
    q.wholesaler_variant.product.title,
    q.wholesaler_variant.size || '-',
    q.wholesaler_variant.color || '-',
    q.wholesaler_variant.other || '-',
    q.wholesaler_variant.product_id,
    q.wholesaler_variant.variant_id,
    q.quantity
  )

  sleep(0.05)
end

# tracker.event(non_interactive: true, campaign_source: 'Forelle', category: 'Helm', action: 'Riddell 360',  label: 'Schwarz-XL', value: 10)
# tracker.event(non_interactive: true, campaign_source: 'Forelle', category: 'Helm', action: 'Riddell 3601', label: 'Schwarz-XL', value: 12)
# tracker.event(non_interactive: true, campaign_source: 'Forelle', category: 'Helm', action: 'Riddell 3602', label: 'Schwarz-XL', value: 103)
# tracker.event(non_interactive: true, campaign_source: 'Forelle', category: 'Helm', action: 'Riddell 3603', label: 'Schwarz-XL', value: 203)
# tracker.event(non_interactive: true, campaign_source: 'Forelle', category: 'Helm', action: 'Riddell 3604', label: 'Schwarz-XL', value: 1)


# tracker.transaction({
#   transaction_id: 1,
#   affiliation: 'clothing',
#   revenue: 17.98,
#   shipping: 2.00,
#   tax: 2.50,
#   currency: 'EUR'
# })


# tracker.transaction_item({
#   transaction_id: 1,
#   name: 'Shirt',
#   price: 8.99,
#   quantity: 2,
#   code: 'afhcka1230',
#   variation: 'red',
#   currency: 'EUR'
# })

