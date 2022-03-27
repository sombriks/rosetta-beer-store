Rails.application.routes.draw do
  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html
  scope "/", defaults: {format: :json} do
    get "beer/list", to: "beer#list"
    get "beer/:idbeer", to: "beer#find"
  end
end
