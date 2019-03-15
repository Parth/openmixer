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
    width: '80%',
  },
  outputAddr: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
    marginTop: 0,
    marginBottom: 10,
    width: '80%',
  },
  outputSplit: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
    marginTop: 0,
    marginBottom: 10,
    width: '10%',
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
    input: 0,
    time: 0,
    outputs: [''],
    splits: [100],
    depositAddr: '',
    txId: '',
    txstatus: []
  };

  handleInput = event => {
    this.setState({ input: parseFloat(event.target.value) });
  };

  handleOutputChange = index => event => {
    this.state.outputs[index] = event.target.value;
    this.setState({})
  }

  handleSplitChange = index => event => {
    this.state.splits[index] = event.target.value;
    this.setState({})
  }

  handleSliderChange = (event, value) => {
    this.setState({ time: value });
  }

  handleTXRecv = (value) => {
    console.log(value)
    this.setState({
      depositAddr: value.data.DepositAddr,
      txId: value.data.id
    })
    this.pollStatus()
  }

  pollStatus = () => {
    axios.post('/tx-status', {
      'id': this.state.txId
    })
      .then(this.handleStatus)
      .catch(function (error) {
        console.log(error);
      })

    setTimeout(() => {
      this.pollStatus()
    }, 500);
  }

  handleStatus = (value) => {
    this.setState({
      txstatus: value.data
    })
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
                label="Amount"
                type="number"
                className={classes.textField}
                onChange={this.handleInput}
                margin="normal"
                helperText="Number of JobCoins you're trying to anonymize."
                InputProps={{
                  endAdornment: <InputAdornment position="end">JC</InputAdornment>,
                }}
              />
              <Typography className={classes.header} variant="h6" component="h6">
                Outputs
              </Typography>

              {this.state.outputs.map((value, index) => {
                return <div>
                  <TextField
                    className={classes.outputAddr}
                    id={'' + index}
                    key={'' + index}
                    value={this.state.outputs[index]}
                    onChange={
                      this.handleOutputChange(index)
                    }
                    label="Output Address"
                  />
                  <TextField
                    className={classes.outputSplit}
                    id={'' + index}
                    key={'' + index}
                    value={this.state.splits[index]}
                    type="number"
                    onChange={
                      this.handleSplitChange(index)
                    }
                    InputProps={{
                      endAdornment: <InputAdornment position="end">%</InputAdornment>,
                    }}
                    label="Split Percentage"
                  />
                </div>
              })}

              <Fab
                color="primary"
                aria-label="Add"
                size="small"
                className={classes.fab}
                onClick={() => {
                  this.state.outputs.push('')
                  this.state.splits.push(1)

                  for (var i = 0; i < this.state.splits.length; i++) {
                    this.state.splits[i] = 100 / this.state.splits.length
                  }

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
        {this.state.depositAddr !== '' &&
          <Card className={classes.card}>
            <Typography className={classes.header} variant="h6" component="h6">
              Deposit 10 JC to {this.state.depositAddr}
            </Typography>
          </Card>
        }
        {this.state.txId !== '' &&
          <Card className={classes.card}>
            {this.state.txstatus.map((value, index) => {
              return (<Typography className={classes.header} variant="h6" component="h6">
                {value}
              </Typography>)
            })}
          </Card>
        }
      </div>
    );
  }
}

App.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(App);