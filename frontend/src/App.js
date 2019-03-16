import { withStyles } from '@material-ui/core/styles'
import PropTypes from 'prop-types'
import React from 'react'
import NewTx from './NewTx'
import TxStatus from './TxStatus'


class App extends React.Component {
  state = {
    depositAddress: '',
    txId: '',
    outputCount: 0,
  }

  onTx = (value) => {
    console.log(value)
    this.setState({
      depositAddr: value.data.depositAddress,
      txId: value.data.id,
      outputCount: value.data.spec.outputs.length
    })

    console.log(this.state)
  }

  handleStatus = (value) => {
    this.setState({
      txstatus: value.data
    })
  }

  render() {
    const { classes } = this.props

    return (
      <div className={classes.root}>
        <NewTx
          onTx={this.onTx} />
        {this.state.txId !== '' &&
          <TxStatus txId={this.state.txId} depositAddr={this.state.depositAddr} outputCount={this.state.outputCount}/>
        }
      </div>
    )
  }
}

App.propTypes = {
  classes: PropTypes.object.isRequired,
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
})

export default withStyles(styles)(App)
