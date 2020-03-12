import React from 'react'
import './Fact.css'

class Fact extends React.Component {

  constructor(props) {
    super(props)

    this.state = {
      text: this.props.text
    }
  }

  static getDerivedStateFromProps(props, state) {
    if (props.text !== state.text) {
      return { text: props.text }
    } else {
      return null
    }
  }

  render() {
    return (
      <p>{this.state.text}</p>
    )
  }
}

export default Fact