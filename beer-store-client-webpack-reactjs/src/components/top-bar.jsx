import React from "react";

import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";

export const TopBar = props => {
  return (
    <AppBar position="static">
      <Toolbar>
        <Typography variant="h4" color="inherit">
          {props.left}
        </Typography>
      </Toolbar>
    </AppBar>
  );
};
