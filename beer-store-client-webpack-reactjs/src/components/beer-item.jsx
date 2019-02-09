import React from "react";

import {mediaservice} from "../api";

import {withStyles} from "@material-ui/core/styles";

import Avatar from "@material-ui/core/Avatar";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import ListItemSecondaryAction from "@material-ui/core/ListItemSecondaryAction";

const BeerItem_ = ({beer, children}) => {
  return (
    <ListItem>
      <ListItemIcon>
        <Avatar alt="beer icon" src={mediaservice.url(beer.idmedia)} />
      </ListItemIcon>
      <ListItemText primary={beer.titlebeer} secondary={beer.descriptionbeer}/>
      <ListItemSecondaryAction>{children}</ListItemSecondaryAction>
    </ListItem>
  );
};

export const BeerItem = withStyles({})(BeerItem_);
