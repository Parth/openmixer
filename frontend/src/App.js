import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import InputAdornment from '@material-ui/core/InputAdornment';
import { CardContent, TextField } from '@material-ui/core';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Card from '@material-ui/core/Card';

const styles = theme => ({
  card: {
    maxWidth: '65%',
    margin: '0 auto',
    marginTop: 50,
  },
  root: {
    flexGrow: 1,
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
    width: 200,
  },
  dense: {
    marginTop: 19,
  },
  menu: {
    width: 200,
  },
});

class App extends React.Component {
  state = {
    input: '',
    output: [],
  };

  handleChange = name => event => {
    this.setState({ [name]: event.target.value });
  };

  render() {
    const { classes } = this.props;

    return (
      <div className={classes.root}>
        <AppBar position="static" color="default">
          <Toolbar>
            <Typography variant="h6" color="inherit">
              Jobcoin Mixer
          </Typography>
          </Toolbar>
        </AppBar>
        <Card className={classes.card}>
          <CardContent>
            <Typography gutterBottom variant="h4" component="h2">
              New Transaction
            </Typography>
            <form className={classes.container} noValidate autoComplete="off">
              <TextField
                id="standard-name"
                label="Amount"
                className={classes.textField}
                value={this.state.input}
                onChange={this.handleChange('name')}
                margin="normal"
                helperText="Input Amount in JCs"
                InputProps={{
                  endAdornment: <InputAdornment position="end">JC</InputAdornment>,
                }}
              />

            </form>
          </CardContent>
        </Card>
      </div>
    );
  }
}

App.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(App);