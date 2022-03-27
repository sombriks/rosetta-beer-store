class CreateBeer < ActiveRecord::Migration[7.0]
  def change
    create_table :beer, id: false do |t|
      t.primary_key :idbeer
      t.timestamp :creationdatebeer, null: false, default: -> {'CURRENT_TIMESTAMP'}
      t.string :titlebeer, null: false
      t.text :descriptionbeer
      t.integer :idmedia
      t.foreign_key :media, column: :idmedia, primary_key: :idmedia
    end
  end
end
