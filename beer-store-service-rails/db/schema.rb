# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `bin/rails
# db:schema:load`. When creating a new database, `bin/rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema[7.0].define(version: 2022_03_27_175937) do
  create_table "beer", primary_key: "idbeer", force: :cascade do |t|
    t.datetime "creationdatebeer", precision: nil, default: -> { "CURRENT_TIMESTAMP" }, null: false
    t.string "titlebeer", null: false
    t.text "descriptionbeer"
    t.integer "idmedia"
  end

  create_table "media", primary_key: "idmedia", force: :cascade do |t|
    t.datetime "creationdatemedia", precision: nil, default: -> { "CURRENT_TIMESTAMP" }, null: false
    t.binary "datamedia", null: false
    t.string "nomemedia", null: false
    t.string "mimemedia", null: false
  end

  add_foreign_key "beer", "media", column: "idmedia", primary_key: "idmedia"
end
