import { CardContent, Fab, TextField } from '@material-ui/core';
import AppBar from '@material-ui/core/AppBar';
import Card from '@material-ui/core/Card';
import InputAdornment from '@material-ui/core/InputAdornment';
import { withStyles } from '@material-ui/core/styles';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import AddIcon from '@material-ui/icons/Add';
import PropTypes from 'prop-types';
import React from 'react';
import Slider from '@material-ui/lab/Slider';
import Icon from '@material-ui/core/Icon';
import Button from '@material-ui/core/Button';

import axios from 'axios';


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
    display: 'block',
    flexWrap: 'wrap',
  },
  textField: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
    marginTop: 0,
    marginBottom: 10,
    width: '100%',
  },
  fab: {
    margin: theme.spacing.unit,
  },
  header: {
    marginTop: 30,
    marginBottom: 0,
  },
  slider: {
    padding: '22px 0px',
    marginLeft: '10px'
  },
  rightIcon: {
    marginLeft: theme.spacing.unit,
  },
});

class App extends React.Component {
  state = {
    input: '',
    time: 0,
    outputs: [''],
  };

  handleChange = name => event => {
    this.setState({ [name]: event.target.value });
  };

  handleOutputChange = index => event => {
    this.state.outputs[index] = event.target.value;
    this.setState({})
  }

  handleSliderChange = (event, value) => {
    this.setState({ ['time']: value })
  }

  render() {
    const { classes } = this.props;

    return (
      <div className={classes.root}>
        <AppBar position="static" color="default">
          <Toolbar>
            <Typography variant="h6" color="inherit">
              Transparent Mixer
          </Typography>
          </Toolbar>
        </AppBar>
        <Card className={classes.card}>
          <CardContent>
            <Typography gutterBottom variant="h4" component="h2">
              New Transaction
            </Typography>
            <form className={classes.container} noValidate autoComplete="off">
              <Typography className={classes.header} variant="h6" component="h6">
                Inputs
              </Typography>
              <TextField
                id="standard-name"
                label="Amount"
                className={classes.textField}
                onChange={this.handleChange('input')}
                margin="normal"
                helperText="Number of JobCoins you're trying to clean."
                InputProps={{
                  endAdornment: <InputAdornment position="end">JC</InputAdornment>,
                }}
              />
              <Typography className={classes.header} variant="h6" component="h6">
                Outputs
              </Typography>

              {this.state.outputs.map((value, index) => {
                return <TextField
                  className={classes.textField}
                  id={'' + index}
                  key={'' + index}
                  value={this.state.outputs[index]}
                  onChange={
                    this.handleOutputChange(index)
                  }
                  label="Output Address"
                />
              })}

              <Fab
                color="primary"
                aria-label="Add"
                size="small"
                className={classes.fab}
                onClick={() => {
                  this.state.outputs.push('')
                  this.setState({})
                }}>
                <AddIcon />
              </Fab>

              <Typography className={classes.header} variant="h6" component="h6">
                Time
              </Typography>
              <Slider
                classes={{ container: classes.slider }}
                value={this.state.time}
                onChange={this.handleSliderChange}>
              </Slider>
              <p>{this.state.time}s</p> 
              <Button variant="contained" color="primary" className={classes.button}
                onClick={() => {
                  axios.post('/new-tx', this.state)
                    .then(this.handleTXRecv)
                    .catch(function (error) {
                      console.log(error);
                    })
                }}>
                Send
                <Icon className={classes.rightIcon}>send</Icon>
              </Button>
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