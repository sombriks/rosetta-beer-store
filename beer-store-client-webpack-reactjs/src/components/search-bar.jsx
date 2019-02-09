import React from "react";

import FormControl from "@material-ui/core/FormControl";
import InputLabel from "@material-ui/core/InputLabel";
import Input from "@material-ui/core/Input";
import Button from "@material-ui/core/Button";

import Grid from "@material-ui/core/Grid";

import {withStyles} from "@material-ui/core/styles";

const styles = theme => ({
  container: {
    display: "flex",
    flexWrap: "wrap",
  },
  formControl: {
    margin: theme.spacing.unit,
    width:"100%"
  },
  button: {
    margin: theme.spacing.unit,
  },
});

class SearchBar_ extends React.Component {
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
    const {list, classes} = this.props;
    return (
      <Grid container spacing={24}>
        <Grid item xs={8}>
          <FormControl className={classes.formControl}>
            <InputLabel htmlFor="component-simple">Search</InputLabel>
            <Input id="component-simple" value={params.search} onChange={this.handleChange} />
          </FormControl>
        </Grid>
        <Grid item xs={2}>
          <Button
            variant="contained"
            color="primary"
            className={classes.button}
            disabled={params.page == 1}
            onClick={this.handlePrev}
          >
            Prev
          </Button>
        </Grid>
        <Grid item xs={2}>
          <Button
            variant="contained"
            color="primary"
            className={classes.button}
            disabled={list.length < params.pageSize}
            onClick={this.handleNext}
          >
            Next
          </Button>
        </Grid>
      </Grid>
    );
  }
}

export const SearchBar = withStyles(styles)(SearchBar_);
