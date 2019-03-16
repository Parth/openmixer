import { StepLabel, Stepper, Step, StepContent } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import axios from 'axios';
import React from 'react';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/Card';

class TxStatus extends React.Component {
  state = {
    currentStep: 0,
  }

  componentDidMount() {
    this.pollStatus()
  }

  pollStatus = () => {
    axios
      .post('/tx-status', {
        'id': this.props.txId
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
      currentStep: value.data.current + 1,
    })
  }

  render() {
    const { classes } = this.props;
    return (
      <Card className={classes.card}>
        <CardContent>
          <Stepper activeStep={this.state.currentStep} orientation="vertical">
            <Step key={0}>
              <StepLabel>Awaiting Deposit to: {this.props.depositAddr}</StepLabel>
              <StepContent>Deposit TODO jobcoins to {this.props.depositAddr}</StepContent>
            </Step>
            {Array(this.props.outputCount).fill().map((label, index) => (
              <Step key={index+1}>
                <StepLabel>Scheduling Payment {index}</StepLabel>
              </Step>
            ))}
          </Stepper>
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
  root: {
    flexGrow: 1,
  },
});

export default withStyles(styles)(TxStatus)