class BeerController < ApplicationController
  def list
    page = params["page"] || "1"
    page_size = params["pageSize"] || "10"
    search = "%#{params["search"]}%" || "%%"

    p = {
      :q => search,
      :limit => page_size.to_i,
      :offset => (page.to_i - 1) * page_size.to_i
    }

    beers = Beer.limit(p[:limit]).offset(p[:offset])
    .where([%q(lower(titlebeer) like lower(:q) 
      or lower(descriptionbeer) like lower(:q)), p])
    render json: beers
  end

  def find
    render json: Beer.find(params["idbeer"])
  end
end
