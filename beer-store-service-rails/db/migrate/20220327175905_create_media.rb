class CreateMedia < ActiveRecord::Migration[7.0]
  def change
    # https://apidock.com/rails/ActiveRecord/ConnectionAdapters/PostgreSQLAdapter/TableDefinition/primary_key
    create_table :media, id: false do |t|
      t.primary_key :idmedia
      t.timestamp :creationdatemedia, null: false, default: -> {'CURRENT_TIMESTAMP'}
      t.binary :datamedia, null: false
      t.string :nomemedia, null: false
      t.string :mimemedia, null: false
    end
  end
end
