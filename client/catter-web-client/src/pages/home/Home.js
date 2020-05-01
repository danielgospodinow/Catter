import React from 'react'
import logo from '../../images/cat2.png'
import Fact from "../../components/fact/Fact"
import CircularProgress from "@material-ui/core/CircularProgress"
import { withStyles } from '@material-ui/core/styles';
import './Home.css'

const ColorCircularProgress = withStyles({
  root: {
    color: '#4D4D4D',
  },
})(CircularProgress);

class Home extends React.Component {

  constructor(props) {
    super(props)

    this.state = {
      currentFact: "Loading ...",
      progress: 0,

      countdownResetId: 0
    }
  }

  componentDidMount() {
    const factUpdateIntervalMs = 10_000

    this.loadNewFact()
    this.resetCircle(factUpdateIntervalMs)

    let initialId = setInterval(() => {
      this.loadNewFact()
      this.resetCircle(factUpdateIntervalMs)
    }, factUpdateIntervalMs, 2000)

    this.setState({
      countdownResetId: initialId
    })
  }

  componentWillUnmount() {
    clearInterval(this.state.countdownResetId)
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} width="12%" className="App-logo" alt="logo" />
          <Fact text={this.state.currentFact}></Fact>
          <ColorCircularProgress className="progress" color="inherit" variant="static" value={this.state.progress} />
        </header>
      </div>
    )
  }

  loadNewFact() {
    fetch(process.env.REACT_APP_FACT_SERVICE_URL + "/api/fact/random")
      .then(res => res.json())
      .then(fact => {
        this.setState({
          currentFact: fact.content
        })
      })
      .catch(err => console.log(err))
  }

  resetCircle(countdownTimeMs) {
    const maxValue = 100
    const totalIterationsCount = 10
    const singleIterationMs = countdownTimeMs / totalIterationsCount
    const singleIterationValue = maxValue / totalIterationsCount

    clearInterval(this.state.countdownResetId)

    this.setState({
      progress: 0
    })

    const currentCountdownResetId = setInterval(() => {
      this.increaseCircle(singleIterationValue)
    }, singleIterationMs)

    this.setState({
      countdownResetId: currentCountdownResetId
    })
  }

  increaseCircle(value) {
    this.setState({
      progress: this.state.progress + value
    })
  }
}

export default Home
