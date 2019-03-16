import { CardContent, Fab, TextField } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import Card from '@material-ui/core/Card'
import Icon from '@material-ui/core/Icon'
import InputAdornment from '@material-ui/core/InputAdornment'
import Typography from '@material-ui/core/Typography'
import AddIcon from '@material-ui/icons/Add'
import Slider from '@material-ui/lab/Slider'
import axios from 'axios'
import React from 'react'
import { withStyles } from '@material-ui/core/styles'

class NewTx extends React.Component {
  handleInput = event => {
    this.setState({ input: parseFloat(event.target.value) })
  }

  handleOutputChange = index => event => {
    let val = event.target.value

    let outputs = [...this.state.outputs]
    let output = { ...outputs[index] }
    output.addr = val
    outputs[index] = output

    this.setState({ outputs })
  }

  handleSplitChange = index => event => {
    let val = parseFloat(event.target.value)

    let outputs = [...this.state.outputs]
    let output = { ...outputs[index] }
    output.split = val
    outputs[index] = output

    this.setState({ outputs })
  }

  handleNewOutput = () => {
    let outputs = [...this.state.outputs]
    outputs.push({
      addr: '',
      split: 100
    })

    for (var i = 0; i < outputs.length; i++) {
      outputs[i].split = 100 / outputs.length
    }

    this.setState({ outputs })
  }

  handleSliderChange = (event, value) => {
    this.setState({ time: value })
  }

  state = {
    input: 0,
    time: 0,
    outputs: [{
      addr: '',
      split: 100
    }],
    depositAddr: '',
  }

  render() {
    const { classes } = this.props
    return (
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
              return <div key={index}>
                <TextField
                  className={classes.outputAddr}
                  key={'addr' + index}
                  value={value.addr}
                  onChange={
                    this.handleOutputChange(index)
                  }
                  label="Output Address"
                />
                <TextField
                  className={classes.outputSplit}
                  key={'split' + index}
                  value={value.split}
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
              onClick={this.handleNewOutput}>
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
                  .then(this.props.onTx)
                  .catch(function (error) {
                    console.log(error)
                  })
              }}>
              Send
                <Icon className={classes.rightIcon}>send</Icon>
            </Button>
          </form>
        </CardContent>
      </Card>
    )
  }
}

const styles = theme => ({
  card: {
    maxWidth: '65%',
    margin: '0 auto',
    marginTop: 50,
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
})

export default withStyles(styles)(NewTx)
