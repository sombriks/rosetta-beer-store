import React from "react";

import {beerservice} from "../api";

import {TopBar} from "../components/top-bar";
import {SearchBar} from "../components/search-bar";
import {BeerItem} from "../components/beer-item";

export class BeerListing extends React.Component {
  // @babel/plugin-proposal-class-properties
  state = {params: {search: "", page: 1, pageSize: 10}, list: []};

  componentDidMount() {
    this.busca();
  }

  busca = params => {
    if (params) {
      this.state.params = params;
      this.setState(this.state);
    }
    beerservice.list(this.state.params).then(ret => {
      this.state.list = ret.data;
      this.setState(this.state);
      console.log("chegou service");
    });
  };

  render() {
    const {params, list} = this.state;
    return (
      <div>
        <TopBar left={"Beer Listing"} />
        <SearchBar params={params} list={list} busca={this.busca} />
        <ul>
          {list.map(beer => (
            <BeerItem beer={beer} />
          ))}
        </ul>
      </div>
    );
  }
}
