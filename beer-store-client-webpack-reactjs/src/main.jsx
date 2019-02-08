import React from "react";
import ReactDOM from "react-dom";
import {HashRouter as Router, Route, Switch, Redirect} from "react-router-dom";

import CssBaseline from "@material-ui/core/CssBaseline";

import {BeerListing} from "./views/beer-listing";
import {BeerDetails} from "./views/beer-details";

const Main = (props) => {
  return (
    <Router>
      <div>
        <CssBaseline />
        <Switch>
          <Redirect exact from="/" to="/beer-listing" />
          <Route path="/beer-listing" component={BeerListing} />
          <Route path="/beer-details/:id" component={BeerDetails} />
        </Switch>
      </div>
    </Router>
  );
};

ReactDOM.render(<Main />, document.getElementById("app"));
