import React from "react";

import {beerservice} from "../api";

import {TopBar} from "../components/top-bar";
import {BeerItem} from "../components/beer-item";

import Button from "@material-ui/core/Button";
import ArrowBackIcon from "@material-ui/icons/ArrowBack";
import {withStyles} from "@material-ui/core";

const style = theme => ({
  button: {
    margin: theme.spacing.unit,
  },
});

class BeerDetails_ extends React.Component {
  state = {beer: null};

  componentDidMount() {
    beerservice.find(this.props.match.params.id).then(ret => {
      this.setState({beer: ret.data});
    });
  }

  render() {
    const {classes} = this.props;
    const {beer} = this.state;
    return (
      <div>
        <TopBar left={"Beer Details"} />
        <Button variant="contained" color="primary" className={classes.button} href="#/beer-listing">
          <ArrowBackIcon />
        </Button>
        {beer && <BeerItem beer={beer} key={beer.idbeer} />}
      </div>
    );
  }
}

export const BeerDetails = withStyles(style)(BeerDetails_);
