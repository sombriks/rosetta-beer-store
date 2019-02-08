import React from "react";

export class SearchBar extends React.Component {
  state = {params: {search: "", page: 1, pageSize: 10}};

  handleChange = ev => {
    this.state.params.search = ev.target.value;
    this.state.params.page = 1;
    this.setState(this.state);
    this.props.busca(this.state.params);
  };

  handlePrev = _ => {
    this.state.params.page--;
    this.setState(this.state);
    this.props.busca(this.state.params);
  };

  handleNext = _ => {
    this.state.params.page++;
    this.setState(this.state);
    this.props.busca(this.state.params);
  };

  render() {
    const {params} = this.state;
    const {list} = this.props;
    return (
      <div>
        <input value={params.search} onChange={this.handleChange} />
        <button disabled={params.page == 1} onClick={this.handlePrev}>
          Prev
        </button>
        <button disabled={list.length < params.pageSize} onClick={this.handleNext}>
          Next
        </button>
      </div>
    );
  }
}
